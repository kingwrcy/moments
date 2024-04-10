# 极简朋友圈

- 支持匿名评论/点赞
- 支持引入网易云音乐,b站视频
- 支持自定义头图,个人头像,网站标题等
- 支持上传图片到S3兼容的云存储,支持本地存储
- 适配手机
- 支持暗黑模式
- 数据库采用sqlite,可随时备份

## Docker启动
Docker启动看[这里](https://github.com/kingwrcy/moments/blob/master/docker-start.sh)

## Docker Compose启动
Docker Compose启动看[这里](https://github.com/kingwrcy/moments/blob/master/docker-compose.yml)

## 源码编译启动

首先设置环境变量:

```
-- sqlite数据库位置
DATABASE_URL="file:/app/data/db.sqlite" 
-- 本地上传的文件目录
UPLOAD_DIR="/app/data/upload"
```

执行命令

```
-- 安装依赖
npm install
-- 脚本迁移
npx prisma migrate dev
-- 执行构建
npm run build
-- 预览
npm run preview
```

[在线DEMO](https://m.mblog.club),欢迎体验.