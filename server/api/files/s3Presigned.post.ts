import { PutObjectCommand, S3Client } from "@aws-sdk/client-s3";
import { getSignedUrl } from "@aws-sdk/s3-request-presigner";
import dayjs from "dayjs";
import prisma from "~/lib/db";
import short from "short-uuid";

export default defineEventHandler(async (event) => {
  const { fileType } = (await readBody(event)) as { fileType: string };


  const user = await prisma.user.findUnique({
    where: {
      id: event.context.userId,
    },
  });

  if (
    !fileType ||
    !user ||
    !user.enableS3 ||
    !user.accessKey ||
    !user.secretKey ||
    !user.bucket ||
    !user.region ||
    !user.endpoint
  ) {
    return {
      success: false,
      message: "S3 not enabled",
      url: "",
      imgUrl: "",
    };
  }

  const client = new S3Client({
    region: user.region,
    endpoint: user.endpoint,
    credentials: {
      accessKeyId: user.accessKey,
      secretAccessKey: user.secretKey,
    },
  });

  const key =
    "moments/" + dayjs(new Date()).format("YYYY/MM/DD/") + short.generate();
  const command = new PutObjectCommand({
    Bucket: user.bucket,
    Key: key,
    ContentType: fileType,
  });
  const url = await getSignedUrl(client, command, { expiresIn: 600 });
  let imgUrl = `${user.domain}/${key}`;
  return {
    success: true,
    url,
    imgUrl,
    message: "",
  };
});
