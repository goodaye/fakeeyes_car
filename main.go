package main

// main enterence
//
import (
	"fmt"

	"github.com/mateemu/fakeeyes_car/controller"
	"github.com/mateemu/wire"
)

func main() {
	var err error
	fmt.Println("abc")
	controller.TurnLeft()
	err = wire.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
