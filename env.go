package env

import (
	"github.com/caarlos0/env/v9"
	"github.com/semichkin-gopkg/conf"
	"reflect"
)

func Fill[T any](updaters ...conf.Updater[Config]) (T, error) {
	var config T

	options := conf.New[Config]().Append(updaters...).Build()

	parsers := make(map[reflect.Type]env.ParserFunc, len(options.Parsers))
	for t, p := range options.Parsers {
		parsers[t] = p
	}

	err := env.ParseWithOptions(&config, env.Options{
		Environment:     options.Environments,
		TagName:         options.TagName,
		RequiredIfNoDef: options.RequiredIfNoDef,
		OnSet:           options.OnSet,
		Prefix:          options.Prefix,
		FuncMap:         parsers,
	})

	return config, err
}

func MustFill[T any](updaters ...conf.Updater[Config]) T {
	filled, err := Fill[T](updaters...)
	if err != nil {
		panic(err)
	}

	return filled
}

type Env[T any] struct {
}

func (e Env[T]) Fill(updaters ...conf.Updater[Config]) (T, error) {
	return Fill[T](updaters...)
}

func (e Env[T]) MustFill(updaters ...conf.Updater[Config]) T {
	return MustFill[T](updaters...)
}
