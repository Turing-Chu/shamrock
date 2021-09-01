// Author: Turing Zhu
// Date: 6/11/21 2:40 PM
// File: logger.go

// Package shamrock Common used function & struct
package shamrock

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type fieldKey string

// fieldMap allows customization of the key names for default fields.
type fieldMap map[fieldKey]string

// ShamrockFormatter Shamrock defined formatter: [timestamp] LEVEL file:LINE function message
type ShamrockFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey string

	// fieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	fieldMap: fieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	fieldMap fieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

func (f *ShamrockFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	if f.DataKey != "" {
		newData := make(logrus.Fields, 4)
		newData[f.DataKey] = data
		data = newData
	}

	prefixFieldClashes(data, f.fieldMap, entry.HasCaller())

	timestampFormat := f.TimestampFormat
	// if timestampFormat == "" {
	// 	timestampFormat = consts.TimestampFormat
	// }

	if !f.DisableTimestamp {
		data[f.fieldMap.resolve(logrus.FieldKeyTime)] = entry.Time.Format(timestampFormat)
	}

	msgLen := len(entry.Message)
	if entry.Message == "" {
		data[f.fieldMap.resolve(logrus.FieldKeyMsg)] = entry.Message
	} else if entry.Message[0] == '[' && entry.Message[msgLen-1] == ']' {
		data[f.fieldMap.resolve(logrus.FieldKeyMsg)] = entry.Message[1 : msgLen-1]
	} else {
		data[f.fieldMap.resolve(logrus.FieldKeyMsg)] = entry.Message
	}

	data[f.fieldMap.resolve(logrus.FieldKeyLevel)] = entry.Level.String()
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:L%d", entry.Caller.File, entry.Caller.Line)
		if f.CallerPrettyfier != nil {
			funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
		}
		if funcVal != "" {
			data[f.fieldMap.resolve(logrus.FieldKeyFunc)] = funcVal[strings.LastIndex(funcVal, ".")+1:]
		}
		if fileVal != ":L" {
			data[f.fieldMap.resolve(logrus.FieldKeyFile)] = fmt.Sprintf("%s:L%d", path.Base(entry.Caller.File), entry.Caller.Line)
		}
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// [time] LEVEL FILE:LINE:FUNCNAME msg
	var line string
	line = fmt.Sprintf("[%-26s] %s %s %-16s\t%s\n",
		data[f.fieldMap.resolve(logrus.FieldKeyTime)],
		strings.ToUpper(data[f.fieldMap.resolve(logrus.FieldKeyLevel)].(string)),
		data[f.fieldMap.resolve(logrus.FieldKeyFile)],
		data[f.fieldMap.resolve(logrus.FieldKeyFunc)],
		data[f.fieldMap.resolve(logrus.FieldKeyMsg)])

	if _, err := b.WriteString(line); err != nil {
		return nil, fmt.Errorf("failed to format msg, %v", err)
	}

	return b.Bytes(), nil
}

func (f fieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

// This is to not silently overwrite `time`, `msg`, `func` and `level` fields when
// dumping it. If this code wasn't there doing:
//
//  logrus.WithField("level", 1).Info("hello")
//
// Would just silently drop the user provided level. Instead with this code
// it'll logged as:
//
//  {"level": "info", "fields.level": 1, "msg": "hello", "time": "..."}
//
// It's not exported because it's still using Data in an opinionated way. It's to
// avoid code duplication between the two default formatters.
func prefixFieldClashes(data logrus.Fields, fieldMap fieldMap, reportCaller bool) {
	timeKey := fieldMap.resolve(logrus.FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["fields."+timeKey] = t
		delete(data, timeKey)
	}

	msgKey := fieldMap.resolve(logrus.FieldKeyMsg)
	if m, ok := data[msgKey]; ok {
		data["fields."+msgKey] = m
		delete(data, msgKey)
	}

	levelKey := fieldMap.resolve(logrus.FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["fields."+levelKey] = l
		delete(data, levelKey)
	}

	logrusErrKey := fieldMap.resolve(logrus.FieldKeyLogrusError)
	if l, ok := data[logrusErrKey]; ok {
		data["fields."+logrusErrKey] = l
		delete(data, logrusErrKey)
	}

	// If reportCaller is not set, 'func' will not conflict.
	if reportCaller {
		funcKey := fieldMap.resolve(logrus.FieldKeyFunc)
		if l, ok := data[funcKey]; ok {
			data["fields."+funcKey] = l
		}
		fileKey := fieldMap.resolve(logrus.FieldKeyFile)
		if l, ok := data[fileKey]; ok {
			data["fields."+fileKey] = l
		}
	}
}
