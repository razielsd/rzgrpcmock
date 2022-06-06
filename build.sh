#!/usr/bin/env bash

ROOT=$(PWD)
TEMPLATE_ROOT="${ROOT}/template"
BUILDER_ROOT="${ROOT}/builder"
MOCK_ROOT="${ROOT}/mock"
GOROOT=`go env | grep -F 'GOPATH' | cut -d"=" -f2 | cut -d "\"" -f2`

show-env()
{
  echo "GOROOT: ${GOROOT}"
  echo "ROOT: ${ROOT}"
  echo "TEMPLATE_ROOT: ${TEMPLATE_ROOT}"
  echo "BUILDER_ROOT: ${BUILDER_ROOT}"
  echo "MOCK_ROOT: ${MOCK_ROOT}"
}

template-init()
{
  echo "INIT ${MOCK_ROOT}"
  cp -R $TEMPLATE_ROOT $MOCK_ROOT
  rm -rf $MOCK_ROOT/vendor
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
  cd $MOCK_ROOT && go get "$PKG"

  #generate mock
  PKG_PATH=`find-package $PKG`
  if [ ! -d "${PKG_PATH}" ]; then
    echo "Package not found: ${PKG_PATH}"
    exit 1
  fi
  find ${PKG_PATH} -name "*grpc.pb.go" | while read -r i
  do
    cd $BUILDER_ROOT && go run . $i ${MOCK_ROOT}/internal/generated ${PKG_NAME}
  done
}

run-server() {
  cd ${MOCK_ROOT} && go run .
}

find-package() {
  PKG=$1
  PKG_NAME=`echo $PKG | cut -d\@ -f1`
  PKG_VERSION=`echo $PKG | cut -d\@ -f2`
  LIMIT=`echo ${PKG} | grep -o '/' | wc -l`
  FULL_PATH="${GOROOT}/pkg/mod/${PKG}"
  if [ -d "${FULL_PATH}" ]; then
      echo "${FULL_PATH}"
      return
  fi
  i=1
  while [ $i -le "$LIMIT" ]
  do
    PKG_PATH=`echo $PKG_NAME | sed 's/\//@'$PKG_VERSION'\//'${i}`
    FULL_PATH="${GOROOT}/pkg/mod/${PKG_PATH}"
    if [ -d "${FULL_PATH}" ]; then
      echo "${FULL_PATH}"
      return
    fi
    i=$(($i+1))
  done
}

show-find-package() {
  PKG=$1
  echo "PACKAGE: ${PKG}"
  PKG_NAME=`echo $PKG | cut -d\@ -f1`
  PKG_VERSION=`echo $PKG | cut -d\@ -f2`
  echo "NAME: ${PKG_NAME}"
  echo "VERSION: ${PKG_VERSION}"
  PKG_PATH=`find-package $PKG`
  echo "PATH: ${PKG_PATH}"
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
  find-pkg)
    show-find-package $2
  ;;
esac
cd $ROOT