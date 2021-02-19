package clutch

import (
	"github.com/mateemu/wire"
	"github.com/stianeikeland/go-rpio/v4"
)

type service struct {
	wire.BaseService
}

func (s service) Init() error {

	var err error
	err = rpio.Open()
	return err
}

func (s service) Stop() error {
	var err error
	err = rpio.Close()
	return err
}

func init() {

	svc := service{}
	wire.Append(svc)
}
