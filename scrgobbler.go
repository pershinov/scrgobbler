package scrgobbler

import (
	"fmt"
	"log"
	"strings"
)

type Severity uint8

const (
	SeverityDebug Severity = iota
	SeverityInfo
	SeverityWarn
	SeverityError
	SeverityCritical
)

const (
	severityDebugString    string = "debug"
	severityInfoString     string = "info"
	severityWarnString     string = "warn"
	severityErrorString    string = "error"
	severityCriticalString string = "critical"
)

var logSeverityMap = map[Severity]string{
	SeverityDebug:    severityDebugString,
	SeverityInfo:     severityInfoString,
	SeverityWarn:     severityWarnString,
	SeverityError:    severityErrorString,
	SeverityCritical: severityCriticalString,
}

func Log(severity Severity, message, method string, params ...any) string {
	logString := fmt.Sprintf(
		"- severity: %s, message: %s, method: %s",
		logSeverityMap[severity],
		message,
		method)

	variadic := make([]string, 0, len(params))
	if params != nil && len(params) != 0 {
		for _, p := range params {
			if p == nil {
				continue
			}
			variadic = append(variadic, fmt.Sprintf("%v", p))
		}
	}
	if len(variadic) != 0 {
		logString += ", params: "
		logString += strings.Join(variadic, ", ")
	}
	log.Println(logString)

	return logString
}
