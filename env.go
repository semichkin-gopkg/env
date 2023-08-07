package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/semichkin-gopkg/conf"
)

func Fill[T any](updaters ...conf.Updater[Config]) (T, error) {
	var config T

	options := conf.New[Config]().Append(updaters...).Build()

	err := env.Parse(&config, env.Options{
		Environment:     options.Environments,
		TagName:         options.TagName,
		RequiredIfNoDef: options.RequiredIfNoDef,
		OnSet:           options.OnSet,
		Prefix:          options.Prefix,
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

type Env[T any] struct{}

func (e Env[T]) Fill(updaters ...conf.Updater[Config]) (T, error) {
	return Fill[T](updaters...)
}

func (e Env[T]) MustFill(updaters ...conf.Updater[Config]) T {
	return MustFill[T](updaters...)
}
