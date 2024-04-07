import fs from "fs/promises";

export default defineEventHandler(async (event) => {
  const filename = getRouterParam(event, "filename");
  if (!filename) {
    return {
      success: false,
      message: "No file found",
    };
  }

  try {
    return await fs.readFile(`/moments/${filename}`);
  } catch (error) {
    return {
      success: false,
      message: "File not found",
    };
  }
});
