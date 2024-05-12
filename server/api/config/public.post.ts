import fs from "fs/promises";
import { PublicConfig } from "~/lib/types";

export default defineEventHandler(async (event) => {
  const config = ((await fs.readFile(`${process.env.CONFIG_FILE}`)).toString())
  return JSON.parse(config).public as PublicConfig;
});
