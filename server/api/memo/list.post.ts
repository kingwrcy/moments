import prisma from "~/lib/db";
import fs from "fs/promises";
import type { SysConfig } from "~/lib/types";
import jwt from "jsonwebtoken";
import { jwtKey } from "~/lib/constant";
import { JwtPayload } from "../user/login.post";

type ListMemoReq = {
  page: number;
  size: number;
  tagname: any;
};

export default defineEventHandler(async (event) => {
  const token = getCookie(event, "token");
  let userId = 0;
  if (token) {
    try {
      const user = jwt.verify(token, jwtKey) as JwtPayload;
      userId = user.userId;
    } catch {}
  }
  const fromUser = () => {
    return userId == 0 ? {} : { userId };
  };
  const fromShowType = () => {
    return userId == 0 ? [{ showType: 1 }] : [{ showType: 1 }, { showType: 0 }];
  };

  const config = (await fs.readFile(`${process.env.CONFIG_FILE}`)).toString();
  const sysConfig = JSON.parse(config) as SysConfig;

  const { page, tagname } = (await readBody(event)) as ListMemoReq;
  const size = sysConfig.public.pageSize;
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
      content: {
        contains: tagname ? "#" + tagname : "",
      },
      OR: [...fromShowType()],
      ...fromUser(),
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
        content: {
          contains: tagname ? "#" + tagname : "",
        },
        OR: [...fromShowType()],
        ...fromUser(),
      },
    });
    if (pinnedMemo) {
      // @ts-ignore
      data = [pinnedMemo, ...data];
    }
  }
  const total = await prisma.memo.count({
    where: {
      content: {
        contains: tagname ? "#" + tagname : "",
      },
    },
  });

  const totalPage = Math.ceil(total / size);
  return {
    data,
    hasNext: page < totalPage,
    success: true,
  };
});
