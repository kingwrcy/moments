import prisma from "~/lib/db";
import fs from "fs/promises";
import { SysConfig } from "~/lib/types";

type DetailMemoReq = {
  id: number;
};

export default defineEventHandler(async (event) => {
  const { id } = (await readBody(event)) as DetailMemoReq;
  const config = ((await fs.readFile(`${process.env.CONFIG_FILE}`)).toString())
  const sysConfig = JSON.parse(config) as SysConfig
  const data = await prisma.memo.findUnique({
    where:{
      id
    },
    include: {
      user: {
        select: {
          username: true,
          nickname: true,
          slogan: true,
          id: true,
          avatarUrl: true,
          coverUrl: true,
        },
      },
      comments: sysConfig.public.enableShowComment ?  {
        orderBy: {
          createdAt: "desc",
        },
      } : undefined,
      _count: {
        select:{
          comments:true
        }
      }
    },    
  });
  return {
    data,
    success: true,
  };
});
