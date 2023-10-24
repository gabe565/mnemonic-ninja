FROM --platform=$BUILDPLATFORM node:20-alpine AS node-builder
WORKDIR /app

COPY package.json package-lock.json .npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY . .
RUN npm run build


FROM nginx:stable-alpine

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=node-builder /app/dist/ /usr/share/nginx/html
