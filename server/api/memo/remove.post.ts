import prisma from "~/lib/db";

type RemoveCommentReq = {
  memoId: number;
};

export default defineEventHandler(async (event) => {
  const { memoId } = (await readBody(event)) as RemoveCommentReq;
  await prisma.memo.delete({
    where: {
      id: memoId,
    },
  });
  return {
    success: true,
  };
});
