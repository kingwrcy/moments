import prisma from "~/lib/db";

export default defineEventHandler(async (event) => {
  const data = await prisma.user.findUnique({
    where: {
      id: 1,
    },
  });
  if(!data){
    throw new Error("User not found");
  }
  return {
    success: true,data
  };
});
