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
	fmt.Println("start car ")
	err = wire.Init()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// beep
	// controller.BeepTimes(1, 1*time.Second)
	// led
	fmt.Println("led car ")
	controller.Led(1, 0, 1)
	time.Sleep(5 * time.Second)
	// controller.Led(0, 1, 0)
	// time.Sleep(5 * time.Second)

	fmt.Println("start run ")
	controller.MotorRun()
	time.Sleep(5 * time.Second)
	fmt.Println("stop car ")
	err = wire.Stop()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
