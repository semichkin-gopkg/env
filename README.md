# Lib for filling struct with env values

This library is a wrapper over [caarlos0/env](https://github.com/caarlos0/env) with generics support and more simple syntax

## Example
```go
package main

import (
	"fmt"
	"github.com/semichkin-gopkg/env"
	"log"
	"os"
)

type Config struct {
	A string `env:"A" envDefault:"A"`
	B string `env:"B" envDefault:"$A B" envExpand:"true"`

	Slice []int `env:"SLICE" envDefault:"1:2:3" envSeparator:":"`
}

func main() {
	for _, name := range []string{"A", "B", "SLICE"} {
		_ = os.Unsetenv(name)
	}

	config, _ := env.Fill[Config](env.WithOnSetFn(func(tag string, value interface{}, isDefault bool) {
		_ = os.Setenv(tag, fmt.Sprintf("%v", value))
	}))

	log.Println(config.A)     // A
	log.Println(config.B)     // A B
	log.Println(config.Slice) // [1 2 3]
}
```