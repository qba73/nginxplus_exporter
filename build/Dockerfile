# syntax=docker/dockerfile:1.4
FROM golang:1.19 as base
ARG VERSION
ARG TARGETARCH

WORKDIR /go/src/github.com/qba73/nginx_exporter

FROM base as builder
COPY --link go.mod go.sum ./
RUN go mod download
COPY --link *.go ./
COPY --link collector ./collector
COPY --link client ./client
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -trimpath -a -ldflags "-s -w -X main.version=${VERSION}" -o nginx_exporter .


FROM scratch as intermediate
COPY --from=base --link /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER 1001:1001
ENTRYPOINT [ "/usr/bin/nginx_exporter" ]


FROM intermediate as container
COPY --from=builder --link /go/src/github.com/qba73/nginx_exporter/nginx_exporter /usr/bin/


FROM intermediate as goreleaser
ARG TARGETARCH
ARG TARGETVARIANT
ARG TARGETPLATFORM

LABEL org.qba73.exporter.image.build.target="${TARGETPLATFORM}"

COPY --link dist/nginx_exporter_linux_$TARGETARCH${TARGETVARIANT:+_7}*/nginx_exporter /usr/bin/
