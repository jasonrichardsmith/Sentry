package limits

import (
	"reflect"

	"github.com/jasonrichardsmith/sentry/config"
	"github.com/jasonrichardsmith/sentry/sentry"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	NAME = "limits"
)

func init() {
	config.Decoder(NAME, QtyHookFunc)
	config.Register(&Config{})
}

type Config struct {
	CPU    MinMax `yaml:"cpu"`
	Memory MinMax `yaml:"memory"`
}

type MinMax struct {
	Min resource.Quantity `yaml:"min"`
	Max resource.Quantity `yaml:"max"`
}

func (c *Config) Name() string {
	return NAME
}

func (c *Config) LoadSentry() sentry.Sentry {
	return LimitSentry{c}
}

func QtyHookFunc(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String {
		return data, nil
	}
	if t != reflect.TypeOf(resource.Quantity{}) {
		return data, nil
	}
	return resource.ParseQuantity(data.(string))

}
