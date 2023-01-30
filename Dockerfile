FROM --platform=$BUILDPLATFORM node:18-alpine AS node-builder
WORKDIR /app

COPY package.json package-lock.json .npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY . .
RUN npm run build


FROM nginx:stable-alpine

COPY --from=node-builder /app/dist/ /usr/share/nginx/html
