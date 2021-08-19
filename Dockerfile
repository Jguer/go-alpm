FROM lopsided/archlinux:devel

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .

RUN pacman -Syu --overwrite=* --needed --noconfirm go && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.42.0 && \
    go mod download && \
    rm -rfv /var/cache/pacman/* /var/lib/pacman/sync/*
