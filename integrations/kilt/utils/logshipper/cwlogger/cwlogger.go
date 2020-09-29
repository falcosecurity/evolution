package cwlogger

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// The Config for the Logger.
type Config struct {
	// The Amazon CloudWatch Logs client created with the AWS SDK for Go.
	// Required.
	Client *cloudwatchlogs.Client

	// The name of the log group to write logs into. Required.
	LogGroupName string

	// An optional function to report errors that couldn't be automatically
	// handled during a PutLogEvents API call and caused a log events to be
	// dropped.
	ErrorReporter func(err error)

	// An optional log group retention time in days. This value is only taken into
	// account when creating a log group that does not yet exist. Set to 0
	// (default) for no retention policy. Refer to the PutRetentionPolicy API
	// documentation for valid values.
	Retention int
}

// A Logger represents a single CloudWatch Logs log group.
type Logger struct {
	name          *string
	svc           *cloudwatchlogs.Client
	streams       *logStreams
	prefix        string
	batcher       *batcher
	wg            sync.WaitGroup
	done          chan bool
	errorReporter func(err error)
	retention     int
}

// New creates a new Logger.
//
// Creates the log group if it doesn't yet exist, and one initial log stream for
// writing logs into.
//
// Returns an error if the configuration is invalid, or if either the creation
// of the log group or log stream fail.
func New(config *Config) (*Logger, error) {
	if config.Client == nil {
		return nil, errors.New("cwlogger: config missing required Client")
	}

	if config.LogGroupName == "" {
		return nil, errors.New("cwlogger: config missing required LogGroupName")
	}

	errorReporter := noopErrorReporter
	if config.ErrorReporter != nil {
		errorReporter = config.ErrorReporter
	}

	lg := &Logger{
		errorReporter: errorReporter,
		name:          &config.LogGroupName,
		svc:           config.Client,
		retention:     config.Retention,
		prefix:        randomHex(32),
		batcher:       newBatcher(),
		done:          make(chan bool),
	}

	lg.streams = newLogStreams(lg)

	if err := lg.createIfNotExists(); err != nil {
		return nil, err
	}
	if err := lg.streams.new(); err != nil {
		return nil, err
	}

	go lg.worker()

	return lg, nil
}

// Log enqueues a log message to be written to a log stream.
//
// The log message must be less than 1,048,550 bytes, and the time must not be
// more than 2 hours in the future, 14 days in the past, or older than the
// retention period of the log group.
//
// This method is safe for concurrent access by multiple goroutines.
func (lg *Logger) Log(t time.Time, s string) {
	lg.wg.Add(1)
	go func() {
		lg.batcher.input <- cloudwatchlogs.InputLogEvent{
			Message:   &s,
			Timestamp: aws.Int64(t.UnixNano() / int64(time.Millisecond)),
		}
		lg.wg.Done()
	}()
}

// Close drains all enqueued log messages and writes them to CloudWatch Logs.
// This method blocks until all pending log messages are written.
//
// The Logger is not meant to be used anymore after this method is called.
// Doing so will result in a panic. Create a new Logger if you wish to write
// more logs.
func (lg *Logger) Close() {
	lg.wg.Wait()       // wait for all log entries to be accepted
	lg.batcher.flush() // wait for all log entries to be batched
	<-lg.done          // wait for all batches to be processed
	lg.streams.flush() // wait for all batches to be sent to CloudWatch Logs
}

func (lg *Logger) worker() {
	for batch := range lg.batcher.output {
		lg.streams.write(batch)
	}
	lg.done <- true
}

func (lg *Logger) createIfNotExists() error {
	ctx := context.Background()

	req := lg.svc.CreateLogGroupRequest(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: lg.name,
	})

	_, err := req.Send(ctx)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == cloudwatchlogs.ErrCodeResourceAlreadyExistsException {
				return nil
			}
		}
	}
	if lg.retention != 0 {
		req := lg.svc.PutRetentionPolicyRequest(&cloudwatchlogs.PutRetentionPolicyInput{
			LogGroupName:    lg.name,
			RetentionInDays: aws.Int64(int64(lg.retention)),
		})
		_, err = req.Send(ctx)
	}
	return err
}

