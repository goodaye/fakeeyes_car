package main

// main enterence
//
import (
	"fmt"
	"time"

	"github.com/mateemu/fakeeyes_car/controller"
	"github.com/mateemu/wire"
)

func main() {
	var err error
	fmt.Println("abc")
	err = wire.Init()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	controller.BeepTimes(3, 1*time.Second)
}
