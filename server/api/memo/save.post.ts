import prisma from "~/lib/db"



type SaveMemoReq = {
  content: string;
  imgUrls?: string[];
};


export default defineEventHandler(async (event) => {

  const body = (await readBody(event)) as SaveMemoReq;
    await prisma.memo.create({
      data:{
        createdAt:new Date(),
        updatedAt:new Date(),
        imgs:body.imgUrls?.join(','),
        content:body.content,
        userId:event.context.userId,
      }
    })
})