import { jwtKey } from "~/lib/constant";
import jwt from "jsonwebtoken";
import { JwtPayload } from "../api/user/login.post";

const needLoginUrl = [
  "/api/memo/save",
  "/api/files/s3Presigned",
  "/api/files/upload",
  "/api/memo/remove",
  "/api/user/settings/save",
  "/api/user/settings/full",
  "/api/memo/doubanBook",
  "/api/memo/doubanMovie",
  "/api/files/removePreviewImg",
];

export default defineEventHandler(async (event) => {
  const token = getCookie(event, "token");
  const url = getRequestURL(event);

  if (token && url.pathname === "/login") {
    await sendRedirect(event, "/", 302);
    return;
  }

  if (url.pathname === "/api/comment/save" && token) {
    try {
      const result = jwt.verify(token, jwtKey);
      const payload = result as JwtPayload;
      event.context.userId = payload.userId;
    } catch (error) {}
  }

  if (needLoginUrl.includes(url.pathname) && !token) {
    throw createError({
      statusCode: 401,
      statusMessage: "Unauthorized",
    });
  }

  if (needLoginUrl.includes(url.pathname) && token) {
    try {
      const result = jwt.verify(token, jwtKey);
      const payload = result as JwtPayload;
      event.context.userId = payload.userId;
    } catch (error) {
      throw createError({
        statusCode: 401,
        statusMessage: "Unauthorized",
      });
    }
  }
});
