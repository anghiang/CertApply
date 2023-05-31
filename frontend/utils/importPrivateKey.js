$(document).ready(function () {
    $('#stringSubmit').click(function (){
        let privateKey = $('#privateKey').val()
        let addr
        $.ajax({
            url:'http://127.0.0.1:5000/stringSubmit',
            type: 'POST',
            data: {
                privateKey_:privateKey
            },
            success: function (data){
                getAccounts(data.address_str,function (result) {
                    if (result===null){
                        alert("该账户已导入")
                    }else {
                        chrome.runtime.sendMessage({action:"priStr",addrStr :addr,privateKeys :data.private_key_str});
                    }
                })
            }
        })
    })
    $('#pemFileSubmit').submit(function (e){
        e.preventDefault()
        var formData = new FormData();
        var file = $('#fileInput')[0].files[0];
        formData.append('file', file);

        $.ajax({
            url: 'http://127.0.0.1:5000/pemSubmit',
            type: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function(response) {
                alert(response)
                // 处理成功响应
            },
            error: function(xhr, status, error) {
                // 处理错误响应
            }
        });
    })

    $('#showList').click(function(){
        chrome.storage.local.get('accounts', function(result) {
            let storedAccount = result.accounts;
            console.log("storedAccount :",storedAccount)
            // if (storedAccount) {
            //     let address = storedAccount.address_;
            //     let privateKey = storedAccount.privateKey_;
            //     alert('存储的地址:'+ address);
            //     alert('存储的私钥:'+ privateKey);
            // } else {
            //     alert('未找到存储的值');
            // }
        });
    })
})



function getAccounts(accounts, callback) {
    chrome.storage.local.get(accounts, function(result) {
        let storedAccount = result.accounts;
        if (!storedAccount) {
            callback(storedAccount);
        } else {
            callback(null);
        }
    });
}