#! /bin/bash

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o pdfpreview main.go

# CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o pdfpreview main.go

