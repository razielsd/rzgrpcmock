package srcbuilder

const serviceTemplate = `package mockservice{{.Index}}

import (
	"context"
	{{.PackageName}} "{{.ModuleName}}"
	"go.uber.org/zap"
	"github.com/razielsd/rzgrpcmock/server/internal/reqmatcher"
	{{.ExtImport}}
)

type Service struct {
	{{.PackageName}}.Unimplemented{{ .InterfaceName }}
	log    *zap.Logger
}

func NewService(lg *zap.Logger) *Service {
	return &Service{
		log: lg,
	}
}
`

const handlerTemplate = `
func (s *Service) {{.Method}}({{.Args}}) (*{{.Response}}, error) {
	matcher := reqmatcher.GetMatcher(reqmatcher.DefaultMatcher)
	resp := &{{.Response}}{}
	err := matcher.Match(arg0 , arg1, resp)
	return resp, err
}
`

const registerHandlerTemplate = `package generated

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	svc{{ .Index }} "github.com/razielsd/rzgrpcmock/server/internal/generated/service_{{ .Index }}"
	sa{{ .Index }} "{{ .ExportModuleName }}"
)

func init() {
	f := func(s *grpc.Server, logger *zap.Logger) {
		service{{ .Index }} := svc{{ .Index }}.NewService(logger)
		sa{{ .Index }}.Register{{ .InterfaceName }}(s, service{{ .Index }})
	}
	mu.Lock()
	defer mu.Unlock()
	handlerRegistration = append(handlerRegistration, f)
}



`
