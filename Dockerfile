FROM --platform=$BUILDPLATFORM node:22-alpine AS node-builder
WORKDIR /app

RUN corepack enable

COPY package.json pnpm-*.yaml .npmrc ./
RUN --mount=type=cache,id=pnpm,target=/root/.cache \
  pnpm install --prod --frozen-lockfile

COPY . .
RUN --mount=type=cache,id=pnpm,target=/root/.cache \
  pnpm run build


FROM nginx:stable-alpine

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=node-builder /app/dist/ /usr/share/nginx/html
