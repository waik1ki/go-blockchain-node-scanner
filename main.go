package main

import (
	"flag"
	"go-blockchain-scope/app"
	"go-blockchain-scope/env"
)

var envFlag = flag.String("env", "./env.toml", "env not found")

func main() {
	// app > root.go 만 호출

	flag.Parse()
	app.NewApp(env.NewEnv(*envFlag))
}
