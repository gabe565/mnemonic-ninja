ARG GO_VERSION=1.18
ARG NODE_VERSION=16

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION-alpine as go-builder
WORKDIR /app

RUN apk add --no-cache gcc g++

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
COPY internal/ internal/
ARG TARGETPLATFORM
# Set Golang build envs based on Docker platform string
RUN --mount=type=cache,target=/root/.cache \
    set -x \
    && case "$TARGETPLATFORM" in \
        'linux/amd64') export GOARCH=amd64 ;; \
        'linux/arm/v6') export GOARCH=arm GOARM=6 ;; \
        'linux/arm/v7') export GOARCH=arm GOARM=7 ;; \
        'linux/arm64') export GOARCH=arm64 ;; \
        *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;; \
    esac \
    && go generate \
    && go build -ldflags='-w -s'


FROM --platform=$BUILDPLATFORM node:$NODE_VERSION-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM alpine
WORKDIR /app
COPY --from=go-builder /app/mnemonic-ninja ./
COPY --from=node-builder /app/dist/ frontend/

ENV MNEMONIC_NINJA_ADDRESS ":80"
CMD ["/app/mnemonic-ninja"]
