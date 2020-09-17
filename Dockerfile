# Copyright (c) 2020 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG BASE=golang:1.15-alpine
FROM ${BASE} AS builder

ARG ALPINE_PKG_BASE="build-base git openssh-client"
ARG ALPINE_PKG_EXTRA=""

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirros we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR $GOPATH/src/github.impcloud.net/RSP-Inventory-Suite/device-llrp-go
COPY go.mod .
RUN go mod download

COPY . .

# To run tests in the build container:
#   docker build --build-arg 'MAKE=build test' .
# This is handy of you do your Docker business on a Mac
ARG MAKE='make build'
RUN $MAKE

FROM alpine

ENV APP_PORT=51992
EXPOSE $APP_PORT

COPY --from=builder /go/src/github.impcloud.net/RSP-Inventory-Suite/device-llrp-go/cmd /
COPY --from=builder /go/src/github.impcloud.net/RSP-Inventory-Suite/device-llrp-go/LICENSE /
COPY --from=builder /go/src/github.impcloud.net/RSP-Inventory-Suite/device-llrp-go/Attribution.txt /

LABEL license='SPDX-License-Identifier: Apache-2.0' \
      copyright='Copyright (c) 2020: Intel Corporation'

ENTRYPOINT ["/device-llrp","-cp=consul://edgex-core-consul:8500","-registry","-confdir=/res/docker"]
