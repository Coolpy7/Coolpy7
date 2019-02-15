#!/usr/bin/env bash
go build -ldflags "-pluginpath=plugin/cp7-auth-$(date +%s)" -buildmode=plugin -o ./auth/cp7_auth.so ./auth/cp7_auth.go

go build -ldflags "-pluginpath=plugin/cp7-pubs-$(date +%s)" -buildmode=plugin -o ./pubs/cp7_pubs.so ./pubs/cp7_pubs.go

go build -ldflags "-pluginpath=plugin/cp7-term-$(date +%s)" -buildmode=plugin -o ./term/cp7_term.so ./term/cp7_term.go