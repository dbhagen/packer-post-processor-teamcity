// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package main

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	TeamCityUrl         *string           `mapstructure:"teamcity_url" cty:"teamcity_url"`
	Username            *string           `mapstructure:"username" cty:"username"`
	Password            *string           `mapstructure:"password" cty:"password"`
	ProjectId           *string           `mapstructure:"project_id" cty:"project_id"`
	CloudImage          *string           `mapstructure:"cloud_image" cty:"cloud_image"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"teamcity_url":               &hcldec.AttrSpec{Name: "teamcity_url", Type: cty.String, Required: false},
		"username":                   &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                   &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"project_id":                 &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"cloud_image":                &hcldec.AttrSpec{Name: "cloud_image", Type: cty.String, Required: false},
	}
	return s
}