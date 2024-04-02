import prisma from "~/lib/db";

type SaveMemoReq = {
  id?: number;
  content: string;
  imgUrls?: string[];
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
    },
    update: {
      imgs: body.imgUrls?.join(","),
      content: body.content,
    },
  });
  return {
    success: true,
  };
});
