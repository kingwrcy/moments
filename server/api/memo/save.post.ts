import prisma from "~/lib/db";

type SaveMemoReq = {
  id?: number;
  content: string;
  imgUrls?: string[];
  music163Url?: string;
  bilibiliUrl?: string;
  location?: string;
  externalUrl?: string;
  externalTitle?: string;
  externalFavicon?: string;
};

export default defineEventHandler(async (event) => {
  const body = (await readBody(event)) as SaveMemoReq;

  const updated = {
    imgs: body.imgUrls?.join(","),
    music163Url: body.music163Url,
    bilibiliUrl: body.bilibiliUrl,
    location: body.location,
    externalUrl: body.externalUrl,
    externalTitle: body.externalTitle,
    externalFavicon: body.externalFavicon,
    content: body.content,
  };
  await prisma.memo.upsert({
    where: {
      id: body.id ?? -1,
    },
    create: {
      userId: event.context.userId,
      ...updated,
    },
    update: {
      ...updated,
    },
  });
  return {
    success: true,
  };
});
