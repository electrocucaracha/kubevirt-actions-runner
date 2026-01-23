#checkov:skip=CKV_DOCKER_2:Ensure that HEALTHCHECK instructions have been added to container images
FROM golang:1.25-alpine3.22 AS build

WORKDIR /app

RUN apk add --no-cache git=2.49.1-r0

COPY . .

SHELL ["/bin/sh", "-o", "pipefail", "-c"]
RUN go build \
    -buildvcs=true \
    -ldflags="-X 'main.gitCommit=$(git rev-parse HEAD)' \
    -X 'main.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')' \
    -X 'main.gitTreeModified=$(git status --porcelain | wc -l)'" \
    -o /go/bin/ ./...

FROM scratch

COPY --from=build /go/bin/kar /opt/kar

USER 10001:10001

ENTRYPOINT [ "/opt/kar"]
