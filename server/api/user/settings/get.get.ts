import prisma from "~/lib/db";

export default defineEventHandler(async (event) => {
  const data = await prisma.user.findUnique({
    where: {
      id: 1,
    },
    select: {
      nickname: true,
      avatarUrl: true,
      slogan: true,
      coverUrl: true,
      favicon: true,
      title: true,
    },
  });
  if (!data) {
    throw createError({ status: 404, message: "Not Found" });
  }
  return {
    success: true,
    data,
  };
});
