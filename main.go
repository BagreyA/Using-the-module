package main

import (
	"factorial/calculations"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	n := int64(5)
	var log bool
	flag.BoolVar(&log, "log", false, "enable logging")
	flag.Parse()

	result := calculations.Calculate(n, log)
	fmt.Println("Factorial:", result)

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Factorial:", result)
}
