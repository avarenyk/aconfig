# aconfig

Package provide method to load json config from file to custom structs.

config file could be located at same dir as executable and named **conf.json**
or path to config file could be provided by :
- command flag -config-location
- at load method via LoadOptions

# Usage

```
package main

import (
	"github.com/avarenyk/config"
	"log"
)

type MyConfig struct {
	Option1 string
	Option2 int
}

func main() {
	c := MyConfig{}
	// make sure reference to config is passed to function
	config.LoadJsonConfig(&c, nil)
	// or
	config.LoadJsonConfig(&c, &config.LoadOptions{"/absolute/path/to/config.json"})

	log.Fatal(c)
}

```