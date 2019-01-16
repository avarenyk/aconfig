# config

Package provide method to load json config from file to custom structs.

config file could be located at same dir as executable and named **conf.json**
or path to config file could be provided by :
- command flag -config-location
- at load method via LoadOptions

# Usage

```
type MyConfig struct {
    option1     string
    option2     int
}

func main (){
    config := MyConfig{}
    // make sure reference to config is passed to function
    LoadJsonConfig(&config, nill)
    // or 
    LoadJsonConfig(&config, &LoadOptions{"/path/to/config.json"})
}

```