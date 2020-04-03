package main

import (
	"fmt"
	"github.com/headfirstgo/calender"
	"log"
)

func main() {
	event := calender.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
