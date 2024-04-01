import { jwtKey } from "~/lib/constant";
import jwt from "jsonwebtoken";
import { JwtPayload } from "../api/user/login.post";

const ignorePaths = ["/api/user/login"];

export default defineEventHandler((event) => {
  const token = getHeader(event, "Authorization");
  const url = getRequestURL(event);
  
  if (ignorePaths.includes(url.pathname)) {
    return;
  }
  if(!token){
    throw createError({
      statusCode:401,
      statusMessage:"Unauthorized"
    })
  }

  try {
    const result = jwt.verify(token, jwtKey)
    const payload = result as JwtPayload
    event.context.userId = payload.userId
  } catch (error) {
    throw createError({
      statusCode:401,
      statusMessage:"Unauthorized"
    })
  }
})