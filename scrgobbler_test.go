package scrgobbler_test

import (
	"testing"

	"github.com/pershinov/scrgobbler"
)

func TestLog(t *testing.T) {
	const methodName string = "TestLog"
	message := "test message"

	testCases := []struct {
		title    string
		severity scrgobbler.Severity
		params   any
		variadic bool

		expectedLog string
	}{
		{
			title:       "debug",
			severity:    scrgobbler.SeverityDebug,
			expectedLog: "- severity: debug, message: test message, method: TestLog",
		},
		{
			title:    "info with params map",
			severity: scrgobbler.SeverityInfo,
			params: map[string]string{
				"key": "value",
			},
			expectedLog: "- severity: info, message: test message, method: TestLog, params: map[key:value]",
		},
		{
			title:       "warn with param string",
			severity:    scrgobbler.SeverityWarn,
			params:      "string",
			expectedLog: "- severity: warn, message: test message, method: TestLog, params: string",
		},
		{
			title:       "error with params int slice",
			severity:    scrgobbler.SeverityError,
			params:      []int{1, 2, 3},
			expectedLog: "- severity: error, message: test message, method: TestLog, params: [1 2 3]",
		},
		{
			title:       "critical with variadic",
			severity:    scrgobbler.SeverityCritical,
			params:      []any{1, "string", map[string]bool{"test": true}},
			variadic:    true,
			expectedLog: "- severity: critical, message: test message, method: TestLog, params: 1, string, map[test:true]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			var log string
			if testCase.variadic {
				params, ok := testCase.params.([]any)
				if !ok {
					t.Errorf("cast %v to []any", params)
				}
				log = scrgobbler.Log(testCase.severity, message, methodName, params...)
			} else {
				log = scrgobbler.Log(testCase.severity, message, methodName, testCase.params)
			}
			if log != testCase.expectedLog {
				t.Errorf("log:\n%s\nis not expected:\n%s", log, testCase.expectedLog)
			}
		})
	}
}
