package testutil

import (
	"os"

	"github.com/sirupsen/logrus"
)

func IsIntegration() {
	if os.Getenv("TEST_TYPE") != "integration" {
		logrus.Infoln("integration test skipped")
		os.Exit(0)
	}
}
