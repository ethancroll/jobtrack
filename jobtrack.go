package main

import (
	"fmt"
	"os"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Magenta = "\033[35m"

var input string
var started bool
var cash int

func main() {
	started = false
	go addcash()
	fmt.Println("welcome to jobtrack, type 'help' for commands")
	checkinput()
}

func checkinput() {
	fmt.Println("awating command")
	fmt.Scan(&input)
	if input == "help" {
		help()
	}
	if input == "start" && started == false {
		start()
	}
	if input == "start" && started == true {
		fmt.Println("timer already started.")
		checkinput()
	}
	if input == "pause" && started == true {
		pause()
	}
	if input == "pause" && started == false {
		fmt.Println("can not pause, timer not started")
		checkinput()
	}
	if input == "save" && started == false {
		save()
	}
	if input == "save" && started == true {
		fmt.Println("pause first to save")
		checkinput()
	}
	if input == "status" {
		status()
	} else {
		fmt.Println("command not recognised. type 'help' for commands")
		checkinput()
	}
}

func help() {
	fmt.Println(Green + "commands" + Reset)
	fmt.Println("start timer: 'start'")
	fmt.Println("pause timer: 'pause'")
	fmt.Println("save to txt: 'save'")
	fmt.Println("check status: 'status'")
	checkinput()
}

func start() {
	started = true
	fmt.Println(Green + "timer started" + Reset)
	checkinput()
}

func pause() {
	started = false
	fmt.Println(Red + "timer paused" + Reset)
	checkinput()
}

func status() {
	fmt.Println(Green+"current cash value:"+Reset, float64(cash)/100.0)
	checkinput()
}

func save() {
	fmt.Println(Magenta + "saving" + Reset)
	date := time.Now().Format("2006-01-02")
	file, _ := os.OpenFile("track.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s %.2f\n", date, float64(cash)/100.0))
	fmt.Println(Green + "save complete" + Reset)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(Red + "exiting" + Reset)
	time.Sleep(5 * time.Second)
	os.Exit(3)
}

func addcash() {
	for {
		if started == true {
			cash = cash + 1
			time.Sleep(3600 * time.Millisecond)
		}
	}
}
