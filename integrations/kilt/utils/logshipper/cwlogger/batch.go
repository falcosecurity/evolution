package cwlogger

import (
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

const (
	maxBatchByteSize = 1048576
	maxBatchLength   = 10000
	logEventOverhead = 26
)

type batch struct {
	logEvents []cloudwatchlogs.InputLogEvent
	size      int
}

func newBatch() *batch {
	return &batch{
		logEvents: []cloudwatchlogs.InputLogEvent{},
	}
}

func (b *batch) add(logEvent cloudwatchlogs.InputLogEvent) (ok bool) {
	size := len(*logEvent.Message) + logEventOverhead
	if size+b.size <= maxBatchByteSize && len(b.logEvents) < maxBatchLength {
		b.logEvents = append(b.logEvents, logEvent)
		b.size += size
		return true
	}
	return false
}

func (b *batch) Len() int {
	return len(b.logEvents)
}

func (b *batch) Less(i, j int) bool {
	return *b.logEvents[i].Timestamp < *b.logEvents[j].Timestamp
}

func (b *batch) Swap(i, j int) {
	b.logEvents[i], b.logEvents[j] = b.logEvents[j], b.logEvents[i]
}

type batcher struct {
	input  chan cloudwatchlogs.InputLogEvent
	output chan []cloudwatchlogs.InputLogEvent
}

func newBatcher() *batcher {
	b := &batcher{
		input:  make(chan cloudwatchlogs.InputLogEvent),
		output: make(chan []cloudwatchlogs.InputLogEvent),
	}
	go b.worker()
	return b
}

func (br *batcher) flush() {
	close(br.input)
}

func (br *batcher) worker() {
	b := newBatch()
	timeout := time.After(time.Second)

	flush := func() {
		if len(b.logEvents) > 0 {
			sort.Sort(b)
			br.output <- b.logEvents
			b = newBatch()
		}
		timeout = time.After(time.Second)
	}

	for {
		select {
		case logEvent, ok := <-br.input:
			if !ok {
				flush()
				close(br.output)
				return
			}
			if ok := b.add(logEvent); !ok {
				flush()
				b.add(logEvent)
			}
		case <-timeout:
			flush()
		}
	}
}
