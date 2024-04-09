import { useStorage } from "@vueuse/core";

export type UploadCallBack = (res: {
  success: boolean;
  message: string;
  filename: string;
}) => void;

export const useUpload = async (file: File, cb: UploadCallBack) => {
  const enableS3 = useStorage("enableS3", false);
  if (enableS3.value) {
    const res = await $fetch("/api/files/s3Presigned", {
      method: "POST",
      body: JSON.stringify({
        fileType: file.type,
      }),
    });
    console.log(res)
    if (res.success) {
      const uploadRes = await $fetch(res.url, {
        method: "PUT",
        body: file,
        // @ts-ignore
        headers: {
          "Content-Type": null,
        },
      });
      cb({
        success: uploadRes.status === 200,
        message: '',
        filename: res.imgUrl,
      });
    }
  } else {
    const formData = new FormData();
    formData.append("file", file);
    const res = await $fetch("/api/files/upload", {
      method: "POST",
      body: formData,
    });
    cb(res);
  }
};
