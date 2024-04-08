# Nuxt 3 builder
FROM node:lts as builder

WORKDIR /app

COPY package*.json ./
COPY prisma ./prisma

# 安装依赖
RUN npm install

# 生成Prisma客户端
RUN npx prisma generate


# 构建Nuxt应用
RUN npm run build

# Nuxt 3 production
FROM node:lts

WORKDIR /app

COPY --from=builder /app/.output  /app/.output
COPY --from=builder /app/prisma  /app/prisma

RUN npm init -y
RUN npm install prisma --save-dev

# 运行数据库迁移
RUN npx prisma migrate deploy

RUN mkdir -p /app/assets/upload

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "node", ".output/server/index.mjs" ]