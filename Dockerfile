# build stage
FROM golang:alpine AS build-stage

WORKDIR /src

RUN apk add gcc musl-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN --mount=source=internal,dst=internal \
    --mount=source=main.go,dst=main.go \
    CGO_ENABLED=1 go build -v -o /fairdrop

# #2 TODO: вернуть, когда появятся тесты
# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v


# build release stage
FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /fairdrop /fairdrop

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/fairdrop"]
