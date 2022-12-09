package processors

import (
	"context"

	"github.com/katalabut/money-tell-api/app/processors/auth"
	"github.com/katalabut/money-tell-api/app/processors/transactions"
)

type Deleter interface {
	Start(ctx context.Context)
}

// Container контейнер процессоров
type Container struct {
	Txn  *transactions.Manager
	Auth *auth.Auth
}

// New создает и возвращает Container
func New(
	txn *transactions.Manager,
	auth *auth.Auth,
) *Container {
	return &Container{
		Txn:  txn,
		Auth: auth,
	}
}
