import prisma from "~/lib/db";

type SaveCommentReq = {
  memoId: number;
  content: string;
  replyTo?: string;
  email?: string;
  website?: string;
  username: string;
  author: Boolean;
};

export default defineEventHandler(async (event) => {
  const { memoId, content, replyTo, username, email, website } =
    (await readBody(event)) as SaveCommentReq;
  const userId = event.context.userId;
  console.log(userId,userId ?true:false)
  await prisma.comment.create({
    data: {
      content,
      replyTo,
      memoId,
      username,
      email,
      website,
      author: userId !== undefined
    },
  });
  return {
    success: true,
  };
});
