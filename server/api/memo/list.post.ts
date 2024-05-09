import prisma from "~/lib/db";
import fs from "fs/promises";
import type { SysConfig } from "~/lib/types";

type ListMemoReq = {
  page: number;
  size: number;
};

export default defineEventHandler(async (event) => {
  const cookie = event.headers.get('cookie') || '';
  let showType = cookie ? [{ showType: 1 },{ showType: 0 }] : [{ showType: 1 }]

  const config = ((await fs.readFile(`${process.env.CONFIG_FILE}`)).toString())
  const sysConfig = JSON.parse(config) as SysConfig

  const { page ,size} = (await readBody(event)) as ListMemoReq;
  let data = await prisma.memo.findMany({
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
      comments: {
        //@ts-ignore
        orderBy: {
          createdAt: sysConfig.private.commentOrderBy,
        },
        take: 5,
      },
      _count: {
        select: {
          comments: true,
        },
      },
    },
    where: {
      pinned: false,
      OR: showType,
    },
    orderBy: {
      createdAt: "desc",
    },
    skip: (page - 1) * size,
    take: size,
  });
  if (page === 1) {
    const pinnedMemo = await prisma.memo.findFirst({
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
        comments: {
          orderBy: {
            createdAt: "desc",
          },
          take: 5,
        },
        _count: {
          select: {
            comments: true,
          },
        },
      },
      where: {
        pinned: true,
      },
    });
    if (pinnedMemo) {
      // @ts-ignore
      data = [pinnedMemo, ...data];
    }
  }
  const total = await prisma.memo.count();
  const totalPage = Math.ceil(total / size);
  return {
    data,
    hasNext: page < totalPage,
    success: true,
  };
});
