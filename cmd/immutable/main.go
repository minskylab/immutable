package main

import (
	"github.com/minskylab/immutable"
)

func main() {
	config, err := immutable.LoadConfigFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	record, err := immutable.NewRecord(config)
	if err != nil {
		panic(err)
	}

	if err := immutable.AddRecordToLogs(config, record); err != nil {
		panic(err)
	}

	// pp.Println(record.RecordLine())
}
