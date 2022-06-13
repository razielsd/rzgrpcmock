# RzGrpcMock

[![codecov](https://codecov.io/gh/razielsd/rzgrpcmock/branch/master/graph/badge.svg)](https://codecov.io/gh/razielsd/rzgrpcmock)
[![Go Report Card](https://goreportcard.com/badge/github.com/razielsd/rzgrpcmock)](https://goreportcard.com/report/github.com/razielsd/rzgrpcmock)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/razielsd/rzgrpcmock)

## Using
```
RzGrpcMock service cli
Available Commands:
  init        init default grpc mock service
  clean       clean default grpc mock service
  gen         generate mock api
```
Создаем мок-сервис
```
rzgrpcmock init <path/to/new/service>
```
Добавляем в него апи для сервиса моков, можно добавить несколько пакетов, эту команду можно выполнять несколько раз:
```
rzgrpcmock gen <path/to/new/service> <package@version>
```
Запускаем сервис:
```
cd <path/to/new/service> && go run .
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