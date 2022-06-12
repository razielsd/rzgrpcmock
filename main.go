package main

import (
	"embed"
	"github.com/razielsd/rzgrpcmock/internal/cmd"
	"github.com/razielsd/rzgrpcmock/internal/template"
)

var (
	//go:embed template
	res embed.FS
)
func main() {
	template.Resource = res
	cmd.Execute()
}
