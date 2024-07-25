import type {ResultVO, SysConfigVO} from "~/types";
import {toast} from "vue-sonner";
import {useGlobalState} from "~/store";

const global = useGlobalState()

export async function useMyFetch<T>(url: string, data?: any) {
    const userinfo = global.value.userinfo
    const headers:Record<string, any> = {}
    if (userinfo.token){
        headers["x-api-token"] = userinfo.token
    }
    try {
        const res = await $fetch<ResultVO<T>>(`/api${url}`, {
            method: "post",
            body: data ? JSON.stringify(data) : null,
            headers:headers
        })
        if (res.code !== 0) {
            toast.error(res.message || "请求失败")
            throw new Error(res.message)
        }
        return res.data
    } catch (e) {
        if (e instanceof Error) {
            throw new Error(e.message)
        }
        throw new Error("接口异常")
    }
}

async function upload2S3(files: FileList) {
    const result = []
    for (let i = files.length - 1; i >= 0; i--) {
        const {preSignedUrl, imageUrl} = await useMyFetch<{
            preSignedUrl: string,
            imageUrl: string
        }>('/file/s3PreSigned',{
            contentType:files[0].type
        })
        const res = await $fetch(preSignedUrl, {
            method: "put",
            body: files[i],
            //@ts-ignore
            headers:{
                'Content-Type': null
            }
        })
        result.push(imageUrl)
    }
    return result
}


export async function useUpload(files: FileList | null) {
    const sysConfig = useState<SysConfigVO>('sysConfig')

    if (!files || files.length === 0) {
        toast.error("没有选择文件")
        return
    }

    const userinfo = global.value.userinfo
    const headers:Record<string, any> = {}
    if (userinfo.token){
        headers["x-api-token"] = userinfo.token
    }

    if (sysConfig.value.enableS3) {
        return await upload2S3(files)
    }
    const form = new FormData()
    for (let i = files.length - 1; i >= 0; i--) {
        form.append("files", files[i])
    }
    try {
        const res = await $fetch<ResultVO<string[]>>(`/api/file/upload`, {
            method: "post",
            body: form,
            headers
        })
        if (res.code !== 0) {
            toast.error(res.message || "请求失败")
            throw new Error(res.message)
        }
        return res.data
    } catch (e) {
        if (e instanceof Error) {
            throw new Error(e.message)
        }
        throw new Error("接口异常")
    }
}
