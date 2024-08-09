function copyToClipboard(text) {
    if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(text)
    } else {
        // 兼容性处理：使用临时文本区域
        const textArea = document.createElement('textarea');
        textArea.value = text;
        document.body.appendChild(textArea);
        textArea.select();
        try {
            document.execCommand('copy');
        } catch (err) {
            console.error('复制失败:', err);
        }
        document.body.removeChild(textArea);
    }
}


document.addEventListener('DOMContentLoaded',()=>{
    setTimeout(()=>{
        const list = document.querySelectorAll(".markdown-content>pre.shiki")
        list.forEach((pre)=>{
            const copyBtn = document.createElement('div')
            copyBtn.innerText='复制'
            copyBtn.classList.add('copyBtn')
            copyBtn.addEventListener('click',()=>{
                copyToClipboard(pre.querySelector("code").innerText)
                copyBtn.innerText = '已复制!'
                copyBtn.style.color= '#ccc'
                setTimeout(()=>{
                    copyBtn.innerText = '复制'
                    copyBtn.style.color= '#000'
                },3000)
            })
            pre.insertAdjacentElement('beforeend',copyBtn)
        })
    },3000)
})