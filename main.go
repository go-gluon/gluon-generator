package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-gluon/gondex"
	"golang.org/x/tools/go/packages"
)

var (
	version    = `0.0.0`
	resourcesF = flag.String("resources", "resources/*", "Resources directory")
	debugF     = flag.Bool("debug", false, "Enable debug logging")
	versionF   = flag.Bool("version", false, "Print version and exit")
)

type generatorConfig struct {
	debug bool
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gluon - .... %s.\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  %s [flags] [packages or directories]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  go generate [flags] [packages or files] (with '//go:generate gluon' in files)\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionF {
		fmt.Println(version)
		os.Exit(0)
	}

	config := generatorConfig{
		debug: *debugF,
	}
	// load golang indexer for current project
	gondexConfig := gondex.CreateDefaultConfig()
	gondexConfig.Mode = packages.NeedModule
	gondexConfig.Debug = *debugF
	// config.DefaultPattern = append(config.DefaultPattern, "github.com/go-gluon/gluon")
	indexer := gondex.CreateIndexer(gondexConfig)
	if e := indexer.Load(); e != nil {
		panic(e)
	}

	// check resource directory
	resources := *resourcesF
	if _, err := os.Stat(strings.TrimSuffix(resources, "/*")); os.IsNotExist(err) {
		panic(fmt.Errorf("Resources directory %s does not exists!", resources))
	}

	// generate configuration
	if err := ConfigGenerator(indexer, config); err != nil {
		panic(err)
	}

	// generate extension
	if err := ExtensionGenerator(indexer, config); err != nil {
		panic(err)
	}
}
