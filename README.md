# 极简朋友圈

![moments github action status](https://img.shields.io/github/actions/workflow/status/kingwrcy/moments/deploy.yml)
[![docker pull](https://img.shields.io/docker/pulls/kingwrcy/moments)](https://hub.docker.com/repository/docker/kingwrcy/moments)


[tg交流群](https://t.me/simple_moments)

S3兼容的对象存储配置方法(不是必须的,只有你需要把图片存储到对象存储时才需要配置,默认图片存在本地,可备份):

[Cloudflare R2配置](https://jerry.mblog.club/moments-r2-config)  

[阿里云OSS配置](https://jerry.mblog.club/moments-config-aliyun)

又拍云 不支持[使用预签名 URL 上传对象](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/userguide/PresignedUrlUploadObject.html),所以不支持又拍云.


[在线DEMO](https://m.mblog.club),欢迎体验.

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

鉴于萝卜青菜各有所爱,每个人情况不一致,特此使用配置文件配置各项特性化需求,目前支持以下环境变量配置见.

<details>

<summary>点我查看</summary>

| KEY  | 默认值 | 描述 |
| ------------- | ------------- | ------------- |
| NUXT_PUBLIC_MOMENTS_COMMENT_ENABLE  | true  | 是否开启评论 |
| NUXT_PUBLIC_MOMENTS_SHOW_COMMENT  | true  | 是否显示评论 |
| NUXT_PUBLIC_MOMENTS_COMMENT_MAX_LENGTH  | 120  | 评论最大字数 |
| NUXT_PUBLIC_MOMENTS_COMMENT_ORDER_BY  | desc  | 评论的显示顺序,desc:倒序,asc:顺序 |
| NUXT_PUBLIC_MOMENTS_TOOLBAR_ENABLE_DOUBAN  | true  | 是否显示引入豆瓣读书/视频按钮 |
| NUXT_PUBLIC_MOMENTS_TOOLBAR_ENABLE_MUSIC163  | true  | 是否显示引入网易云音乐按钮 |
| NUXT_PUBLIC_MOMENTS_TOOLBAR_ENABLE_VIDEO  | true  | 是否显示引入youtube,b站,在线视频按钮 |
| NUXT_PUBLIC_MOMENTS_MAX_LINE  | 4  | 单条发言最大行数,最大10行,超过折叠 |
| NUXT_PUBLIC_GOOGLE_RECAPTCHA_SITE_KEY  | 无  | google recaptchaV3 HTML 代码中使用此网站密钥 |
| NUXT_GOOGLE_RECAPTCHA_SECRET_KEY  | 无  | google recaptchaV3 网站和 reCAPTCHA 之间的通信密钥 |

</details>




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