import prisma from "~/lib/db";
import fs from "fs/promises";

export default defineEventHandler(async (event) => {
  const data = await prisma.user.findUnique({
    where: {
      id: 1,
    },
  });
  if(!data){
    throw new Error("User not found");
  }
  const config = (await fs.readFile(`${process.env.CONFIG_FILE}`)).toString()
  return {
    success: true,data:{
      ...data,config
    }
  };
});
