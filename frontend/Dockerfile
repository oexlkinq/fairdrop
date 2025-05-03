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
RUN pnpm run build-only


FROM base AS dev
CMD [ "pnpm", "dev", "--port", "8080", "--host" ]


FROM nginx:alpine AS main
COPY --from=build /src/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf