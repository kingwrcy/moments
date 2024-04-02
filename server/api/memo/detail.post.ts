import prisma from "~/lib/db";

type DetailMemoReq = {
  id: number;
};

export default defineEventHandler(async (event) => {
  const { id } = (await readBody(event)) as DetailMemoReq;
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
      comments: {
        orderBy: {
          createdAt: "desc",
        },
      },
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
