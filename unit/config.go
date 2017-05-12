package unit

import "github.com/jacobstr/confer"

// LoadConfig loads default test_config.yaml file
func LoadConfig() *confer.Config {
	config := confer.NewConfig()
	err := config.ReadPaths("test_config.yaml")
	if err != nil {
		panic(err)
	}
	return config
}
