// 匹配页面文本并替换
function replaceTextOnPage() {
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

// 监听消息，执行替换
chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    if (request.action === 'replaceText') {
        replaceTextOnPage();
    }
});

// 第一次注入脚本时执行替换
replaceTextOnPage();
