# build stage
FROM node:slim AS frontend-base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM frontend-base AS build-frontend
WORKDIR /src

COPY frontend .
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build-only


FROM golang:alpine AS build-backend
WORKDIR /src

RUN apk add gcc musl-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY --from=build-frontend /src/dist /src/frontend/dist
COPY internal ./internal
COPY main.go .
COPY frontend/embed.go frontend

ENV GOCACHE=/gocache
RUN --mount=type=cache,target="/gocache" CGO_ENABLED=1 go build -v -o /fairdrop

# #2 TODO: вернуть, когда появятся тесты
# Run the tests in the container
# FROM build-backend AS run-test-stage
# RUN go test -v


# build release stage
FROM alpine AS build-release-stage
WORKDIR /

COPY --from=build-backend /fairdrop /fairdrop

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/fairdrop"]
