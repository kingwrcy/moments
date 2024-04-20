import prisma from "~/lib/db";

type LikeMemoReq = {
  memoId?: number;
  pinned: boolean;
};

export default defineEventHandler(async (event) => {
  const { memoId,pinned } = (await readBody(event)) as LikeMemoReq;
  await prisma.memo.updateMany({
    where: {
      pinned: true,
    },
    data: {
      pinned: false,
    },
  });

  await prisma.memo.update({
    where: {
      id: memoId,
    },
    data: {
      pinned,
    },
  });
  return {
    success: true
  };
});
