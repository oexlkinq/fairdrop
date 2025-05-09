FROM golang:1.24.3-alpine AS build
WORKDIR /src

RUN apk add gcc musl-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY internal ./internal
COPY main.go .

ENV GOCACHE=/gocache
RUN --mount=type=cache,target="${GOCACHE}" CGO_ENABLED=1 go build -v -o /fairdrop

# #2 TODO: вернуть, когда появятся тесты
# Run the tests in the container
# FROM build AS run-test-stage
# RUN go test -v


FROM alpine AS main
WORKDIR /

COPY --from=build /fairdrop /fairdrop

ENV GIN_MODE=release
ENV PORT=8080

ENTRYPOINT ["/fairdrop"]
