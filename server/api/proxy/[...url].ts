export default defineEventHandler(async (event) => {
  const url = getRouterParam(event, "url") as string;
  if (url && url.startsWith("http")) {
    const body = await $fetch(url, {
      method: "GET",
      headers: {
        userAgent:
          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
      },
    });
    return body;
  }
});
