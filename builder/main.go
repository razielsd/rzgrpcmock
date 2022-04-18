package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/razielsd/rzgrpcmock/builder/internal/srcbuilder"
	"github.com/razielsd/rzgrpcmock/builder/internal/srcparser"
)

const moduleName = "github.com/razielsd/rzgrpcmock/server"

func main() {
	const argCount = 4
	if len(os.Args) < argCount {
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
