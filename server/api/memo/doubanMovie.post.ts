import { PutObjectCommand, S3Client } from "@aws-sdk/client-s3";
import { load } from "cheerio";
import dayjs from "dayjs";
import fs from "fs/promises";
import short from "short-uuid";
import prisma from "~/lib/db";
type Request = {
  id: number;
};

export default defineEventHandler(async (event) => {
  const { id } = (await readBody(event)) as Request;
  const userId = event.context.userId;
  let res: any;
  let url = `https://movie.douban.com/subject/${id}/`;

  const user = await prisma.user.findUnique({
    where: {
      id: userId,
    },
  });

  try {
    res = await $fetch(url, {
      timeout: 5000,
    });
  } catch (e) {
    return {
      success: false,
      title: "",
      favicon: "",
      message: "获取豆瓣读书信息异常,请重试",
      desc: "",
      image: "",
      director: "",
      rating: "",
      releaseDate: "",
      id: 0,
      runtime:"",
      actors:"",
      url,
    };
  }

  const $ = load(res as string);

  const metas = $("meta[property]");
  let title = "";
  let desc = "";
  let image = "";
  let director = "";
  let rating = "";
  let releaseDate = "";
  let actors = "";
  let runtime = "";
  metas.each((i, el) => {
    const property = $(el).attr("property");
    if (property === "og:title") {
      title = $(el).attr("content") || "";
    } else if (property === "og:description") {
      desc = $(el).attr("content") || "";
    } else if (property === "og:image") {
      image = $(el).attr("content") || "";
    } else if (property === "video:director") {
      director = $(el).attr("content") || "";
    } else if (property === "video:actor") {
      actors += ($(el).attr("content") || "") + "/";
    }
  });

  releaseDate =
    $("span[property='v:initialReleaseDate']").attr("content") || "";
  runtime = $("span[property='v:runtime']").attr("content") || "";
  rating = $("strong.rating_num").text() || "未知评分";

  if (image && image.startsWith("http")) {
    const body = (await $fetch(image, {
      method: "GET",
      headers: {
        userAgent:
          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
      },
    })) as any as Response;

    const data = await body.arrayBuffer();
    const fileBuffer = Buffer.from(data);
    if (user!.enableS3) {
      const client = new S3Client({
        region: user!.region!,
        endpoint: user!.endpoint!,
        credentials: {
          accessKeyId: user!.accessKey!,
          secretAccessKey: user!.secretKey!,
        },
      });
      const key =
        "moments/" + dayjs(new Date()).format("YYYY/MM/DD/") + short.generate();
      const command = new PutObjectCommand({
        Bucket: user!.bucket!,
        Key: key,
        ContentType: "image/jpg",
        Body: fileBuffer,
      });
      await client.send(command);
      image = `${user!.domain}/${key}`;
    } else {
      const filename = short.generate() + ".jpg";
      const filepath = `${process.env.UPLOAD_DIR}/${filename}`;
      try {
        await fs.writeFile(filepath, fileBuffer);
      } catch (e) {
        console.log("filepath is : ", filepath);
        console.log("writeFile error is : ", e);
      }
      image = "/upload/" + filename;
    }
  }

  return {
    title,
    desc,
    image,
    director,
    url,
    rating,
    releaseDate,
    actors,
    message: "",
    id,
    runtime,
    success: true,
  };
});
