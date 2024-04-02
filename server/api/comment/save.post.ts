import prisma from "~/lib/db";

type SaveCommentReq = {
  memoId: number;
  content: string;
  replyTo?: string;
  email?: string;
  website?: string;
  username: string;
};

export default defineEventHandler(async (event) => {
  const { memoId, content, replyTo, username, email, website } =
    (await readBody(event)) as SaveCommentReq;
  await prisma.comment.create({
    data: {
      content,
      replyTo,
      memoId,
      username,
      email,
      website,
    },
  });
  return {
    success: true,
  };
});
