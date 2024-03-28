package calculations

import (
	"github.com/sirupsen/logrus"
)

func Calculate(n int64, useLogging bool) int64 {
	if useLogging {
		logrus.Info("Start calculations...")
		logrus.Infof("Calculate %d!", n)
	}

	var result int64 = 1
	for i := int64(1); i <= n; i++ {
		result *= i
	}

	if useLogging {
		logrus.Info("Calculations complete!")
	}

	return result
}
