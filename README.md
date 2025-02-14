# Moments - 极简朋友圈

[![release](https://img.shields.io/badge/release-更新记录-blue)](https://github.com/kingwrcy/moments/releases)
[![docker-release-status](https://img.shields.io/github/actions/workflow/status/kingwrcy/moments/docker-image-release.yml)](https://github.com/kingwrcy/moments/actions/workflows/docker-image-release.yml)
[![docker-pull](https://img.shields.io/docker/pulls/kingwrcy/moments)](https://hub.docker.com/repository/docker/kingwrcy/moments)
[![telegram-group](https://img.shields.io/badge/Telegram-group-blue)](https://t.me/simple_moments)
[![discussion](https://img.shields.io/badge/moments-论坛-blue)](https://discussion.mblog.club)

从 v0.2.1 开始，使用 Golang 重写了服务端，目前已经基本实现了 v0.2.0 版本的功能，同时软件包体积也更小了。

注意：如果你在找 v0.2.0 版本，请 [点击这里](https://github.com/kingwrcy/moments/blob/master/README.md)。

# 功能列表

## 用户

- 默认用户名和密码是 admin/a123456，登录后可以在后台修改
- 多用户模式，可以在后台打开或关闭用户注册功能

## Memo

- 支持设置标签
- 支持上传图片，可以上传到服务器，也可以在后台开启上传到 S3
- 支持生成缩略图，但是目前只支持直接上传到服务器时生成缩略图，将在后续版本中支持上传到 S3 时生成缩略图
- 支持 Markdown 语法，但是目前只适配了常用的几个标签，将在后续版本中支持更多的标签
- 支持点赞
- 支持评论，可以在后台打开或关闭评论功能

## 其他

- 支持回到顶部按钮，PC 端和手机端都有

更多说明可以[点击这里查看](https://discussion.mblog.club/post/pto2hqoFzDKzZMpvoPZKYuP)。

# 使用教程

## 环境变量

Moments 支持以下环境变量，可按需修改进行配置：

| 变量名            | 解释                  | 默认值                                                 |
| ----------------- | --------------------- | ------------------------------------------------------ |
| PORT              | 监听端口              | 3000                                                   |
| JWT_KEY           | JWT 密钥              | 不填写则每次启动时随机生成，每次重启后需要重新登录     |
| DB                | sqlite 数据库存放目录 | 工作目录下的 db.sqlite，/app/data/db.sqlite            |
| UPLOAD_DIR        | 上传文件本地目录      | 工作目录下的 upload，/app/data/upload                  |
| LOG_LEVEL         | 日志级别              | info，可选 debug                                       |
| ENABLE_SWAGGER    | 是否启用 swagger 文档 | false，可选 true，启用后可访问路径 /swagger/index.html |
| ENABLE_SQL_OUTPUT | 是否启用 SQL 调试日志 | false                                                  |

除了直接配置环境变量外，还可以可选地在工作目录下创建 `.env` 文件，Moments 在启动时也会读取该文件来决定最终生效的配置，例如：

```
JWT_KEY=
LOG_LEVEL=info
```

## 使用 Docker

注意：需要将以下示例中的 `$JWT_KEY` 替换为你的真实 `JWT_KEY`，生成方式可[参考这里](#jwt_key-生成)。

使用命令行启动容器：

```bash
docker run -d \
  -e PORT=3000 \
  -e JWT_KEY=$JWT_KEY \
  -p 3000:3000 \
  -v /var/moments:/app/data \
  --name moments \
  kingwrcy/moments:latest
```

注意：

- 容器默认的工作目录是 `/app/data`，所以示例中将 `/app/data` 目录挂载到主机的 `/var/moments` 目录，实现持久化数据的目的，可以按需修改
- 示例中使用的是 `latest` 标签来拉取最新的发布版；也可以通过 `dev` 标签来拉取最新的开发版，通常开发版会包含最新的功能和问题修复，但是相对于发布版的稳定较差
- 无论使用 `latest` 版本还是 `dev` 版本，我们都 _强烈建议将容器的数据持久化，并（定期或在升级前）备份到安全的介质中_

也可以使用 Docker Compose:

```yaml
services:
  moments:
    image: kingwrcy/moments:latest
    container_name: moments
    restart: always
    environment:
      PORT: 3000
      JWT_KEY: $JWT_KEY
    ports:
      - 3000:3000
    volumes:
      - /var/moments:/app/data # 持久化数据到主机的 /var/moments 目录，可以按需修改
```

## 使用可执行文件

首先在 [Release](https://github.com/kingwrcy/moments/releases) 列表根据自己的平台下载最新版本的可执行文件。

例如以下是用于 `windows-amd64` 的文件：

| 文件名                                       | 说明                                           |
| -------------------------------------------- | ---------------------------------------------- |
| moments-windows-amd64-x.x.x.exe.zip          | 包含可执行文件的压缩包，解压后可得到可执行文件 |
| moments-windows-amd64-x.x.x.exe-checksum.txt | 包含对应可执行文件的 `MD5` 校验码              |

下载并解压完成后，可以可选地检查可执行文件的校验码是否匹配，然后通过环境变量或 `.env` 文件进行配置，最后直接打开可执行文件即可。

## JWT_KEY 生成

### 使用 OpenSSL

执行以下命令生成 (仅测试了 Linux)：

```bash
openssl rand -hex 32
```

### 使用 SHA256

执行以下命令生成 (仅测试了 Linux)：

```bash
echo $RANDOM | sha256sum
```

### 在线生成

打开 [https://tool.lu/uuid](https://tool.lu/uuid) 生成不带 `-` 的 `UUID` 作为 `JWT_KEY`。

## 配置前端地址

项目启动后，请在系统设置中设置前端地址并保存，如：`https://www.example.com`。如果没有配置域名，也可以使用`<protocol>://<ip>:<port>`的形式。

如果不配置前端地址，RSS和邮件通知等功能将无法正确生成原文链接。

# 开发

## 环境

配置以下开发环境：

- Go 1.22.5（或更高）
- NodeJS 18（或更高）
- PNPM 包管理工具

对于 `VSCode`，我们推荐安装以下插件：

- eamodio.gitlens
- esbenp.prettier-vscode
- dbaeumer.vscode-eslint
- Nuxtr.nuxt-vscode-extentions
- golang.go
- qwtel.sqlite-viewer

另外，也可以直接通过 `VSCode` 的 [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers) 来启动配置好的开发环境。

## 启动

对于安装了 `make` 命令的用户，可以使用以下命令启动项目：

后端：

```bash
# 进入项目根目录
cd moments

# 编译并启动后端，注意此时的工作目录是项目根目录 moments
make backend-dev
```

单独创建一个终端，启动前端：

```bash
# 进入项目根目录
cd moments

# 安装前端依赖
make frontend-install

# 启动前端
make frontend-dev
```

如果没有 `make` 命令，也可以使用以下命令启动项目：

后端：

```bash
# 进入后端目录
cd moments/backend

# 编译后端
go build -ldflags="-X main.version=local -X main.commitId=local" -o ./dist/moments

# 启动后端，注意此时的工作目录是后端目录 moments/backend
./dist/moments
```

单独创建一个终端，启动前端：

```bash
# 进入前端目录
cd moments/front

# 安装前端依赖
pnpm install

# 启动前端
pnpm run dev
```

## 访问

启动项目后可以通过 `http://localhost:3000` 来访问前端，并且开发环境的接口请求会通过前端代理转发到后端来避免跨域问题。

# 其他版本

| 项目                                                            | 演示地址                                                             |
| --------------------------------------------------------------- | -------------------------------------------------------------------- |
| [RandallAnjie/moments](https://github.com/RandallAnjie/moments) | [https://moments.randallanjie.com](https://moments.randallanjie.com) |

# Contributors

感谢这些贡献代码的朋友。

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/kingwrcy"><img src="https://avatars.githubusercontent.com/u/1247324?v=4?s=80" width="80px;" alt="kingwrcy"/><br /><sub><b>kingwrcy</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/RandallAnjie"><img src="https://avatars.githubusercontent.com/u/84122428?v=4?s=80" width="80px;" alt="Randall"/><br /><sub><b>Randall</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Jonnyan404"><img src="https://avatars.githubusercontent.com/u/20352705?v=4?s=80" width="80px;" alt="jonny"/><br /><sub><b>jonny</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/akarikun"><img src="https://avatars.githubusercontent.com/u/11921182?v=4?s=80" width="80px;" alt="akari"/><br /><sub><b>akari</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/douseful"><img src="https://avatars.githubusercontent.com/u/52767905?v=4?s=80" width="80px;" alt="yee"/><br /><sub><b>yee</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://www.jschef.com"><img src="https://avatars.githubusercontent.com/u/38160059?v=4?s=80" width="80px;" alt="Chef"/><br /><sub><b>Chef</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://xwsir.cn"><img src="https://avatars.githubusercontent.com/u/17978673?v=4?s=80" width="80px;" alt="小王先森"/><br /><sub><b>小王先森</b></sub></a><br /></td>
    </tr>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://www.gooth.org"><img src="https://avatars.githubusercontent.com/u/126313?v=4?s=80" width="80px;" alt="Athurg Gooth"/><br /><sub><b>Athurg Gooth</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/xuewenG"><img src="https://avatars.githubusercontent.com/u/32838722?v=4?s=80" width="80px;" alt="xuewenG"/><br /><sub><b>xuewenG</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Secretlovez"><img src="https://avatars.githubusercontent.com/u/40491055?v=4?s=80" width="80px;" alt="Secretlovez"/><br /><sub><b>Secretlovez</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/jkjoy"><img src="https://avatars.githubusercontent.com/u/23159890?v=4?s=80" width="80px;" alt="浪子"/><br /><sub><b>浪子</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/lateautumn2"><img src="https://avatars.githubusercontent.com/u/57248164?v=4?s=80" width="80px;" alt="lateautumn2"/><br /><sub><b>lateautumn2</b></sub></a><br /></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/Jinvic"><img src="https://avatars.githubusercontent.com/u/77521861?v=4?s=80" width="80px;" alt="Jinvic"/><br /><sub><b>Jinvic</b></sub></a><br /></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=kingwrcy/moments&type=Date)](https://star-history.com/#kingwrcy/moments&Date)
