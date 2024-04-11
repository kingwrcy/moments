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
      message: "无法打开网页，请检查网址是否正确或网络是否正常",
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
  if (icons.length > 0) {
    href = icons.attr("href") || "";
  }
  if (!href.startsWith("http")) {
    href = new URL(url).origin + (href.startsWith('/') ? href : '/'+href);
  }

  return {
    title: $("title").text(),
    favicon: href,
    message: "",
    success: true,
  };
});
