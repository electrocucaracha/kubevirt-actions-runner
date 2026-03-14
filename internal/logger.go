package runner

import (
	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

// This file provides a centralized logging utility for the runner package.
// It ensures consistent logging behavior and allows for easy replacement of the logging implementation.

// GetLogger returns the singleton logger instance.
func GetLogger() *utils.LoggerImpl {
	return utils.GetLogger()
}
