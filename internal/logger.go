package runner

import (
	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

// GetLogger returns the singleton logger instance.
// This centralizes logger access for the runner package and allows for
// easy replacement of the logging implementation in the future.
func GetLogger() *utils.LoggerImpl {
	return utils.GetLogger()
}
