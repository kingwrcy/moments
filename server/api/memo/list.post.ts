import prisma from "~/lib/db";

type ListMemoReq = {
  page: number;
};

export default defineEventHandler(async (event) => {
  const { page } = (await readBody(event)) as ListMemoReq;
  const size = 20;
  const data = await prisma.memo.findMany({
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
        take:5,
      },
      _count: {
        select:{
          comments:true
        }
      }
    },
    orderBy: {
      createdAt: "desc",
    },
    skip: (page - 1) * size,
    take: size,
  });
  return {
    data,
    success: true,
  };
});
