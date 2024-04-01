import bcrypt from "bcrypt";
import jwt from "jsonwebtoken";
import { jwtKey } from "~/lib/constant";
import prisma from "~/lib/db";

type loginReq = {
  username: string;
  password: string;
};


export type JwtPayload = {
  username:string
  exp:number
  userId:number
}

export default defineEventHandler(async (event) => {
  const body = (await readBody(event)) as loginReq;

  if (!body.username || !body.password) {
    return {
      message: "用户名或密码不能为空",
    };
  }


  const hash = bcrypt.hashSync(body.password, 10);
  const user = await prisma.user.findFirst({
    where: {
      username: body.username,
    },
  });

  if (!user || bcrypt.compareSync(hash, user.password)) {
    return {
      message: "用户名或密码错误",
    };
  }

  const token = jwt.sign(
    {
      exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24,
      username: user.username,
      userId:user.id,
    },
    jwtKey
  );

  return {
    username: user.username,
    token
  };
});
