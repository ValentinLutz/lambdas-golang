package provider

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-random-go/random/v11/provider"
)

func NewRandomProvider(scope constructs.Construct) provider.RandomProvider {
	return provider.NewRandomProvider(scope, jsii.String("random-provider"), &provider.RandomProviderConfig{})
}
