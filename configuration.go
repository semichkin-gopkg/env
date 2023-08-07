package env

import (
	"github.com/semichkin-gopkg/conf"
)

type (
	Environments = map[string]string
	OnSetFn      = func(tag string, value interface{}, isDefault bool)
)

type Config struct {
	// Environments keys and values that will be accessible for the service.
	Environments map[string]string

	// TagName specifies another tag name to use rather than the default env.
	TagName string

	// RequiredIfNoDef automatically sets all env as required if they do not declare 'envDefault'
	RequiredIfNoDef bool

	// OnSet allows to run a function when a value is set
	OnSet OnSetFn

	// Prefix define a prefix for each key
	Prefix string
}

func WithEnvironments(environments Environments) conf.Updater[Config] {
	return func(c *Config) {
		c.Environments = environments
	}
}

func WithEnvironment(key string, val string) conf.Updater[Config] {
	return func(c *Config) {
		if c.Environments == nil {
			c.Environments = Environments{}
		}

		c.Environments[key] = val
	}
}

func WithTagName(name string) conf.Updater[Config] {
	return func(c *Config) {
		c.TagName = name
	}
}

func WithRequiredIfNoDef(required bool) conf.Updater[Config] {
	return func(c *Config) {
		c.RequiredIfNoDef = required
	}
}

func WithOnSetFn(fn OnSetFn) conf.Updater[Config] {
	return func(c *Config) {
		c.OnSet = fn
	}
}

func WithPrefix(prefix string) conf.Updater[Config] {
	return func(c *Config) {
		c.Prefix = prefix
	}
}
