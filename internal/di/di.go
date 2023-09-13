//go:build wireinject

package di

import (
	"github.com/davidtrse/go-clean-arch/internal/account/handler"
	accrepo "github.com/davidtrse/go-clean-arch/internal/account/repository"
	accsrv "github.com/davidtrse/go-clean-arch/internal/account/service"
	"github.com/google/wire"
)

func InitAccountHandler() *handler.AccountHandler {
	panic(wire.Build(
		initDB,
		accrepo.NewAccountRepository,
		accsrv.NewAccountService,
		handler.NewAccountHandler,
	))
}
