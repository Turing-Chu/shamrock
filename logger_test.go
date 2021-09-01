// Author: Turing Zhu
// Date: 2021/9/1 3:41 PM
// File: logger_test.go

package shamrock

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestShamrockFormatter(t *testing.T)  {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&ShamrockFormatter{
		TimestampFormat: StdMicroTimeFmt,
	})

	logger.Info("")
	logger.Info("message test")
}
