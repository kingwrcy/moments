# Nuxt 3 builder
FROM node:22.2.0-alpine as builder

RUN corepack enable && corepack prepare pnpm@latest --activate
ENV PNPM_HOME=/usr/local/bin
RUN pnpm add --global prisma

ARG VERSION

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

# 安装生产依赖
RUN pnpm install
ENV NODE_ENV=production

# 复制整个项目
COPY . .

# 生成Prisma客户端
RUN npx prisma generate

RUN echo $VERSION > /app/version

# 构建Nuxt应用
RUN pnpm run build

# Nuxt 3 production
FROM node:22-alpine
RUN corepack enable && corepack prepare pnpm@latest --activate
ENV PNPM_HOME=/usr/local/bin
RUN pnpm add --global prisma

WORKDIR /app

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/data/db.sqlite"
ENV UPLOAD_DIR="/app/data/upload"
ENV CONFIG_FILE="/app/data/config.json"
ENV MOMENTS_VERSION=$VERSION

RUN mkdir -p /app/data/upload

COPY --from=builder /app/.output /app/.output
COPY --from=builder /app/prisma /app/prisma
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/version /app/version
COPY --from=builder /app/config.json .

RUN chmod +x /app/start.sh

EXPOSE 3000

CMD /app/start.sh