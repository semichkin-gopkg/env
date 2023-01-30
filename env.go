package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/semichkin-gopkg/configurator"
)

func Fill[T any](updaters ...configurator.Updater[Configuration]) (T, error) {
	var config T

	options := configurator.New[Configuration]().Append(updaters...).Apply()

	err := env.Parse(&config, env.Options{
		Environment:     options.Environments,
		TagName:         options.TagName,
		RequiredIfNoDef: options.RequiredIfNoDef,
		OnSet:           options.OnSet,
		Prefix:          options.Prefix,
	})

	return config, err
}

func MustFill[T any](updaters ...configurator.Updater[Configuration]) T {
	filled, err := Fill[T](updaters...)
	if err != nil {
		panic(err)
	}

	return filled
}
