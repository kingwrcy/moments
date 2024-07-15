import {createGlobalState, useStorage} from '@vueuse/core'
import type {Userinfo} from "@/types/user";

export const useGlobalState = createGlobalState(
    () => useStorage<Userinfo>('userinfo', {
        username: '',
        token: '',
        id: -1
    }),
)