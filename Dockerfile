#checkov:skip=CKV_DOCKER_2:Ensure that HEALTHCHECK instructions have been added to container images
FROM golang:1.24-alpine3.22 AS build

WORKDIR /go

COPY . .

RUN go build -o /kar cmd/kar/main.go

FROM scratch

COPY --from=build /kar /opt/kar

USER 10001:10001

ENTRYPOINT [ "/opt/kar"]
