### 常见问题

## 配置改坏了无法启动/打开首页报错如何处理?

删除你映射的/app/data目录下的config.json文件,然后重新启动镜像

## 如何加入访问量统计?

使用如下自定义JS

```js
document.addEventListener("DOMContentLoaded", () => {
  setTimeout(() => {
    document
      .querySelector(".footer")
      .insertAdjacentHTML(
        "afterend",
        `<div class='flex text-sm justify-center'><div id="busuanzi_container_site_pv">本站总访问量<span id="busuanzi_value_site_pv"></span>次</div></div>`
      );
    var script = document.createElement("script");
    script.src = "//busuanzi.ibruce.info/busuanzi/2.3/busuanzi.pure.mini.js";
    script.async = true;
    document.head.appendChild(script);
  }, 500);
});
```


## 如何去除底部github star?

虽然我不太愿意你去掉版权相关的东西,但是已经开放了自定义JS功能,所以这个可以通过自定义JS实现,以下方法选其一即可

#### 方式一 自定义JS

```js
document.addEventListener("DOMContentLoaded", () => {
  setTimeout(() => {
    document.querySelector(".footer a[rel='noopener noreferrer']").parentNode.remove()
  }, 500);
});
```

#### 方式二 自定义CSS

```css
.footer a[rel='noopener noreferrer']{display:none;}
```



## 如何更改全站字体

这里以落霞孤鹜字体为例子

自定义js

```js
function changeFont() { 
  const link = document.createElement("link");
  link.rel = "stylesheet";
  link.type = "text/css";
  link.href = "https://cdn.staticfile.org/lxgw-wenkai-screen-webfont/1.6.0/lxgwwenkaiscreen.css";
  document.head.append(link);
};
changeFont();
```

自定义css

```css
*{font-family: "LXGW WenKai Screen", sans-serif !important;}
```