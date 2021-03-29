package controller

import (
	"github.com/mateemu/wire"
)

// init module
func init() {
	svc := service{}
	wire.Append(svc)
}

type service struct {
	wire.BaseService
}

func (s service) init() error {

	return nil
}
