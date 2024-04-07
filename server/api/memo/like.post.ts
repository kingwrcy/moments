import prisma from "~/lib/db";

type LikeMemoReq = {
  memoId?: number;
  like: boolean;
};

export default defineEventHandler(async (event) => {
  const { memoId, like } = (await readBody(event)) as LikeMemoReq;
  await prisma.memo.update({
    where: {
      id: memoId,
    },
    data: {
      favCount: {
        increment: like ? 1 : -1,
      },
    },
  });

  const data = await prisma.memo.findUnique({
    where: {
      id: memoId,
    },
    select: {
      favCount:true
    },
  });
  return {
    success: true,data
  };
});
