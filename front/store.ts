import { useStorage } from '@vueuse/core'
import type {LoginResp} from "~/types";

export type GlobalVO = {
    userinfo:LoginResp
}

export const useGlobalState = createGlobalState(
    () => useStorage<GlobalVO>('global',{
        userinfo:{},
    }),
)

