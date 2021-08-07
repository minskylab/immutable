package main

import "github.com/minskylab/immutable"

func main() {
	config, err := immutable.LoadConfigFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	if err := immutable.GeneratePDF(config); err != nil {
		panic(err)
	}
}
