# RzGrpcMock

## Using
Для настройки мок-сервера используйте `build.sh`, параметры:
 * `./build.sh init` - инициализации мок-сервера
 * `./build.sh clean` - привести мок-сервер к начальному состоянию
 * `./build.sh add <go module>` - как в `go get ...`
 * `./build.sh run` - запуск сервера

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