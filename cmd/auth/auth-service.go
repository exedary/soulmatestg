package main

import "github.com/exedary/soulmates/internal/runner"

func main() {
	runner := runner.
		NewGinRunner().
		UseConfiguration("..\\..\\configs\\users-service\\config.yaml").
		UseSwagger()

	runner.Run()
}
