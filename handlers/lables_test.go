package handlers

import "testing"

func Test_getMinReplicaCount(t *testing.T) {
	scenarios := []struct {
		name   string
		labels *map[string]string
		output int
	}{
		{
			name:   "nil map returns default",
			labels: nil,
			output: initialReplicasCount,
		},
		{
			name:   "empty map returns default",
			labels: &map[string]string{},
			output: initialReplicasCount,
		},
		{
			name:   "empty map returns default",
			labels: &map[string]string{OFFunctionMinReplicaCount: "2"},
			output: 2,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			output := getMinReplicaCount(s.labels)
			if output == nil {
				t.Errorf("getMinReplicaCount should not return nil pointer")
			}

			value := int(*output)
			if value != s.output {
				t.Errorf("expected: %d, got %d", s.output, value)
			}
		})
	}
}

func Test_parseLabels(t *testing.T) {
	scenarios := []struct {
		name         string
		labels       *map[string]string
		functionName string
		output       map[string]string
	}{
		{
			name:         "nil map returns just the function name",
			labels:       nil,
			functionName: "testFunc",
			output:       map[string]string{OFFunctionNameLabel: "testFunc"},
		},
		{
			name:         "empty map returns just the function name",
			labels:       &map[string]string{},
			functionName: "testFunc",
			output:       map[string]string{OFFunctionNameLabel: "testFunc"},
		},
		{
			name:         "non-empty map does not overwrite the function name label",
			labels:       &map[string]string{OFFunctionNameLabel: "anotherValue", "customLabel": "test"},
			functionName: "testFunc",
			output:       map[string]string{OFFunctionNameLabel: "testFunc", "customLabel": "test"},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			output := parseLabels(s.functionName, s.labels)
			if output == nil {
				t.Errorf("parseLabels should not return nil map")
			}

			outputFuncName := output[OFFunctionNameLabel]
			if outputFuncName != s.functionName {
				t.Errorf("parseLabels should always set the function name: expected %s, got %s", s.functionName, outputFuncName)
			}

			for key, value := range output {
				expectedValue := s.output[key]
				if value != expectedValue {
					t.Errorf("Incorrect output label for %s, expected: %s, got %s", key, expectedValue, value)
				}
			}

		})
	}
}
