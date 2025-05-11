import type { ResultVO, SysConfigVO } from "~/types"
import { toast } from "vue-sonner"
import { useGlobalState } from "~/store"
import markdownit from "markdown-it"
import { fromHighlighter } from "@shikijs/markdown-it/core"
import { createHighlighterCore } from "shiki/core"

const global = useGlobalState()

export const useMyFetch = async <T>(url: string, data?: any) => {
  const headers: Record<string, string> = {}

  const userinfo = global.value.userinfo
  if (userinfo.token) {
    headers["x-api-token"] = userinfo.token
  }

  const res = await $fetch<ResultVO<T>>(`/api${url}`, {
    method: "post",
    body: data ? JSON.stringify(data) : null,
    headers: headers,
  })

  if (!res || res.code !== 0) {
    if (!res) {
      throw new Error("请求失败")
    }

    if (res.code === 3 || res.code === 4) {
      global.value.userinfo = {}
      window.location.href = "/"
      throw new Error(res.message || "请求失败")
    }

    throw new Error(res.message)
  }

  return res.data
}

type OnProgressCallback = (progress: number) => void

type OnTotalProgressCallback = (
  totalCount: number,
  currentCount: number,
  name: string,
  progress: number,
) => void

const upload2S3WithProgress = async (
  preSignedUrl: string,
  file: File,
  onProgress: OnProgressCallback,
): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.addEventListener("load", () => {
      resolve()
    })
    xhr.addEventListener("error", () => reject(new Error("File upload failed")))
    xhr.addEventListener("abort", () =>
      reject(new Error("File upload aborted")),
    )
    xhr.upload.addEventListener("progress", e => onProgress(e.loaded / e.total))

    xhr.open("PUT", preSignedUrl, true)
    xhr.send(file)
  })

const upload2S3 = async (
  files: FileList,
  onProgress?: OnTotalProgressCallback,
): Promise<string[]> => {
  const result: string[] = []

  for (let i = 0; i < files.length; i++) {
    try {
      const res = await useMyFetch<{
        preSignedUrl: string
        imageUrl: string
      }>("/file/s3PreSigned", {
        contentType: files[0].type,
      })

      if (!res || !res.preSignedUrl) {
        toast.error("获取 S3 上传地址失败")
        continue
      }

      const file = files[i]
      await upload2S3WithProgress(res.preSignedUrl, file, progress => {
        if (onProgress) {
          onProgress(files.length, i + 1, file.name, progress)
        }
      })
      result.push(res.imageUrl)
    } catch (err) {
      toast.error(`上传图片到 S3 失败, ${err}`)
    }
  }

  return result
}

const uploadFile2ServerWithProgress = (
  url: string,
  file: File,
  onProgress: OnProgressCallback,
): Promise<string[]> =>
  new Promise<string[]>((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.addEventListener("load", () => {
      const res = JSON.parse(xhr.responseText)
      if (!res || res.code !== 0) {
        return reject(new Error(`${res?.message || "请求失败"}`))
      }

      resolve(res.data || [])
    })
    xhr.addEventListener("error", () => reject(new Error("File upload failed")))
    xhr.addEventListener("abort", () =>
      reject(new Error("File upload aborted")),
    )
    xhr.upload.addEventListener("progress", e => onProgress(e.loaded / e.total))

    xhr.open("POST", url, true)

    const userinfo = global.value.userinfo
    if (userinfo.token) {
      xhr.setRequestHeader("x-api-token", userinfo.token)
    }

    const formData = new FormData()
    formData.append("files", file)

    xhr.send(formData)
  })

const uploadFile2Server = async (
  files: FileList,
  onProgress?: OnTotalProgressCallback,
): Promise<string[]> => {
  const result: string[] = []

  for (let i = 0; i < files.length; i++) {
    try {
      const file = files[i]
      const urlList = await uploadFile2ServerWithProgress(
        "/api/file/upload",
        file,
        progress => {
          if (onProgress) {
            onProgress(files.length, i + 1, file.name, progress)
          }
        },
      )

      if (!urlList.length) {
        toast.error(`上传图片到服务器失败`)
        continue
      }

      result.push(...urlList)
    } catch (e) {
      toast.error(`上传图片到服务器失败, ${e}`)
    }
  }

  return result
}

export const useUpload = async (
  files: FileList,
  onProgress?: OnTotalProgressCallback,
): Promise<string[]> => {
  if (files.length === 0) {
    toast.error("没有选择文件")
    return []
  }

  const sysConfig = useState<SysConfigVO>("sysConfig")
  if (sysConfig.value.enableS3) {
    return upload2S3(files, onProgress)
  }

  return uploadFile2Server(files, onProgress)
}

export const md = markdownit({
  html: true,
  linkify: true,
  typographer: true,
  breaks: true,
})

createHighlighterCore({
  themes: [import("shiki/themes/github-dark.mjs")],
  langs: [
    import("shiki/langs/c.mjs"),
    import("shiki/langs/css.mjs"),
    import("shiki/langs/html.mjs"),
    import("shiki/langs/javascript.mjs"),
    import("shiki/langs/json.mjs"),
    import("shiki/langs/python.mjs"),
    import("shiki/langs/shellscript.mjs"),
    import("shiki/langs/sql.mjs"),
    import("shiki/langs/tsx.mjs"),
    import("shiki/langs/xml.mjs"),
    import("shiki/langs/yaml.mjs"),
    import("shiki/langs/go.mjs"),
  ],
  loadWasm: import("shiki/wasm"),
}).then(highlighter => {
  md.use(
    //@ts-ignore
    fromHighlighter(highlighter, {
      themes: {
        light: "github-dark",
        dark: "github-dark",
      },
    }),
  )
})

export const getGuestId = () => {
  // 优先从 localStorage 获取
  let storedData = localStorage.getItem("guest_info");

  // 尝试从 cookie 获取（如果 localStorage 没有）
  if (!storedData) {
    const cookieValue = document.cookie.split('; ')
      .find(row => row.startsWith('guest_info='))
      ?.split('=')[1];
      
    if (cookieValue) {
      try {
        storedData = decodeURIComponent(cookieValue);
        localStorage.setItem("guest_info", storedData);
      } catch (e) {
        console.error('解析 cookie 失败', e);
      }
    }
  }

  if (storedData) {
    try {
      const { guestId, timestamp } = JSON.parse(storedData);
      // 有效期为 15 天
      if (guestId && Date.now() - timestamp < 15 * 24 * 60 * 60 * 1000) {
        return guestId;
      }
    } catch (e) {
      console.error('解析访客信息失败', e);
    }
  }

  // 生成短随机 ID（6位字符）
  const CHARS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let shortId = '';
  for (let i = 0; i < 6; i++) {
    shortId += CHARS.charAt(Math.floor(Math.random() * CHARS.length));
  }
  
  const newGuestId = `访客_${shortId}`;
  const data = JSON.stringify({ guestId: newGuestId, timestamp: Date.now() });
  
  // 同时写入 localStorage 和 cookie
  try {
    localStorage.setItem("guest_info", data);
  } catch (e) {
    console.warn("无法写入 localStorage（可能处于无痕模式）", e);
  }
  
  document.cookie = `guest_info=${encodeURIComponent(data)}; path=/; max-age=1296000; secure; samesite=lax`;
  
  return newGuestId;
};