type writeError struct {
	batch  []cloudwatchlogs.InputLogEvent
	stream *logStream
	err    error
}

type logStreams struct {
	logger  *Logger
	streams []*logStream
	writers map[*logStream]chan []cloudwatchlogs.InputLogEvent
	writes  chan []cloudwatchlogs.InputLogEvent
	errors  chan *writeError
	wg      sync.WaitGroup
}

func newLogStreams(lg *Logger) *logStreams {
	streams := &logStreams{
		logger:  lg,
		streams: []*logStream{},
		writers: make(map[*logStream]chan []cloudwatchlogs.InputLogEvent),
		writes:  make(chan []cloudwatchlogs.InputLogEvent),
		errors:  make(chan *writeError),
	}
	go streams.coordinator()
	return streams
}

func (ls *logStreams) new() error {
	name := ls.logger.prefix + "." + strconv.Itoa(len(ls.streams))
	stream := &logStream{
		name:   &name,
		logger: ls.logger,
	}

	err := stream.create()
	if err != nil {
		return err
	}

	ls.streams = append(ls.streams, stream)
	ls.writers[stream] = make(chan []cloudwatchlogs.InputLogEvent)
	go ls.writer(stream)

	return nil
}

func (ls *logStreams) write(b []cloudwatchlogs.InputLogEvent) {
	ls.wg.Add(1)
	go func() {
		ls.writes <- b
	}()
}

func (ls *logStreams) writer(stream *logStream) {
	for batch := range ls.writers[stream] {
		batch := batch // create new instance of batch for the goroutine
		err := stream.write(batch)
		if err != nil {
			go func() {
				ls.errors <- &writeError{
					batch:  batch,
					stream: stream,
					err:    err,
				}
			}()
		} else {
			ls.wg.Done()
		}
	}
}

func (ls *logStreams) coordinator() {
	i := 0
	for {
		select {
		case batch := <-ls.writes:
			i = (i + 1) % len(ls.streams)
			stream := ls.streams[i]
			ls.writers[stream] <- batch
		case err := <-ls.errors:
			ls.handle(err)
		}
	}
}

func (ls *logStreams) handle(writeErr *writeError) {
	if isErrorCode(writeErr.err, errCodeThrottlingException) {
		ls.new()
	}
	if shouldRetry(writeErr.err) {
		go func() {
			ls.writes <- writeErr.batch
		}()
	} else {
		ls.wg.Done()
		ls.logger.errorReporter(writeErr.err)
	}
}

func (ls *logStreams) flush() {
	ls.wg.Wait()
}

type logStream struct {
	name          *string
	logger        *Logger
	sequenceToken *string
}

func (ls *logStream) create() error {
	req := ls.logger.svc.CreateLogStreamRequest(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  ls.logger.name,
		LogStreamName: ls.name,
	})
	_, err := req.Send(context.Background())
	return err
}

func (ls *logStream) write(b []cloudwatchlogs.InputLogEvent) error {
	req := ls.logger.svc.PutLogEventsRequest(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  ls.logger.name,
		LogStreamName: ls.name,
		LogEvents:     b,
		SequenceToken: ls.sequenceToken,
	})

	req.Sign()
	resp, err := ls.logger.svc.Client.Config.HTTPClient.Do(req.HTTPRequest)

	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data putLogEventsSuccessResponse
		if err := dec.Decode(&data); err != nil {
			return err
		}
		ls.sequenceToken = &data.NextSequenceToken
	} else {
		var data putLogEventsErrorResponse
		if err := dec.Decode(&data); err != nil {
			return err
		}
		if data.ExpectedSequenceToken != nil {
			ls.sequenceToken = data.ExpectedSequenceToken
		}
		return Error{
			Code:    data.Code,
			Message: data.Message,
		}
	}

	return nil
}

func randomHex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}
