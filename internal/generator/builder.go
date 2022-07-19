package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/razielsd/rzgrpcmock/internal/cli"
	"github.com/razielsd/rzgrpcmock/internal/generator/srcbuilder"
	"github.com/razielsd/rzgrpcmock/internal/generator/srcparser"
)

const (
	moduleName            = "github.com/razielsd/rzgrpcmock/server"
	generatedPathLocation = "/internal/generated"
)

type Builder struct {
	printer        *cli.InfoPrinter
	packageName    string
	packageVersion string
	packagePath    string
	fileList       []string
	projectDir     string
	fakeUsageFile  string
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
	if err := b.getPackage(packageName); err != nil {
		return err
	}
	if err := b.makeFakeUsage(); err != nil {
		return err
	}
	if err := b.goModVendor(); err != nil {
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
	if err := b.clean(); err != nil {
		return err
	}
	if err := b.goModVendor(); err != nil {
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

func (b *Builder) makeFakeUsage() error {
	b.printer.Action("Make fake usage")
	template := "package generated\nimport \"%s\""
	code := fmt.Sprintf(template, b.packageName)
	b.fakeUsageFile = b.projectDir + string(os.PathSeparator) + generatedPathLocation  +
		string(os.PathSeparator) + "fake_" + srcbuilder.MakeHash(b.packageName) + ".go"
	if err := ioutil.WriteFile(b.fakeUsageFile, []byte(code), 0644); err != nil {
		b.printer.Push(cli.StateFail)
		return err
	}
	b.printer.Push(cli.StateOk)
	return nil
}

func (b *Builder) clean() error {
	b.printer.Action("Clean")
	if err := os.Remove(b.fakeUsageFile); err != nil {
		b.printer.Push(cli.StateFail)
		return err
	}
	b.printer.Push(cli.StateOk)
	return nil
}

func (b *Builder) getPackage(pkgName string) error {
	b.printer.Action("Run go get package")
	if err := cli.ExecCmd(b.projectDir, "go", "get", pkgName); err != nil {
		b.printer.Push(cli.StateFail)
		log.Fatal(err)
	}
	b.printer.Push(cli.StateOk)
	return nil
}

func (b *Builder) searchPackage() error {
	b.printer.Action("Search package")
	locator := newPackageLocator(b.projectDir)
	path, err := locator.Search(b.packageName)
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
	parser := srcparser.NewFileParser(filename)
	if err := parser.Parse(); err != nil {
		return err
	}
	generator := &srcbuilder.Builder{
		ModuleName:       moduleName,
		ExportModuleName: b.packageName,
		PackageName:      parser.PackageName,
		SaveDir:          saveDir,
	}
	for _, field := range parser.GetInterfaceList() {
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

func (b *Builder) goModVendor() error {
	b.printer.Action("Run go mod vendor")
	if err := cli.ExecCmd(b.projectDir, "go", "mod", "vendor"); err != nil {
		b.printer.Push(cli.StateFail)
		log.Fatal(err)
	}
	b.printer.Push(cli.StateOk)
	return nil
}
