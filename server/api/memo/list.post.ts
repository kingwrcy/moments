import prisma from "~/lib/db";

type ListMemoReq = {
  page: number;
  size: number;
  tagname: any;
};

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();


  const { page, tagname } = (await readBody(event)) as ListMemoReq;
  // const size = config.pageSize;
    const size = 10;
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
          createdAt: config.public.momentsCommentOrderBy,
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
        contains: tagname? '#'+tagname: '',
      },
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
          contains: tagname? '#'+tagname: '',
        },
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
        contains: tagname? '#'+tagname: '',
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
