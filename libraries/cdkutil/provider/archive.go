package provider

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-archive-go/archive/v10/provider"
)

func NewArchiveProvider(scope constructs.Construct) provider.ArchiveProvider {
	return provider.NewArchiveProvider(scope, jsii.String("archive-provider"), &provider.ArchiveProviderConfig{})
}
