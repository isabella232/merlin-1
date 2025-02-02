# Copyright 2020 The Merlin Authors
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

# ============================================================
# Build stage 1: Build web server
# ============================================================
FROM golang:1.18-alpine as go-builder

ARG BRANCH
ARG REVISION
ARG VERSION

RUN apk update && apk add --no-cache git ca-certificates bash
RUN mkdir -p src/api

WORKDIR /src/api

# Caching dependencies
COPY api/go.mod .
COPY api/go.sum .
COPY python/batch-predictor/go.mod ../python/batch-predictor/go.mod
COPY python/batch-predictor/go.sum ../python/batch-predictor/go.sum

RUN go mod download

COPY api .
COPY python/batch-predictor ../python/batch-predictor

RUN go mod tidy
RUN go build -o bin/transformer \
    ./cmd/transformer/main.go

# ============================================================
# Build stage 2: Copy binary
# ============================================================
FROM alpine:3.12

COPY --from=go-builder /src/api/bin/transformer /usr/bin/transformer
COPY --from=go-builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip

# Adding the grpc_health_probe
ENV GRPC_HEALTH_PROBE_VERSION=v0.4.4
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV ZONEINFO=/zoneinfo.zip

ENTRYPOINT ["transformer"]
