#!/usr/bin/env bash
go build -ldflags "-pluginpath=plugin/cp7-auth-$(date +%s)" -buildmode=plugin -o ../../plugin/cp7_auth_$(date +%s).so ./auth/cp7_auth.go

go build -ldflags "-pluginpath=plugin/cp7-pubs-$(date +%s)" -buildmode=plugin -o ../../plugin/cp7_pubs_$(date +%s).so ./pubs/cp7_pubs.go

go build -ldflags "-pluginpath=plugin/cp7-subs-$(date +%s)" -buildmode=plugin -o ../../plugin/cp7_subs_$(date +%s).so ./subs/cp7_subs.go

go build -ldflags "-pluginpath=plugin/cp7-unsubs-$(date +%s)" -buildmode=plugin -o ../../plugin/cp7_unsubs_$(date +%s).so ./unsubs/cp7_unsubs.go

go build -ldflags "-pluginpath=plugin/cp7-term-$(date +%s)" -buildmode=plugin -o ../../plugin/cp7_term_$(date +%s).so ./term/cp7_term.go