package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultConfigLocation = "conf.json"

type LoadOptions struct {
	configLocation string //absolute path to conf file
}

var (
	defaultLoadOptions = LoadOptions{
		configLocation: "",
	}
	configLocation string
	executable     string
)

func init() {
	var err error
	executable, err = os.Executable()
	if err != nil {
		panic(err)
	}
	flag.StringVar(&configLocation, "config-location", defaultConfigLocation, "an absolute path to config file")
}

func LoadJsonConfig(config interface{}, options *LoadOptions) {

	if options == nil {
		options = &defaultLoadOptions
	}

	// first step is to get json file path and check that file exists
	var (
		configLocation string
		err            error
	)
	// if error handling set to Exit or Panic do it before return
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	// parse command line arguments
	if !flag.Parsed() {
		flag.Parse()
	}

	configLocation = getConfigLocation(options)

	// check that given path is absolute
	if !isValidConfigLocation(configLocation) {
		err = fmt.Errorf("path to config file must be an absolute. %s given", configLocation)
		return
	}

	// read contents of file
	fileContents, err := ioutil.ReadFile(configLocation)
	if err != nil {
		return
	}

	// unmarshal json to a provided struct
	err = json.Unmarshal(fileContents, config)
}

// get config location path
func getConfigLocation(options *LoadOptions) string {

	// in priority config given in command line
	if configLocation != defaultConfigLocation {
		return configLocation
	}

	// if path provided in options use it
	if len(options.configLocation) > 0 {
		return options.configLocation
	}

	return filepath.Join(filepath.Dir(executable), configLocation)
}

// check config location path
func isValidConfigLocation(configLocation string) bool {
	return filepath.IsAbs(configLocation)
}
