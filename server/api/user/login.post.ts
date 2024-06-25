import bcrypt from "bcrypt";
import jwt from "jsonwebtoken";
import prisma from "~/lib/db";

type loginReq = {
  username: string;
  password: string;
};

export type JwtPayload = {
  username: string;
  exp: number;
  userId: number;
};

export default defineEventHandler(async (event) => {
  const runtimeConfig = useRuntimeConfig()

  const { username, password } = (await readBody(event)) as loginReq;
  let token = "";
  if (!username || !password) {
    return {
      success: false,
      message: "用户名或密码不能为空",
      token,
    };
  }

  const user = await prisma.user.findFirst({
    where: {
      username,
    },
  });

  if (!user || !bcrypt.compareSync(password, user.password)) {
    return {
      message: "用户名或密码错误",
      success: false,
      token,
    };
  }

  token = jwt.sign(
    {
      exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24,
      username: user.username,
      userId: user.id,
    },
    runtimeConfig.jwtKey
  );
  setCookie(event, "token", token, {
    expires: new Date(Date.now() + 60 * 60 * 24 * 1000),
  });

  return {
    success: true,
    username: user.username,
    message: "",
  };
});
