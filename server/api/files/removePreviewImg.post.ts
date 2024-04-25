import fs from "fs/promises";

type RemoveCommentReq = {
  path: string;
};

export default defineEventHandler(async (event) => {
  const { path } = (await readBody(event)) as RemoveCommentReq;
  if (!path || path.startsWith("http")) {
    return {
      success: false,
    };
  }
  const filename = path.replace("/upload/", "");
  try {
    fs.rm(`${process.env.UPLOAD_DIR}/${filename}`, {
      force: true,
      recursive: true,
    });
  } catch (error) {
    return {
      success: false,
    };
  }
  return {
    success: true,
  };
});
