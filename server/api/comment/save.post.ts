import prisma from "~/lib/db";
import { sendEmail } from '~/utils/sendEmail';

type SaveCommentReq = {
  memoId: number;
  content: string;
  replyTo?: string;
  replyToId?: number;
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
  const { content, email, memoId, replyToId, username, website, token } = request;

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

    } else {
      return {
        success: false,
        message: "小样儿,你是不是人机?",
      };
    }
  }
  let flag = true;
  if(replyToId !== undefined && replyToId !== 0){
    const comment = await prisma.comment.findUnique({
      where: {
        id: replyToId,
      }
    });
    if(comment !== null && comment.email !== null && comment.email !== ''){
      if(comment.email === process.env.NOTIFY_MAIL){
        flag = false;
      }
      // 邮箱通知被回复者
      const result = await sendEmail({
        email: comment.email,
        subject: '新回复',
        message: `您在moments中的评论有新回复！
                用户名为： ${username} 回复了您的评论(${comment.content})，他回复道：${content}，点击查看：${process.env.SITE_URL}/detail/${memoId}`,
      });
    }
  }

  // 非管理员
  if(event.context.userId == undefined && flag){
    // 判断process.env.SITE_URL是否以/结尾，如果是则去掉
    let siteUrl = process.env.SITE_URL;
    if(siteUrl === undefined || siteUrl === '' || siteUrl === 'undefined' || siteUrl === 'null'){
      siteUrl = '';
    }
    if(siteUrl.endsWith('/')){
      siteUrl = siteUrl.slice(0, -1);
    }

    // 邮箱通知管理员
    const result = await sendEmail({
      email: process.env.NOTIFY_MAIL || '',
      subject: '新评论',
      message: `您的moments有新评论！
            用户名为： ${username} 在您的moment中发表了评论：${content}，点击查看：${siteUrl}/detail/${memoId}`,
    });
  }
  await insertComment(userId, request);
  return {
    success: true,
    message: "",
  };
});
