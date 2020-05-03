#!/usr/bin/env bash

set -e

goa gen github.com/kaito2/rest-api-sample/design
goa example github.com/kaito2/rest-api-sample/design

mkdir -p doc
mv gen/http/openapi.json doc
mv gen/http/openapi.yaml doc
