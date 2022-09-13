package processors

import (
	"context"

	"github.com/katalabut/money-tell-api/app/processors/auth"
	"github.com/katalabut/money-tell-api/app/processors/pays"
)

type Deleter interface {
	Start(ctx context.Context)
}

// Container контейнер процессоров
type Container struct {
	Pays *pays.Manager
	Auth *auth.Auth
}

// New создает и возвращает Container
func New(
	pays *pays.Manager,
	auth *auth.Auth,
) *Container {
	return &Container{
		Pays: pays,
		Auth: auth,
	}
}
