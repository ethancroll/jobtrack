package main

import (
	"fmt"
	"time"
	"os"
)

var initinput string
var started bool
var cash int

func main() {
	go timer()
	fmt.Println("welcome to the job portal, type 'help' for commands")
	fmt.Scan(&initinput)
	started = false
	help()
}

func help() {
	if initinput == "help" {
		fmt.Println("commands")
		fmt.Println("start timer: 'start'")
		fmt.Println("pause timer: 'pause'")
		fmt.Println("save to txt: 'save'")
		fmt.Println("check status: 'status'")
		welcomeback()
	} else {
		start()
	}
}

func start() {
	if initinput == "start" {
		fmt.Println("timer started")
		started = true
		welcomeback()
	} else {
		pause()
	}
}

func pause() {
	if initinput == "pause" && started == true {
		fmt.Println("timer paused")
		started = false
		welcomeback()
	} else if initinput == "pause" {
		fmt.Println("track has not started yet")
		welcomeback()
	} else {
		save()
	}
}

func save() {
	if initinput == "save" && started == false {
		fmt.Println("saving")
		date := time.Now().Format("2006-01-02")
		file, _ := os.OpenFile("track.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s %.2f\n", date, float64(cash)/100.0))
		fmt.Println("save complete")	
		welcomeback()
	} else if initinput == "save" {
		fmt.Println("track is still running. pause to save")
		welcomeback()
	} else {
		status()	
	}
}

func status() {
	if initinput == "status" {
		fmt.Println("current cash value:", float64(cash)/100.0)
		welcomeback()
	} else {
		fmt.Println("sorry, that was not recognized. try 'help' for functions")
		welcomeback()
	}
}

func welcomeback() {
	fmt.Println("awating next action...")
	fmt.Scan(&initinput)
	help()
}

func timer() {
	for {
	if started == true {
		cash = cash + 1
		time.Sleep(3600 * time.Millisecond)
	}
}
}
