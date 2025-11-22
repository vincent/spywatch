#!/bin/sh

~/go/bin/oapi-codegen -package=changedetection  ./changedetection.json > ./changedetection.go
