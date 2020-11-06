package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// GetProviderSchema function
func (s *RawProviderServer) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) *tfprotov5.GetProviderSchemaResponse {

	cfgSchema := GetProviderConfigSchema()

	resSchema := GetProviderResourceSchema()

	return &tfprotov5.GetProviderSchemaResponse{
		Provider:        cfgSchema,
		ResourceSchemas: resSchema,
	}
}
