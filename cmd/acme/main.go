package main

import (
	"flag"
	"log"
	"os"

	"acme/pkg/config"
	"acme/pkg/server"
)

func main() {
	var (
		conf       config.Config
		configFile string
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&configFile, "config.file", "", "configuration file to load")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatalf("error parsing flags: %v\n", err)
	}

	if configFile == "" {
		log.Fatalln("-config.file flag required")
	} else if err := config.LoadConfig(configFile, &conf); err != nil {
		log.Fatalf("error loading config file %s: %v\n", configFile, err)
	}

	server.Start(conf)
}
