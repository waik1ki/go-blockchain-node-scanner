package main

import (
	"flag"
	"go-blockchain-node-scanner/app"
	"go-blockchain-node-scanner/env"
)

var envFlag = flag.String("env", "./env.toml", "env not found")

func main() {
	// app > root.go 만 호출

	flag.Parse()
	app.NewApp(env.NewEnv(*envFlag))
}
