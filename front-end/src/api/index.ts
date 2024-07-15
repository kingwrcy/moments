import axios from "axios";
const service = axios.create({
    // baseURL 需要设置为反向代理前缀，不能设置绝对路径URL
    baseURL: '/api',
    timeout: 5000,
    withCredentials: false,
})
// 添加请求拦截器
service.interceptors.request.use(function (req) {
    const token = localStorage.getItem('token')
    if (token) {
        req.headers.set('x-api-token', token)
    }
    return req;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});
//添加响应拦截器
service.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    const res = response.data;
    if (response.status === 200) {

        if (res.code === 3 || res.code === 4) {
            toast.error('未登录')
            return Promise.reject(res.message);
        }else if(res.code !== 0){
            toast.error(res.message)
            return Promise.reject(res.message);
        }
    }
    return res;
}, function (error) {
    toast.error(error)
    return Promise.reject(error);
});

export default service