# dev stage
FROM node:22-slim AS dev

WORKDIR /app

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./
RUN pnpm i --frozen-lockfile

COPY . .

EXPOSE 80

CMD ["pnpm", "dev"]


# build stage
FROM node:22-slim AS build

WORKDIR /app

COPY --chown=node:node package.json pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm i --frozen-lockfile

COPY --chown=node:node . .
RUN pnpm build


# runtime stage
FROM nginx:1.27-alpine

COPY --from=build /app/dist /var/www
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
