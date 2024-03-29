package env

import (
	"os"
	"testing"
)

func TestFill(t *testing.T) {
	for _, name := range []string{"A", "B", "C", "D"} {
		if err := os.Unsetenv(name); err != nil {
			t.Error(err)
		}
	}

	if err := os.Setenv("B", "true"); err != nil {
		t.Error(err)
	}

	type Config struct {
		A string `env:"A"`
		B bool   `env:"B"`
		C int    `env:"C" envDefault:"10"`
		D bool   `env:"D,expand" envDefault:"${B}"`
	}

	filled, err := Fill[Config]()
	if err != nil {
		t.Error(err)
	}

	if filled.A != "" {
		t.Error("filled.A must be empty string, got:", filled.A)
	}

	if filled.B != true {
		t.Error("filled.B must be true, got:", filled.B)
	}

	if filled.C != 10 {
		t.Error("filled.C must be 10, got:", filled.C)
	}

	if filled.D != true {
		t.Error("filled.D must be true, got:", filled.D)
	}
}

func TestEnv_Fill(t *testing.T) {
	for _, name := range []string{"A", "B", "C"} {
		if err := os.Unsetenv(name); err != nil {
			t.Error(err)
		}
	}

	if err := os.Setenv("B", "true"); err != nil {
		t.Error(err)
	}

	type Config struct {
		Env[Config]

		A string `env:"A"`
		B bool   `env:"B"`
		C int    `env:"C" envDefault:"10"`
	}

	filled, err := Config{}.Fill()
	if err != nil {
		t.Error(err)
	}

	if filled.A != "" {
		t.Error("filled.A must be empty string, got:", filled.A)
	}

	if filled.B != true {
		t.Error("filled.B must be true, got:", filled.B)
	}

	if filled.C != 10 {
		t.Error("filled.C must be 10, got:", filled.C)
	}
}
