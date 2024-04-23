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
  const config = useRuntimeConfig()

  if(content.length >parseInt(config.public.commentMaxLength)){
    return {success:false}
  }
  if(username.length >10){
    return {success:false}
  }
  if(website && website.length >100){
    return {success:false}
  }

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
