# 极简朋友圈

[![docker pull](https://img.shields.io/badge/moments-更新记录-blue)](https://github.com/kingwrcy/moments/blob/master/release.md)
[![docker pull](https://img.shields.io/badge/moments-常见问题-blue)](https://github.com/kingwrcy/moments/blob/master/q&a.md)
![moments github action status](https://img.shields.io/github/actions/workflow/status/kingwrcy/moments/deploy.yml)
[![docker pull](https://img.shields.io/docker/pulls/kingwrcy/moments)](https://hub.docker.com/repository/docker/kingwrcy/moments)
[![docker pull](https://img.shields.io/badge/Telegram-group-blue)](https://t.me/simple_moments)
[![docker pull](https://img.shields.io/badge/1panel-本地安装-blue)](https://ono.ee/?p=1713750155422)



## 贡献者

<a href="https://github.com/kingwrcy/moments/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=kingwrcy/moments" />
</a>


## S3配置教程

[![docker pull](https://img.shields.io/badge/CF-R2配置-blue)](https://jerry.mblog.club/moments-r2-config)
[![docker pull](https://img.shields.io/badge/阿里云-OSS配置-blue)](https://jerry.mblog.club/moments-config-aliyun)


S3兼容的对象存储配置方法(不是必须的,只有你需要把图片存储到对象存储时才需要配置,默认图片存在本地,可备份):


又拍云 不支持[使用预签名 URL 上传对象](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/userguide/PresignedUrlUploadObject.html),所以不支持又拍云.


[在线DEMO](https://m.mblog.club),欢迎体验.目前不支持多用户,多用户版本已由[RandallAnjie](https://github.com/RandallAnjie)自行实现了,有需要的可以去[看看](https://moments.randallanjie.com/)

[1panel本地安装](https://ono.ee/?p=1713750155422),感谢[包子叔](https://ono.ee)提供的教程.

- 支持匿名评论/点赞
- 支持引入网易云音乐,b站视频,插入链接等
- 支持自定义头图,个人头像,网站标题等
- 支持上传图片到S3兼容的云存储,支持本地存储
- 适配手机
- 支持暗黑模式
- 数据库采用sqlite,可随时备份
- 支持引入豆瓣读书/豆瓣电影,样式来源于[这里](https://github.com/TankNee/hexo-douban-card/blob/master/templates/assets/style.css)

有其他需求欢迎提issues.

默认用户名密码:`admin/a123456`,登录进去后后台可以自己修改密码.

## 自定义其他配置

打开[配置管理器](https://mconfig.mblog.club),配置好后点击一键复制配置,然后进入moments后台拉到最底下,导入进去,保存即可生效.


## 使用google recaptchaV3(可选)

自行去[google recaptchaV3 admin console](https://www.google.com/recaptcha/admin/create)开通,每月100万次免费调用.
开通成功后复制网站密钥和通信密钥,填入上方的环境变量对应的key里面.


## Docker启动
Docker首次启动看[这里](https://github.com/kingwrcy/moments/blob/master/docker-start.sh)

Docker更新看[这里](https://github.com/kingwrcy/moments/blob/master/docker-update.sh)

## Docker Compose启动
Docker Compose启动看[这里](https://github.com/kingwrcy/moments/blob/master/docker-compose.yml)

## 源码编译启动

首先设置环境变量:

```
-- sqlite数据库位置
DATABASE_URL="file:/app/data/db.sqlite" 
-- 本地上传的文件目录
UPLOAD_DIR="/app/data/upload"
-- 配置文件目录(可以复制项目根目录的)
CONFIG_FILE=/app/data/config.json
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

## 编辑SQLITE数据库

```
# 容器内部执行
npx prisma studio
```

执行上面的命令会在容器内部暴露一个5555端口,暴露到主机后可以通过 `http://容器IP:5555` 访问数据库,直接修改/删除/新增数据.


## 配置S3(可选)

由于使用了[使用预签名 URL 上传对象](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/userguide/PresignedUrlUploadObject.html)方案来上传图片到S3,简单来说就是前端直接上传文件到S3,不经过服务端.

不支持这个预签名技术的S3无法上传,据我所知,号称兼容S3的云存储大部分都支持这个特性.比如腾讯云,七牛云,阿里云等.

另外,要求在S3上配置跨域,配置你当前的域名能够访问S3的资源,不配置的话,是无法使用的.

比如我这里使用的是[缤纷云](https://www.bitiful.com/),配置如下:

![缤纷云](https://yoyo.s3.bitiful.net/2024/04/12/6618b41d6b65c.png?fmt=webp)

## 重置密码

目前没有别的办法重置密码,只有修改数据库.见[编辑SQLITE数据库](https://github.com/kingwrcy/moments?tab=readme-ov-file#%E7%BC%96%E8%BE%91sqlite%E6%95%B0%E6%8D%AE%E5%BA%93).

或者任何能正常打开SQLITE数据库的工具都行,数据库见前面的环境变量部分.

打开[bcrypt-generator](https://bcrypt-generator.com/)或者其他类似的bcrypt在线加密的网站,加密你的密码.

复制加密后的密码,编辑数据库,更新User表pwd字段,更新完后记得关掉5555端口的映射,执行`npx prisma studio`命令停止5555端口.

## 打赏

如果你觉得这个项目对你有帮助,可以对我打赏,感谢!

![1713695645770.png](https://yoyo.s3.bitiful.net/2024/04/21/6624eb9a4fd18.png)