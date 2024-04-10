import fs from "node:fs";

export default defineEventHandler((event) => {
  const filename = getRouterParam(event, "filename");
  return sendStream(event, fs.createReadStream(`${process.env.UPLOAD_DIR}/${filename}`));
});
