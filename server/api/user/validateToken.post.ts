import jwt from "jsonwebtoken";

export default defineEventHandler(async (event) => {
  const runtimeConfig = useRuntimeConfig()

  const token = getCookie(event, "token");
  if (token) {
    try {
      jwt.verify(token, runtimeConfig.jwtKey);
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
