import prisma from "~/lib/db";

type RemoveCommentReq = {
  commentId: number;
};

export default defineEventHandler(async (event) => {
  const { commentId } =
    (await readBody(event)) as RemoveCommentReq;
  await prisma.comment.delete({
    where: {
      id:commentId,
    },
  });
  return {
    success: true
  };
});
