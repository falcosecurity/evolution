module kilt-cfn

go 1.14

require (
	github.com/Jeffail/gabs/v2 v2.6.0
	github.com/aws/aws-lambda-go v1.19.1
	github.com/aws/aws-sdk-go v1.34.27
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/rs/zerolog v1.19.0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/yudai/gojsondiff v1.0.0
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/admiral0/evolution/integrations/kilt/lib v0.0.0-20200921154944-8926dbb5c8ea
)

replace github.com/admiral0/evolution/integrations/kilt/lib => ./../../lib
