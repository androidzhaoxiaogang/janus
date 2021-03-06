package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/hellofresh/janus/pkg/oauth"
	"github.com/hellofresh/janus/pkg/proxy"
)

// Spec Holds an api definition and basic options
type Spec struct {
	*Definition
	Manager oauth.Manager
}

// Config represents the configuration for a plugin
type Config map[string]interface{}

// Plugin represents the plugins for an API
type Plugin struct {
	Name    string `bson:"name" json:"name"`
	Enabled bool   `bson:"enabled" json:"enabled"`
	Config  Config `bson:"config" json:"config"`
}

// Definition Represents an API that you want to proxy
type Definition struct {
	Name    string            `bson:"name" json:"name" valid:"required"`
	Active  bool              `bson:"active" json:"active"`
	Proxy   *proxy.Definition `bson:"proxy" json:"proxy"`
	Plugins []Plugin          `bson:"plugins" json:"plugins"`
}

// NewDefinition creates a new API Definition with default values
func NewDefinition() *Definition {
	return &Definition{
		Active:  true,
		Plugins: make([]Plugin, 0),
		Proxy:   proxy.NewDefinition(),
	}
}

// Validate validates proxy data
func (d *Definition) Validate() (bool, error) {
	return govalidator.ValidateStruct(d)
}
