package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"strings"
)

func FromS3(path string, decompress bool) string {

	sess, err := session.NewSession()

	if err != nil {
		panic(fmt.Errorf("could not create aws session: %w", err))
	}

	splitPath := strings.SplitN(path, "/", 2)

	if len(splitPath) != 2 {
		panic(fmt.Errorf("invalid path specified: expected bucket/objectkey, got '%s'", path))
	}

	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: &splitPath[0],
		Key: &splitPath[1],
	})

	if err != nil {
		panic(fmt.Errorf("could not retrieve %s: %w", path, err))
	}

	config, err := ioutil.ReadAll(obj.Body)
	defer obj.Body.Close()

	if err != nil {
		panic(fmt.Errorf("could not read %s: %w", path, err))
	}

	if decompress {
		config = decompressBytes(config)
	}

	return string(config)
}