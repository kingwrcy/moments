import { load } from "cheerio";

type Request = {
  url: string;
};

export default defineEventHandler(async (event) => {
  const { url } = (await readBody(event)) as Request;
  let res: any;
  try {
    res = await $fetch(url, {
      timeout: 3000,
    });
  } catch (e) {
    return {
      success: false,
      title: "",
      favicon: "",
      message: "无法打开网页.手动填写标题吧!",
    };
  }

  const $ = load(res as string);

  let icons = $("link[rel='icon']");
  if (icons.length === 0) {
    icons = $("link[rel='shortcut icon']");
  }
  if (icons.length === 0) {
    icons = $("link[rel='apple-touch-icon']");
  }

  let href = "";
  const urlObject =  new URL(url)
  if (icons.length > 0) {
    href = icons.attr("href") || '';
  }else{
    href = urlObject.origin+'/favicon.ico';
  }
  if (!href.startsWith("http")) {
    href = urlObject.origin + (href.startsWith('/') ? href : '/'+href);
  }

  return {
    title: $("title").text(),
    favicon: href,
    message: "",
    success: true,
  };
});
