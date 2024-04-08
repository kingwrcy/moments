# Nuxt 3 builder
FROM node:lts-alpine as builder

WORKDIR /app

COPY package*.json ./

# 安装生产依赖
RUN npm ci 

# 复制整个项目
COPY . .

# 生成Prisma客户端
RUN npx prisma generate

ENV NODE_ENV=production

# 构建Nuxt应用
RUN npm run build

# Nuxt 3 production
FROM node:lts-alpine

WORKDIR /app

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/data/db.sqlite"

RUN mkdir -p /app/data/upload

COPY --from=builder /app/.output /app/.output
COPY --from=builder /app/prisma /app/prisma

RUN npm init -y
RUN npm install -g prisma
# 运行数据库迁移
RUN npx prisma migrate deploy

EXPOSE 3000

CMD [ "node", ".output/server/index.mjs" ]