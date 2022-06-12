package srcbuilder

import (
	"bytes"
	"crypto/md5" //nolint:gosec
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/razielsd/rzgrpcmock/internal/generator/srcparser"
)

const generatedFilePermission = 0o755

type Builder struct {
	ModuleName       string
	ExportModuleName string
	PackageName      string
	SaveDir          string
	Key              string
}

type serviceItem struct {
	Index            int
	InterfaceName    string
	ExportModuleName string
}

func (b *Builder) buildRegisterHandler(item serviceItem) error {
	params := map[string]string{
		"Index":            b.Key,
		"InterfaceName":    item.InterfaceName,
		"ExportModuleName": item.ExportModuleName,
	}
	t := template.New("")
	tmpl, err := t.Parse(registerHandlerTemplate)
	if err != nil {
		return err
	}

	src := bytes.NewBufferString("")
	err = tmpl.Execute(src, params)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s/register_handler_%s.go", b.SaveDir, params["Index"])
	err = overwriteFile(filename, src.String())
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) BuildService(field *srcparser.InterfaceField) error {
	b.Key = makeHash(b.ExportModuleName + "/" + field.Name)
	src := b.buildServiceHandler(field)
	path := fmt.Sprintf("%s/service_%s", b.SaveDir, b.Key)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename := fmt.Sprintf("%s/service_%s.go", path, b.Key)
	err := overwriteFile(filename, src)
	if err != nil {
		return err
	}
	service := serviceItem{
		InterfaceName:    field.Name,
		ExportModuleName: b.ExportModuleName,
	}
	return b.buildRegisterHandler(service)
}

func (b *Builder) buildServiceHandler(field *srcparser.InterfaceField) string {
	handlerSrc, _ := b.buildServiceHeader(field)
	for _, v := range field.MethodList {
		src, _ := b.buildHandler(v)
		handlerSrc += src
	}
	return handlerSrc
}

func (b *Builder) buildServiceHeader(field *srcparser.InterfaceField) (string, error) {
	t := template.New("")
	tmpl, err := t.Parse(serviceTemplate)
	if err != nil {
		return "", err
	}
	extImport := ""
	for k, importSpec := range field.ImportList {
		extImport += fmt.Sprintf("%s %s\n", k, importSpec.Path.Value)
	}
	params := map[string]interface{}{
		"PackageName":   b.PackageName,
		"ModuleName":    b.ExportModuleName,
		"Index":         b.Key,
		"InterfaceName": field.Name,
		"ServiceName":   strings.TrimSuffix(field.Name, "Server"),
		"ExtImport":     template.HTML(extImport), //nolint: gosec
	}
	src := bytes.NewBufferString("")
	err = tmpl.Execute(src, params)
	if err != nil {
		return "", err
	}
	return src.String(), nil
}

func (b *Builder) buildHandler(method *srcparser.InterfaceMethod) (string, error) {
	t := template.New("")
	tmpl, err := t.Parse(handlerTemplate)
	if err != nil {
		return "", err
	}
	var argList []string
	for i, v := range method.Args {
		arg := fmt.Sprintf("arg%d %s", i, v.String())
		argList = append(argList, arg)
	}
	params := map[string]string{
		"Method":   method.Name,
		"Response": method.Result[0].FullName(),
		"Args":     strings.Join(argList, ", "),
	}

	src := bytes.NewBufferString("")
	err = tmpl.Execute(src, params)
	if err != nil {
		return "", err
	}
	return src.String(), nil
}

func overwriteFile(filename, data string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, generatedFilePermission)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func makeHash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data)) //nolint:gosec
}
