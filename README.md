# RzGrpcMock

[![codecov](https://codecov.io/gh/razielsd/rzgrpcmock/branch/master/graph/badge.svg)](https://codecov.io/gh/razielsd/rzgrpcmock)
[![Go Report Card](https://goreportcard.com/badge/github.com/razielsd/rzgrpcmock)](https://goreportcard.com/report/github.com/razielsd/rzgrpcmock)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/razielsd/rzgrpcmock)

## Using
```
RzGrpcMock service cli

Usage:
  rzgrpcmock [command]

Available Commands:
  clean       clean default grpc mock service
  gen         generate mock api
  help        Help about any command
  init        init default grpc mock service

Flags:
  -h, --help   help for rzgrpcmock
```

## Configure

### Mock Server
Мок-сервер запускается на порту 9099 (env: GRPC_ADDR)

### Mock API

API для настройки запускается на порту 9010(env: API_ADDR), доступные апи:
 * `GET /api/form` - форма для отправки мока, для дебага
 * `POST /api/mock/add` - добавить мок, структура запроса


   Подробнее в [swagger](https://editor.swagger.io/?url=https://raw.githubusercontent.com/razielsd/rzgrpcmock/master/doc/swagger.json)

###  Monitoring
API содержит несколько методов для мониторинга приложения:
* _/mertics_ - метрики приложения для prometheus
* _/health/liveness_ - k8s liveness probe
* _/health/readiness_ - k8s readiness probe