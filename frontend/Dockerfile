FROM node:slim AS corepack

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM corepack AS deps
WORKDIR /src
COPY package.json pnpm-lock.yaml ./

RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
    pnpm install --frozen-lockfile


FROM deps AS base
COPY . .


FROM base AS build
ARG BASE_PATH
RUN pnpm run build-only --base ${BASE_PATH}


FROM base AS dev
ARG BASE_PATH
ENV BASE_PATH=${BASE_PATH}
CMD [ "sh", "-c", "pnpm dev --port 8080 --host --base ${BASE_PATH}" ]


FROM nginx:alpine AS main
COPY --from=build /src/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf