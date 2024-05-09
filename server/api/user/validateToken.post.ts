import jwt from "jsonwebtoken";
import { jwtKey } from "~/lib/constant";

export default defineEventHandler(async (event) => {
  const token = getCookie(event, "token");
  if (token) {
    try {
      jwt.verify(token, jwtKey);
    } catch (e) {
      return {
        success: false,
      };
    }
  }
  return {
    success: true,
  };
});
