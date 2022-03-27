package srcbuilder

const serviceTemplate = `package mockservice{{.Index}}

import (
	"context"
	{{.PackageName}} "{{.ModuleName}}"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	{{.PackageName}}.Unimplemented{{ .InterfaceName }}
	log    *log.Logger
}

func NewService(lg *log.Logger) *Service {
	return &Service{log: lg}
}
`

const handlerTemplate = `
func (s *Service) {{.Method}}({{.Args}}) (*{{.Response}}, error) {
	resp := &{{.Response}}{}
	return resp, nil
}
`

const registerHandlerTemplate = `package generated

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	svc{{ .Index }} "github.com/razielsd/rzgrpcmock/server/internal/generated/service_{{ .Index }}"
	sa{{ .Index }} "{{ .ExportModuleName }}"
)

func init() {
	f := func(s *grpc.Server, logger *log.Logger) {
		service{{ .Index }} := svc{{ .Index }}.NewService(logger)
		sa{{ .Index }}.Register{{ .InterfaceName }}(s, service{{ .Index }})
	}
	mu.Lock()
	defer mu.Unlock()
	handlerRegistration = append(handlerRegistration, f)
}



`
