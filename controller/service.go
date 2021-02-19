package controller

import (
	"github.com/mateemu/wire"
)

type service struct {
	wire.BaseService
}

func (s service) init() error {

	return nil
}

// init module
func init() {
	svc := service{}
	wire.Append(svc)
}
