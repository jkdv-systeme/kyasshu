FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.20 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

RUN apt install ca-certificates

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags "-X github.com/jkdv-systeme/kyasshu/internal/config.Version=$(cat .version) -X github.com/jkdv-systeme/kyasshu/internal/config.Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`" -o kyasshu .

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /opt/kyasshu
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/kyasshu ./

ENTRYPOINT ["./kyasshu", "serve"]