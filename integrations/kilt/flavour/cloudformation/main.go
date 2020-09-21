package main

import (
	"context"
	"encoding/json"
	"kilt-cfn/config"
	"os"

	"kilt-cfn/cfnpatcher"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MacroInput struct {
	Region string `json:"region"`
	AccountID string `json:"accountId"`
	RequestID string `json:"requestId"`
	TransformID string `json:"transformId"`
	Fragment json.RawMessage `json:"fragment"`
}

type MacroOutput struct {
	RequestID string `json:"requestId"`
	Status string `json:"status"`
	Fragment json.RawMessage `json:"fragment"`
}



func HandleRequest(configuration *cfnpatcher.Configuration, ctx context.Context, event MacroInput) (MacroOutput, error) {
	l := log.With().
		Str("region", event.Region).
		Str("account", event.AccountID).
		Str("requestId", event.RequestID).
		Str("transformId", event.TransformID).
		Logger()
	loggerCtx := l.WithContext(ctx)
	result, err := cfnpatcher.Patch(loggerCtx, configuration, event.Fragment)
	if err != nil {
		return MacroOutput{event.RequestID, "failure", result}, err
	}
	return MacroOutput{event.RequestID, "success", result}, nil
}



func ConfigClosure() interface{} {
	definition := os.Getenv("KILT_DEFINITION")
	definitionType := os.Getenv("KILT_DEFINITION_TYPE")
	optIn := os.Getenv("KILT_OPT_IN")
	var fullDefinition string
	switch definitionType {
	case config.S3:
		fullDefinition = config.FromS3(definition, false)
	case config.S3Gz:
		fullDefinition = config.FromS3(definition, true)
	case config.Http:
		fullDefinition = config.FromWeb(definition)
	case config.Base64:
		fullDefinition = config.FromBase64(definition, false)
	case config.Base64Gz:
		fullDefinition = config.FromBase64(definition, true)
	default:
		panic("unrecognized definition type - " + definitionType)
	}

	configuration := &cfnpatcher.Configuration{
		Kilt:  fullDefinition,
		OptIn: optIn != "",
	}

	return func(ctx context.Context, event MacroInput) (MacroOutput, error) {
		return HandleRequest(configuration, ctx, event)
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	lambda.Start(ConfigClosure())
}


