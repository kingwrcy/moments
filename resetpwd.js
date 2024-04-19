import { PrismaClient } from '@prisma/client'
import { createInterface } from "readline";
import bcrypt from "bcrypt";

const rl = createInterface({
  input: process.stdin,
  output: process.stdout,
});

const prisma = new PrismaClient()

// 设置一个问询，提示用户输入
rl.question("输入你的新密码: ", async (answer) => {
  if (!answer) {
    console.log("密码不能为空！");
    rl.close();
    return;
  }
  console.log(`你的新密码是: ${answer}`);

  const pwd = bcrypt.hashSync(answer, 10);
  console.log(`加密后是: ${pwd}`);

  await prisma.user.update({
    where: {
      username: "admin",
    },
    data: {
      password: pwd,
    },
  });

  console.log("密码修改成功！请去重新登录吧!");
  // 关闭readline接口
  rl.close();
});
