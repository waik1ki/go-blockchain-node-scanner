package app

import (
	"go-blockchain-scope/chain"
	"go-blockchain-scope/env"
	"go-blockchain-scope/repo"
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
