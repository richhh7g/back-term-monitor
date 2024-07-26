package main

import (
	"fmt"
	"path"

	"github.com/richhh7g/term-alarms/pkg/environment"
)

func init() {
	err := environment.NewEnvLoader(&environment.EnvLoaderParams{
		Type: "env",
		File: ".env",
		Path: path.Join("cmd", "server"),
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(environment.Get[string]("APP_ENV"))
}
