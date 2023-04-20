package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/semichkin-gopkg/conf"
)

func Fill[T any](updaters ...conf.Updater[Configuration]) (T, error) {
	var config T

	options := conf.New[Configuration]().Append(updaters...).Build()

	err := env.Parse(&config, env.Options{
		Environment:     options.Environments,
		TagName:         options.TagName,
		RequiredIfNoDef: options.RequiredIfNoDef,
		OnSet:           options.OnSet,
		Prefix:          options.Prefix,
	})

	return config, err
}

func MustFill[T any](updaters ...conf.Updater[Configuration]) T {
	filled, err := Fill[T](updaters...)
	if err != nil {
		panic(err)
	}

	return filled
}
