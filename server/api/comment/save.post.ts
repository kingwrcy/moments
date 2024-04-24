import prisma from "~/lib/db";

type SaveCommentReq = {
  memoId: number;
  content: string;
  replyTo?: string;
  email?: string;
  website?: string;
  username: string;
  author: Boolean;
  token: string;
};

type recaptchaResponse = {
  success: boolean;
  challenge_ts: string;
  hostname: string;
  score: number;
  action: string;
};

const insertComment = async (userId: string, request: SaveCommentReq) => {
  await prisma.comment.create({
    data: {
      content: request.content,
      replyTo: request.replyTo,
      memoId: request.memoId,
      username: request.username,
      email: request.email,
      website: request.website,
      author: userId !== undefined,
    },
  });
};

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const request = (await readBody(event)) as SaveCommentReq;
  const { content, email, username, website, token } = request;

  const userId = event.context.userId;
  if (config.googleRecaptchaSecretKey && !token) {
    return { success: false, message: "小样儿,你是不是人机?" };
  }
  if (content.length > parseInt(config.public.momentsCommentMaxLength)) {
    return { success: false, message: "评论超长了,老板" };
  }
  if (username.length > 10) {
    return { success: false, message: "用户名你起这么长啥意思?" };
  }
  if (email && email.length > 30) {
    return { success: false, message: "邮箱你起这么长啥意思?" };
  }
  if (website && website.length > 100) {
    return { success: false, message: "网站地址也没这么长的啊" };
  }

  if (config.googleRecaptchaSecretKey) {
    const response = (await $fetch(
      `https://recaptcha.net/recaptcha/api/siteverify?secret=${config.googleRecaptchaSecretKey}&response=${token}`
    )) as any as recaptchaResponse;
    if (response.score > 0.5) {
      await insertComment(userId, request);
      return {
        success: true,
        message: "",
      };
    } else {
      return {
        success: false,
        message: "小样儿,你是不是人机?",
      };
    }
  } else {
    await insertComment(userId, request);
    return {
      success: true,
      message: "",
    };
  }
});
