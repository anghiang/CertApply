
// 监听来自内容脚本的消息
chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    // 弹出钱包页面
    if (request.action === "openPopup") {
        chrome.windows.create({
            url: "/views/index.html",
            type: "popup",
            width: 360,
            height: 600,
            top: 0,
            left: screen.availWidth - 180
        });
    }
    // 监听通过私钥明文导入账户并保存私钥到插件本地
    if (request.action === "priStr") {
        let addr = request.addrStr
        let pri = request.privateKeys
        let account = {
            address_ : addr,
            privateKey_ :pri
        }
        chrome.storage.local.set({accounts : account },function (){
            alert('值已成功存储');
        })

        // 读取值
     /
        
    }

  

    
});


