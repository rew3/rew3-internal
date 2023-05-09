#!/bin/bash

URL="http://142.93.6.93:8081/artifactory/ext-snapshot-local/go"
export GOPROXY=$(go env GOPROXY | grep -q -F $URL || echo "$(go env GOPROXY);"+$URL)
export GONOSUMDB=*
go mod tidy
go mod vendor
curl -u rew3:rew3 -X PUT $URL/rew3.com/app-core/v1.0.0 \
-T ./go.sum -T ./go.mod -T ./vendor/modules.txt -T ./vendor/modules/*
