package main

import (
	"fmt"
	"os"
	"time"
)

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
	fmt.Scan(&input)
	if input == "help" {
		help()
	}
	if input == "start" {
		start()
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
	}
	if input == "status" {
		status()
	}
}

func help() {
	fmt.Println("commands")
	fmt.Println("start timer: 'start'")
	fmt.Println("pause timer: 'pause'")
	fmt.Println("save to txt: 'save'")
	fmt.Println("check status: 'status'")
	checkinput()
}

func start() {
	started = true
	fmt.Println("timer started")
	checkinput()
}

func pause() {
	started = false
	fmt.Println("timer paused")
	checkinput()
}

func status() {
	fmt.Println("current cash value:", float64(cash)/100.0)
	checkinput()
}

func save() {
	fmt.Println("saving")
	date := time.Now().Format("2006-01-02")
	file, _ := os.OpenFile("track.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s %.2f\n", date, float64(cash)/100.0))
	fmt.Println("save complete")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("exiting")
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
