import prisma from "~/lib/db";

type GetCommentReq = {
  memoId: number;
};

export default defineEventHandler(async (event) => {
  const { memoId } =
    (await readBody(event)) as GetCommentReq;
  const data = await prisma.comment.findMany({
    where: {
      memoId,
    },
    orderBy: {
      createdAt:"desc"
    },
    take:5
  });
  return {
    success: true,data
  };
});
