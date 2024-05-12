import prisma from "~/lib/db";
import { MemoExt } from "~/lib/types";

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
  ext:MemoExt,
  showType: number;
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
    ext:JSON.stringify(body.ext),
    showType: body.showType ? 1 : 0
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
