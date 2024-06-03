package main
import (
	"fmt"
	"time"
)

func loading() {
	clearscreen()
	fmt.Println("Loading")
	fmt.Println("█")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading")
	fmt.Println("████")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading")
	fmt.Println("██████")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading")
	fmt.Println("████████")
	time.Sleep(1 * time.Second)
	clearscreen()
	fmt.Println("Loading")
	fmt.Println("██████████")
	time.Sleep(1 * time.Second)
	clearscreen()
}