# weibo_api

授权码获取方法：打开 https://api.weibo.com/2/oauth2/authorize?client_id=82966982&response_type=token&display=js&redirect_uri=https://api.weibo.com/oauth2/default.html ，授权后出现一个空白页面，不要关闭，按 Ctrl+U 查看源代码，找到 "accesstoken":"2.00tlp**_"。嗯，授权码就找到了。
