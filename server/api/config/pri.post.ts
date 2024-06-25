import fs from "fs/promises";
import { PrivateConfig } from "~/lib/types";

export default defineEventHandler(async (event) => {
  const config = ((await fs.readFile(`${process.env.CONFIG_FILE}`)).toString())
  return JSON.parse(config).private as PrivateConfig;
});
