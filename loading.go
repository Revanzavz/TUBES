package main
import (
	"fmt"
	"time"
)

func loading() {
	clearscreen()
	fmt.Println("Loading.")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading..")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading...")
	time.Sleep(1 * time.Second)
	clearscreen()
}