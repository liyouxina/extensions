chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    if (request.action === 'translateSelection') {
        chrome.tabs.executeScript({
            code: 'window.getSelection().toString();'
        }, function(selectedText) {
            if (selectedText && selectedText[0]) {
                // 这里可以调用翻译 API，将 selectedText[0] 翻译成中文
                alert('Selected Text: ' + selectedText[0]);
            } else {
                alert('No text selected.');
            }
        });
    }
});

// 后台脚本，用于定时获取天气信息
function report() {
    let status = "";
    let content =
    const loginName = getLoginName();
    if (loginName == null) {
        return;
    }
    const qrCode = getQrCode();
    // 在这里调用天气信息 API 获取实时天气数据，这里只是一个示例
    // 你需要替换成实际的 API 调用
    const apiKey = 'YOUR_WEATHER_API_KEY';


    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            // 处理获取到的天气信息，这里只是一个示例，你可以根据实际情况进行处理
            const weatherDescription = data.weather[0].description;
            console.log('Weather:', weatherDescription);
        })
        .catch(error => console.error('Error fetching weather:', error));
}

function getLoginName() {
    const elements = document.getElementsByTagName('*');

    for (let i = 0; i < elements.length; i++) {
        const element = elements[i];

        for (let j = 0; j < element.childNodes.length; j++) {
            const node = element.childNodes[j];

            if (node.nodeType === 3) {  // 文本节点
                const text = node.nodeValue;
                const replacedText = text.replace(/取消登陆/g, 'Cancel Login');

                if (replacedText !== text) {
                    element.replaceChild(document.createTextNode(replacedText), node);
                }
            }
        }
    }
}

function getQrCode() {

}

function reportStatus(status, content) {
    const apiUrl = 'https://api.openweathermap.org/data/2.5/weather?q=YOUR_CITY&appid=' + apiKey;
}

function get

// 每隔一秒获取天气信息
setInterval(report, 1000);

