package env

import (
	"github.com/semichkin-gopkg/conf"
	"reflect"
)

type (
	Environments = map[string]string
	OnSetFn      = func(tag string, value interface{}, isDefault bool)
	ParserFunc   = func(v string) (interface{}, error)
)

type Config struct {
	// Environments keys and values that will be accessible for the service.
	Environments Environments

	// TagName specifies another tag name to use rather than the default env.
	TagName string

	// RequiredIfNoDef automatically sets all env as required if they do not declare 'envDefault'
	RequiredIfNoDef bool

	// OnSet allows to run a function when a value is set
	OnSet OnSetFn

	// Prefix define a prefix for each key
	Prefix string

	// Parsers defines parse functions for different types.
	Parsers map[reflect.Type]ParserFunc
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

func WithParsers(p map[reflect.Type]ParserFunc) conf.Updater[Config] {
	return func(c *Config) {
		c.Parsers = p
	}
}

func WithParser(t reflect.Type, f ParserFunc) conf.Updater[Config] {
	return func(c *Config) {
		if c.Parsers == nil {
			c.Parsers = map[reflect.Type]ParserFunc{}
		}

		c.Parsers[t] = f
	}
}
