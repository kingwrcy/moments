import prisma from "~/lib/db";

type SaveMemoReq = {
  id?: number;
  content: string;
  imgUrls?: string[];
  music163Url?: string;
  bilibiliUrl?: string;
  location?: string;
};

export default defineEventHandler(async (event) => {
  const body = (await readBody(event)) as SaveMemoReq;
  await prisma.memo.upsert({
    where: {
      id: body.id ?? -1,
    },
    create: {
      imgs: body.imgUrls?.join(","),
      content: body.content,
      userId: event.context.userId,
      music163Url: body.music163Url,
      bilibiliUrl: body.bilibiliUrl,
      location: body.location,
    },
    update: {
      imgs: body.imgUrls?.join(","),
      content: body.content,
      music163Url: body.music163Url,
      bilibiliUrl: body.bilibiliUrl,
      location: body.location,
    },
  });
  return {
    success: true,
  };
});
