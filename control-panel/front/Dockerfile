FROM node:16.17-alpine3.16 as builder

WORKDIR /app

COPY . .

RUN yarn install && \
    yarn build

FROM nginx:1.23-alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /app/build .

EXPOSE 8080

ENTRYPOINT ["nginx", "-g", "daemon off;"]
