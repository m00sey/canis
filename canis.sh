#!/bin/bash

PWD=$(pwd)
export CANIS_ROOT=$PWD

export CGO_CFLAGS=-I"${CANIS_ROOT}"/include