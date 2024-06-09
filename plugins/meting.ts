// plugins/meting.js
export default defineNuxtPlugin(async (nuxtApp) => {
    if (process.client) {
        const response = await fetch('/api/getMetingApi');
        const data = await response.json();
        if(data.success && data.data.value && data.data.value !== ''){
            if(data.data.value.endsWith('/')){
                window.meting_api = data.data.value + 'api?server=:server&type=:type&id=:id&auth=:auth&r=:r';
            }else{
                window.meting_api = data.data.value + '/api?server=:server&type=:type&id=:id&auth=:auth&r=:r';
            }
        }else{
            window.meting_api = 'https://meting-dd.2333332.xyz/api?server=:server&type=:type&id=:id&auth=:auth&r=:r';
        }
    }
});