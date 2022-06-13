package template

import (
	"embed"
	"fmt"
	"github.com/razielsd/rzgrpcmock/internal/cli"
	"log"
	"os"
	"strings"
)

var (
	Resource embed.FS
)

type ExtractorService struct {
	printer *cli.InfoPrinter
	path    string
}

func NewTemplateService() *ExtractorService {
	return &ExtractorService{
		printer: cli.NewInfoPrinter(),
	}
}

func (t *ExtractorService) Init(path string) error {
	t.path = path
	if err := t.createDir(); err != nil {
		return err
	}

	if err := t.copyFiles(); err != nil {
		return err
	}

	if err := t.goModInit(); err != nil {
		return err
	}

	if err := t.goModTidy(); err != nil {
		return err
	}

	return nil
}

func (t *ExtractorService) Clean(path string) error {
	t.path = path
	if err := t.removeDir(); err != nil {
		return err
	}

	return t.Init(path)
}

func (t *ExtractorService) removeDir() error {
	if _, err := os.Stat(t.path); err != nil {
		if os.IsNotExist(err) {
			t.printer.Action(fmt.Sprintf("Directory not exists: %s", t.path))
			t.printer.Push(cli.StateOk)
			return nil
		}
		t.printer.Action(fmt.Sprintf("Directory stat: %s", t.path))
		t.printer.Push(cli.StateFail)
		return err
	}
	t.printer.Action(fmt.Sprintf("Remove directory: %s", t.path))
	err := os.RemoveAll(t.path)
	if err != nil {
		t.printer.Push(cli.StateFail)
		return err
	}
	t.printer.Push(cli.StateOk)

	return nil
}


func (t *ExtractorService) createDir() error {
	t.printer.Action(fmt.Sprintf("Create dir: %s", t.path))
	err := os.MkdirAll(t.path, 0750)
	if err != nil {
		t.printer.Push(cli.StateFail)
		return err
	}
	t.printer.Push(cli.StateOk)

	return nil
}

func (t *ExtractorService) copyFiles() error {
	t.printer.Action(fmt.Sprintf("Copy files into %s", t.path))
	err := copyDir(Resource, t.path, "template", "")
	if err != nil {
		t.printer.Push(cli.StateFail)
		return fmt.Errorf("copy error: %w", err)
	}
	t.printer.Push(cli.StateOk)

	return nil
}

func (t *ExtractorService) goModInit() error {
	gomodPath := t.path + string(os.PathSeparator) + "go.mod"
	if _, err := os.Stat(gomodPath); err == nil {
		t.printer.Action("go.mod allready exists")
		t.printer.Push(cli.StateOk)
		return nil
	}
	t.printer.Action("Run go mod init")
	if err := cli.ExecCmd(t.path, "go", "mod", "init", "github.com/razielsd/rzgrpcmock/template"); err != nil {
		t.printer.Push(cli.StateFail)
		log.Fatal(err)
	}
	t.printer.Push(cli.StateOk)

	return nil
}

func (t *ExtractorService) goModTidy() error {
	t.printer.Action("Run go mod tidy")
	if err := cli.ExecCmd(t.path,"go", "mod", "tidy"); err != nil {
		t.printer.Push(cli.StateFail)
		log.Fatal(err)
	}
	t.printer.Push(cli.StateOk)
	return nil
}

func copyDir(fs embed.FS, origin, fsDirName, dirName string) error {
	files, err := fs.ReadDir(fsDirName)
	if err != nil {
		return err
	}
	saveDir := fmt.Sprintf("%s/%s", origin, dirName)
	saveDir = strings.TrimSuffix(saveDir, "/")
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		err := os.MkdirAll(saveDir, 0750)
		if err != nil {
			return err
		}
	}

	for _, file := range files {
		if file.IsDir() {
			err := copyDir(
				fs,
				origin,
				fmt.Sprintf("%s/%s", fsDirName, file.Name()),
				fmt.Sprintf("%s/%s", dirName, file.Name()),
			)
			if err != nil {
				return err
			}
			continue
		}
		fileContent, err := fs.ReadFile(fmt.Sprintf("%s/%s", fsDirName, file.Name()))
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("%s/%s", saveDir, file.Name())
		if err := os.WriteFile(filename, fileContent, 0666); err != nil {
			return err
		}
	}
	return nil
}
