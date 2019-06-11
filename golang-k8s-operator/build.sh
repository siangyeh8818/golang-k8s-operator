#!/bin/bash

#mac 版的binary
go build  -o k8sclone *.go
mv k8sclone ../binary/mac

#Linux版的binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o k8sclone *.go
mv k8sclone ../binary/Linux

#Windows版的binary
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o k8sclone.exe *.go
mv k8sclone.exe ../binary/Windows