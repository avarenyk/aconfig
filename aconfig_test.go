package aconfig

import "testing"

func TestGetConfigLocation(t *testing.T) {

	var tests = []struct {
		executable            string
		configLocation        string
		optionsConfigLocation string
		expected              string
	}{
		{"/go/src/project/", "/var/lib/conf.json", "", "/var/lib/conf.json"},
		{"/go/src/project/", "/var/lib/conf.json", "/opt/conf.json", "/var/lib/conf.json"},
		{"/go/src/project/", "conf.json", "/opt/conf.json", "/opt/conf.json"},
		{"/go/src/project/", defaultConfigLocation, "", "/go/src/project/conf.json"},
	}

	for _, test := range tests {
		executable = test.executable
		configLocation = test.configLocation
		loadOptions := &LoadOptions{test.optionsConfigLocation}

		location := getConfigLocation(loadOptions)
		if location != test.expected {
			t.Errorf("Test: %v, expected %s, received %s", test, test.expected, location)
		}
	}
}
