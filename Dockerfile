ARG GO_VERSION=1.17
ARG NODE_VERSION=16

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION-alpine as go-builder
WORKDIR /app

RUN apk add --no-cache gcc g++

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY internal/ internal/
ARG TARGETPLATFORM
# Set Golang build envs based on Docker platform string
RUN set -x \
    && case "$TARGETPLATFORM" in \
        'linux/amd64') export GOARCH=amd64 ;; \
        'linux/arm/v6') export GOARCH=arm GOARM=6 ;; \
        'linux/arm/v7') export GOARCH=arm GOARM=7 ;; \
        'linux/arm64') export GOARCH=arm64 ;; \
        *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;; \
    esac \
    && mkdir -p frontend/dist \
    && echo stub >frontend/dist/gitkeep \
    && go generate \
    && go build -ldflags="-w -s"


FROM --platform=$BUILDPLATFORM node:$NODE_VERSION-alpine AS node-builder
WORKDIR /app

RUN apk add --no-cache g++ make python3

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM alpine
WORKDIR /app
RUN apk add --no-cache lame
COPY --from=go-builder /app/mnemonic-ninja ./
COPY --from=node-builder /app/dist/ dist/

ENV MNEMONIC_NINJA_ADDRESS ":80"
ENV MNEMONIC_NINJA_STATIC "dist"
CMD ["./mnemonic-ninja"]
