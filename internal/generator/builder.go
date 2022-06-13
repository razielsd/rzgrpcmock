package generator

import (
	"fmt"
	"github.com/razielsd/rzgrpcmock/internal/cli"
	"github.com/razielsd/rzgrpcmock/internal/generator/srcbuilder"
	"github.com/razielsd/rzgrpcmock/internal/generator/srcparser"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	moduleName = "github.com/razielsd/rzgrpcmock/server"
	generatedPathLocation = "/internal/generated"
)

type Builder struct {
	printer *cli.InfoPrinter
	packageName string
	packageVersion string
	packagePath string
	fileList []string
	projectDir string
}

func NewBuilder() *Builder {
	return &Builder{
		printer: cli.NewInfoPrinter(),
	}
}

func (b *Builder) Run(projectDir, packageName string) error {
	b.projectDir = strings.TrimSuffix(projectDir, string(os.PathSeparator))
	if err := b.extractPackageName(packageName); err != nil {
		return err
	}
	if err := b.searchPackage(); err != nil {
		return err
	}
	if err := b.searchGrpcServerSpec(); err != nil {
		return err
	}
	if err := b.generateMockServer(); err != nil {
		return err
	}
	if err := b.goModTidy(); err != nil {
		return err
	}
	return nil
}

func (b *Builder) extractPackageName(packageName string) error {
	b.printer.Action("Check package name")
	parts := strings.SplitN(packageName, "@", 2)
	if len(parts) != 2 {
		b.printer.Push(cli.StateFail)
		return fmt.Errorf("invalid package format: name@version, got: %s\n", packageName)
	}
	b.printer.Push(cli.StateOk)
	b.packageName = parts[0]
	b.packageVersion = parts[1]
	return nil
}

func (b *Builder) searchPackage() error {
	b.printer.Action("Search package")
	locator := newPackageLocator()
	path, err := locator.Search(b.packageName, b.packageVersion)
	if err != nil {
		b.printer.Push(cli.StateFail)
		return err
	}
	b.printer.Push(cli.StateOk)
	b.packagePath = path
	return nil
}

func (b *Builder) searchGrpcServerSpec() error {
	b.printer.Action("Search grpc-server spec")

	cmd := exec.Command("find", b.packagePath, "-name", "*grpc.pb.go")
	stdout, err := cmd.Output()
	if err != nil {
		b.printer.Push(cli.StateFail)
		return err
	}
	b.fileList = strings.Split(strings.TrimSpace(string(stdout)), "\n")
	b.printer.Push(cli.StateOk)
	b.printer.Action(fmt.Sprintf("Found %d grpc-server header", len(b.fileList)))
	if len(b.fileList) == 0 {
		b.printer.Push(cli.StateFail)
		return fmt.Errorf("no grpc-server header found")
	}
	b.printer.Push(cli.StateOk)
	return nil
}

func (b *Builder) generateMockServer() error {
	saveDir := b.projectDir + string(os.PathSeparator) + generatedPathLocation
	for _, filename := range b.fileList {
		b.printer.Action(fmt.Sprintf("Generate mock: %s", filename))
		err := b.build(filename, saveDir)
		if err != nil {
			b.printer.Push(cli.StateFail)
			return err
		}
		b.printer.Push(cli.StateOk)
	}
	return nil
}

func (b *Builder) build(filename, saveDir string) error {
	extractor := srcparser.NewInterfaceExtractor()
	err := extractor.Parse(filename)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
	}
	generator := &srcbuilder.Builder{
		ModuleName:       moduleName,
		ExportModuleName: b.packageName,
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
		err := generator.BuildService(field)
		if err != nil {
			fmt.Printf("ERR: %s", err.Error())
		}
	}
	return nil
}

func (b *Builder) goModTidy() error {
	b.printer.Action("Run go mod tidy")
	if err := cli.ExecCmd(b.projectDir,"go", "mod", "tidy"); err != nil {
		b.printer.Push(cli.StateFail)
		log.Fatal(err)
	}
	b.printer.Push(cli.StateOk)
	return nil
}
