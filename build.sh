#!/usr/bin/env bash

ROOT=$(PWD)
TEMPLATE_ROOT="${ROOT}/template"
BUILDER_ROOT="${ROOT}/builder"
MOCK_ROOT="${ROOT}/mock"

HashFunc="shasum"
hashCheck=`which ${HashFunc}`
if [ -z ${hashCheck} ]; then
  HashFunc="sha1sum"
fi

show-env()
{
  echo "ROOT: ${ROOT}"
  echo "TEMPLATE_ROOT: ${TEMPLATE_ROOT}"
  echo "BUILDER_ROOT: ${BUILDER_ROOT}"
  echo "MOCK_ROOT: ${MOCK_ROOT}"
  echo "HashFunc: ${HashFunc}"
}

template-init()
{
  echo "INIT ${MOCK_ROOT}"
  cp -R $TEMPLATE_ROOT $MOCK_ROOT
}

clean-mock()
{
  echo "CLEAN ${MOCK_ROOT}"
  if [[ -d "$MOCK_ROOT" ]]
  then
    rm -rf $MOCK_ROOT
  fi
  template-init
}

add-module()
{
  #import module
  PKG=$1
  echo "ADD ${PKG}"

  PKG_NAME=`echo $PKG | cut -d\@ -f1`
  FILENAME=`echo $PKG_NAME | ${HashFunc} | cut -d" " -f1`
  FAKE_USE_FILE="${MOCK_ROOT}/internal/generated/${FILENAME}.go"
  echo "package generated" > $FAKE_USE_FILE
  echo "import \"${PKG_NAME}\"" >> $FAKE_USE_FILE
  cd $MOCK_ROOT && go get "$PKG" && go mod vendor
  rm -rf $FAKE_USE_FILE

  #generate mock

  find ${MOCK_ROOT}/vendor/${PKG_NAME} -name "*grpc.pb.go" | while read -r i
  do
    cd $BUILDER_ROOT && go run . $i ${MOCK_ROOT}/internal/generated ${PKG_NAME}
  done

  cd ${MOCK_ROOT} && go mod vendor
}

run-server() {
  cd ${MOCK_ROOT} && go run .
}

MODE=$1

case $MODE in
  clean)
    clean-mock
  ;;
  init)
    template-init
  ;;
  add)
    add-module $2
  ;;
  env)
    show-env
  ;;
  run)
    run-server
  ;;
esac
cd $ROOT

# find . -name "*grpc.pb.go"




#print-args "eval set -- \"$opts\"" "$@"