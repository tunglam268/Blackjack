package main

import (
	"awesomeProject2/service"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	fmt.Println("Welcome to the Card Game!")

	for {
		//Start round
		indexMapResult := service.StartRound()

		for i, result := range indexMapResult {
			fmt.Println("-----------------RESULT-----------------")
			fmt.Printf("%d. %v\n", i, result)
		}

		//Restart round
		var playAgain int
		fmt.Print("\nDo you want to play again? (1=yes/2=no): ")
		fmt.Scanln(&playAgain)
		if playAgain == 2 {
			//Clear terminal
			CallClear()
			fmt.Println("Thanks for playing!")
			break
		}
		//Clear terminal
		CallClear()
	}
}
