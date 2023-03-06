FROM lopsided/archlinux:devel
LABEL maintainer="Jguer,joaogg3 at google mail"

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .

RUN pacman -Syu --overwrite=* --needed --noconfirm go && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.51.2 && \
    go mod download && \
    rm -rfv /var/cache/pacman/* /var/lib/pacman/sync/*
