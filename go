#!/bin/bash
set -euo pipefail

SCRIPT_DIR=$(
  cd $(dirname $0)
  pwd -P
)

usage() {
  # print usage information
  echo "Commands:"
  echo
  echo "help     Print this help message."
  echo "build    Build the service."
  echo "test     Run tests."
  echo "deploy   Run deployments."
  echo "  terraform     Apply terraform."
  echo "  k8s           Deploy kubernetes resources."
  echo "  service       Deploy service."
}

build() {
  # trigger build process
  echo
}

test() {
  # run test suite
  echo
}

deploy() {
  case "${2}" in
  "k8s")
    echo "Deploying kubernetes resources..."
    kubectl apply -f $SCRIPT_DIR/k8s-manifests
    ;;
  *)
    echo "Unknown deployment target: ${2}"
    usage
    exit 1
    ;;
  esac
}

nuke() {
  case "${2}" in
  "k8s")
    echo "Nuking kubernetes resources..."
    kubectl delete -f $SCRIPT_DIR/k8s-manifests
    ;;
  *)
    echo "Unknown nuke target: ${2}"
    usage
    exit 1
    ;;
  esac
}

if [[ $# -lt 1 ]]; then
  usage
  exit 1
fi

TARGET=$1
case $TARGET in
"help")
  usage
  ;;
"build")
  build $@
  ;;
"test")
  test $@
  ;;
"deploy")
  deploy $@
  ;;
*)
  echo "Unknown command '${TARGET}'"
  usage
  exit 1
  ;;
esac
