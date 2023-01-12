package main

import "github.com/exedary/soulmates/internal/runner"

func main() {
	runner := runner.
		NewGinRunner().
		UseConfiguration("").
		UseSwagger()

	runner.Run()
}
