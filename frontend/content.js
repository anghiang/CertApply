// 监听按钮点击事件
document.getElementById("login").addEventListener("click",function(){
    chrome.runtime.sendMessage({action: "openPopup"});
})