package main

import (
	"fmt"

	"github.com/go-gluon/gondex"
)

func ConfigGenerator(indexer *gondex.Indexer, config generatorConfig) error {

	configReader := indexer.Interface("github.com/go-gluon/gluon/config.ConfigReader")
	// find all struct gluon:config
	tmp := indexer.FindStructsByAnnotation("gluon:config")
	for _, e := range tmp {
		if !e.Implements(configReader) {
			fmt.Printf("Config %v\n", e.Id())
		}
	}
	return nil
}
