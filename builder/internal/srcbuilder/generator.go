package srcbuilder

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/razielsd/rzgrpcmock/builder/internal/srcparser"
	"html/template"
	"log"
	"os"
	"strings"
)

type Builder struct {
	ModuleName       string
	ExportModuleName string
	PackageName      string
	SaveDir          string
	Key              string
	serviceList      []serviceItem
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
	src, err := b.buildServiceHandler(field)
	if err != nil {
		log.Fatalln(err)
	}
	path := fmt.Sprintf("%s/service_%s", b.SaveDir, b.Key)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename := fmt.Sprintf("%s/service_%s.go", path, b.Key)
	err = overwriteFile(filename, src)
	if err != nil {
		return err
	}
	service := serviceItem{
		InterfaceName:    field.Name,
		ExportModuleName: b.ExportModuleName,
	}
	b.buildRegisterHandler(service)
	return nil
}

func (b *Builder) buildServiceHandler(field *srcparser.InterfaceField) (string, error) {
	handlerSrc, _ := b.buildServiceHeader(field)
	for _, v := range field.MethodList {
		src, _ := b.buildHandler(v)
		handlerSrc = handlerSrc + src
	}
	return handlerSrc, nil
}

func (b *Builder) buildServiceHeader(field *srcparser.InterfaceField) (string, error) {
	t := template.New("")
	tmpl, err := t.Parse(serviceTemplate)
	if err != nil {
		return "", err
	}
	params := map[string]string{
		"PackageName":   b.PackageName,
		"ModuleName":    b.ExportModuleName,
		"Index":         b.Key,
		"InterfaceName": field.Name,
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
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	f.Write([]byte(data))
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func makeHash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}
