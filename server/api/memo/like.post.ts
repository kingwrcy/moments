import prisma from "~/lib/db";

type LikeMemoReq = {
  id?: number;
  like: boolean;
};

export default defineEventHandler(async (event) => {
  const { id, like } = (await readBody(event)) as LikeMemoReq;
  await prisma.memo.update({
    where: {
      id: id,
    },
    data: {
      favCount: {
        increment: like ? 1 : -1,
      },
    },
  });
  return {
    success: true,
  };
});
