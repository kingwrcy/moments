import { jwtKey } from "~/lib/constant";
import jwt from "jsonwebtoken";
import { JwtPayload } from "../api/user/login.post";

const needLoginUrl = ["/api/memo/save"];

export default defineEventHandler(async (event) => {
  const token = getCookie(event,'token')
  const url = getRequestURL(event);

  if(token && url.pathname === "/login"){
    await sendRedirect(event, '/', 302)
    return 
  }
  
  if (!needLoginUrl.includes(url.pathname)) {
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