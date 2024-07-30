FROM node:22.2.0-alpine as front
WORKDIR /app
COPY front/package*.json ./
# 安装生产依赖
RUN npm install
# 复制整个项目
COPY . .
# 构建Nuxt应用
RUN npm run generate

FROM golang:alpine as backend
ENV CGO_ENABLED 1
WORKDIR /app
COPY backend/go.mod .
COPY backend/go.sum .
RUN go mod download
COPY --from=front /app/.output/public /app/public
COPY . .
RUN apk update --no-cache && apk add --no-cache tzdata
RUN go build -ldflags="-s -w" -o /app/moments .

FROM alpine
ARG VERSION
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=backend /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
ENV VERSION $VERSION
COPY --from=backend /app/moments /app/moments
CMD ["/app/moments"]