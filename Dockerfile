# Nuxt 3 builder
FROM node:20.12.1-alpine as builder

ARG VERSION

WORKDIR /app

COPY package*.json ./

# 安装生产依赖
RUN npm ci 

# 复制整个项目
COPY . .

# 生成Prisma客户端
RUN npx prisma generate

RUN echo $VERSION > /app/version

ENV NODE_ENV=production

# 构建Nuxt应用
RUN npm run build

# Nuxt 3 production
FROM node:20.12.1-alpine

WORKDIR /app

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/data/db.sqlite"
ENV UPLOAD_DIR="/app/data/upload"
ENV CONFIG_FILE="/app/config.json"
ENV MOMENTS_VERSION=$VERSION

RUN mkdir -p /app/data/upload

COPY --from=builder /app/.output /app/.output
COPY --from=builder /app/prisma /app/prisma
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/version /app/version
COPY --from=builder /app/config.json /app/config.json

RUN npm init -y
RUN npm install -g prisma
RUN chmod +x /app/start.sh

EXPOSE 3000

CMD /app/start.sh