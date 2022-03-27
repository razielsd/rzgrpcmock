package main

import (
	"fmt"
	"github.com/razielsd/rzgrpcmock/builder/internal/srcbuilder"
	"github.com/razielsd/rzgrpcmock/builder/internal/srcparser"
	"log"
	"os"
	"strings"
)

const moduleName = "github.com/razielsd/rzgrpcmock/server"

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Require 3 argument: grpc-file save-dir module-name")
	}
	saveDir := os.Args[2]
	exportModuleName := os.Args[3]
	extractor := srcparser.NewInterfaceExtractor()
	err := extractor.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
	}
	builder := &srcbuilder.Builder{
		ModuleName:       moduleName,
		ExportModuleName: exportModuleName,
		PackageName:      extractor.PackageName,
		SaveDir:          saveDir,
	}
	for _, field := range extractor.InterfaceList {

		if !strings.HasSuffix(field.Name, "Server") {
			continue
		}
		if strings.HasPrefix(field.Name, "Unsafe") {
			continue
		}
		err := builder.BuildService(field)
		if err != nil {
			fmt.Printf("ERR: %s", err.Error())
		}
	}
}
