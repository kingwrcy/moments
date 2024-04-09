import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const getImgUrl = (url: string) => {
  if (url.startsWith("http")) {
    return url + "?fmt=avif";
  }
  return url;
};
