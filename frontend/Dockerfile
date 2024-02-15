FROM nginx:1.25 as builder-nginx
ARG TIME_ZONE="Asia/Tokyo"

RUN mkdir -p /opt/var/cache/nginx && \
    cp -a --parents /usr/lib/nginx /opt && \
    cp -a --parents /usr/share/nginx /opt && \
    cp -a --parents /var/log/nginx /opt && \
    cp -aL --parents /var/run /opt && \
    cp -a --parents /etc/nginx /opt && \
    cp -a --parents /etc/passwd /opt && \
    cp -a --parents /etc/group /opt && \
    cp -a --parents /usr/sbin/nginx /opt && \
    cp -a --parents /usr/sbin/nginx-debug /opt && \
    cp -R -a --parents /lib/x86_64-linux-gnu /opt && \
    cp -a --parents /usr/lib/x86_64-linux-gnu/libssl.so.* /opt && \
    cp -a --parents /usr/lib/x86_64-linux-gnu/libcrypto.so.* /opt && \
    cp /usr/share/zoneinfo/${TIME_ZONE:-ROC} /opt/etc/localtime && \
    rm -rf /opt/etc/nginx/conf.d/default.conf && \
    rm -rf /opt/usr/share/nginx/html

FROM node:21.6-slim as builder-app

WORKDIR /build
COPY . .

ARG BASE_ENDPOINT
ARG NEXT_PUBLIC_FIREBASE_API_KEY
ARG NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN
ARG NEXT_PUBLIC_FIREBASE_PROJECT_ID
ARG NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET
ARG NEXT_PUBLIC_FIREBASE_MESSAGE_SENDER_ID
ARG NEXT_PUBLIC_FIREBASE_APP_ID
ARG NEXT_PUBLIC_ENVIRONMENT
ARG NEXT_PUBLIC_USER_AUTH

RUN npm install && \
    npm run build

FROM gcr.io/distroless/base-debian12
WORKDIR /app
ENV NODE_ENV=production

COPY --from=builder-nginx /opt /
COPY --from=builder-app /build/out/ /usr/share/nginx/html/
COPY --from=builder-app /build/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]