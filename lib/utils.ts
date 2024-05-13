import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import type { User } from "./types";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const getImgUrl = (url: string) => {
  const userinfo = useState<User>("userinfo");
  const suffix = userinfo.value.thumbnailSuffix || "";
  if (url.startsWith("http") && suffix) {
    return url + (suffix.startsWith("?") ? suffix : "?" + suffix);
  }
  return url;
};

export const insertTextAtCursor = (
  text: string,
  textarea: HTMLTextAreaElement | undefined
) => {
  if (!textarea) return; // 检查textarea是否存在

  var cursor = textarea.selectionStart;
  var textLength = textarea.value.length;
  var selectedText = textarea.value.substring(cursor, textarea.selectionEnd);

  // 如果选中了文本，则替换选中的文本，否则插入新文本
  var newText = selectedText.length > 0 ? text : selectedText + text;

  // 边界检查
  if (cursor > textLength) cursor = textLength;
  if (cursor < 0) cursor = 0;

  // 更新文本内容
  textarea.value =
    textarea.value.substring(0, cursor) +
    newText +
    textarea.value.substring(textarea.selectionEnd);

  // 重新设置光标位置
  textarea.setSelectionRange(cursor + text.length, cursor + text.length);

  // 确保新插入的文本可见
  textarea.scrollTop = textarea.scrollHeight;
};
