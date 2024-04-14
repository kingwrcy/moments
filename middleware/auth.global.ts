export default defineNuxtRouteMiddleware((to, from) => {
  const cookie = useCookie('token')

  if (to.fullPath === '/settings' && !cookie.value) {
    return navigateTo('/login')
  }
  if (to.fullPath === '/login' && cookie.value) {
    return navigateTo('/')
  }
})