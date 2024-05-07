### 常见问题

## 如何去除底部github star?

虽然我不太愿意你去掉版权相关的东西,但是已经开放了自定义JS功能,所以这个可以通过自定义JS实现,以下方法选其一即可

#### 方式一 自定义JS

```js
document.querySelector(".footer a[rel='noopener noreferrer']").parentNode.remove()
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