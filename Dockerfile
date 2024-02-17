FROM node:21.6-slim as builder-frontend

WORKDIR /build/frontend
COPY ./frontend .

ARG BASE_ENDPOINT
ARG NEXT_PUBLIC_FIREBASE_API_KEY
ARG NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN
ARG NEXT_PUBLIC_FIREBASE_PROJECT_ID
ARG NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET
ARG NEXT_PUBLIC_FIREBASE_MESSAGE_SENDER_ID
ARG NEXT_PUBLIC_FIREBASE_APP_ID
ARG NEXT_PUBLIC_ENVIRONMENT
ARG NEXT_PUBLIC_USER_AUTH

ENV NEXT_PUBLIC_BASE_ENDPOINT=$BASE_ENDPOINT \
    NEXT_PUBLIC_FIREBASE_API_KEY=$NEXT_PUBLIC_FIREBASE_API_KEY \
    NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=$NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN \
    NEXT_PUBLIC_FIREBASE_PROJECT_ID=$NEXT_PUBLIC_FIREBASE_PROJECT_ID \
    NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET=$NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET \
    NEXT_PUBLIC_FIREBASE_MESSAGE_SENDER_ID=$NEXT_PUBLIC_FIREBASE_MESSAGE_SENDER_ID \
    NEXT_PUBLIC_FIREBASE_APP_ID=$NEXT_PUBLIC_FIREBASE_APP_ID \
    NEXT_PUBLIC_ENVIRONMENT=$NEXT_PUBLIC_ENVIRONMENT \
    NEXT_PUBLIC_USER_AUTH=$NEXT_PUBLIC_USER_AUTH

RUN npm install && \
    npm run build
    
FROM golang:1.22 as builder-backend
WORKDIR /build/backend

COPY ./backend .

RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o ikigai.app -ldflags="-s -w" cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy
COPY --from=builder-backend /build/backend/ikigai.app /
COPY --from=builder-frontend /build/frontend/out /static
EXPOSE 3000
CMD ["/ikigai.app"]