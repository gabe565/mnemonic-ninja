FROM alpine
LABEL org.opencontainers.image.source="https://github.com/gabe565/mnemonic-ninja"
WORKDIR /app

RUN apk add --no-cache tzdata

COPY mnemonic-ninja ./

ARG USERNAME=mnemonic-ninja
ARG UID=1000
ARG GID=$UID
RUN addgroup -g "$GID" "$USERNAME" \
    && adduser -S -u "$UID" -G "$USERNAME" "$USERNAME"
USER $UID

ENV MNEMONIC_NINJA_ADDRESS ":80"
CMD ["/app/mnemonic-ninja"]
