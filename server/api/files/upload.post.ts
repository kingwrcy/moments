import fs from "fs/promises";
import short from "short-uuid";
type FileInfo = { name: string; filename: string; data: Buffer; type: string };

export default defineEventHandler(async (event) => {
  const formData = await readMultipartFormData(event);
  if (!formData) {
    return {
      success: false,
      message: "No file found",
      filename: "",
    };
  }
  const file = formData[0] as FileInfo;
  if (!file.type.startsWith("image")) {
    return {
      success: false,
      message: "只支持上传图片文件",
      filename: "",
    };
  }
  const filetype = file.type.split("/")[1];

  const filename = short.generate() + "." + filetype;
  
  const filepath = `${process.env.UPLOAD_DIR}/${filename}`;
  await fs.writeFile(filepath, file.data);

  return {
    success: true,
    filename: "/upload/" + filename,
    message: "上传文件成功!",
  };
});
