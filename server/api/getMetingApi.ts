import fs from "fs/promises";
import { PublicConfig } from "~/lib/types";

export default defineEventHandler(async (event) => {
    try{
        const config = ((await fs.readFile(`${process.env.CONFIG_FILE}`)).toString())
        const music_api = JSON.parse(config).public.music_api;
        return {
            success: true,
            data: {
                value: music_api
            }
        };
    }catch(e){
        return {
            success: false,
            data: {
                value: ''
            }
        };
    }
});