package main

import (
	"factorial/calculations"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main1() {
	args := os.Args
	n, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	var useLogging bool
	flag.BoolVar(&useLogging, "log", false, "enable logging")
	flag.Parse()

	result := calculations.Calculate(n, useLogging)

	if useLogging {
		logrus.Info("Calculations complete!")
	}
	fmt.Println("Factorial result:", result)
}
