package app

import (
	"go-blockchain-node-scanner/chain"
	"go-blockchain-node-scanner/env"
	"go-blockchain-node-scanner/repo"
)

type App struct {
	env *env.Env

	repo *repo.Repo
}

func NewApp(env *env.Env) {
	a := &App{env: env}

	var err error

	if a.repo, err = repo.NewRepo(env); err != nil {
		panic(err)
	} else {
		chain.ScanBlock(env, a.repo, env.Node.StartBlock, env.Node.EndBlock)

	}
}
