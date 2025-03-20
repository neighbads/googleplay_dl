# googleplay_dl
Google Play Apk 下载器

## 用法

1. 打开地址 [Google账户登录](https://accounts.google.com/embedded/setup/v2/android) 登录账户, 登录到最后一步

2. 在浏览器中打开开发者工具，切换到 应用程序 -> 存储 -> Cookies, 找到 `oauth_token`

3. 下载 apk
```bash
cd bin

# 登录
# 登录信息, 在 token 路径
./googleplay_dl -abi arm64-v8a -login oauth2_4/0Adeu5B...

# 下载
# 可多次下载, 在 apks 路径
./googleplay_dl -abi arm64-v8a -dl com.google.android.gm
```

- 一些常见的包名
```
# MicroAuth:    com.azure.authenticator
# MicroEdge:    com.microsoft.emmx
# Gmail    :    com.google.android.gm
# X        :    com.twitter.android
# Telegram :    org.telegram.messenger
# GoogleMap:    com.google.android.apps.maps
# Github   :    com.github.android
# Notion   :    notion.id
# Chrome   :    com.android.chrome
# Bitwarden:    com.x8bit.bitwarden
# QQ       :    com.tencent.mobileqq
```

4. 安装
```bash
# 单 apk 文件
adb install apks/com.microsoft.emmx-312406805.apk

# Split apk 文件
adb install-multiple apks/com.android.chrome-699809533.apk apks/com.android.chrome-chrome-699809533.apk apks/com.android.chrome-config.zh-699809533.apk apks/com.android.chrome-google3-699809533.apk
```
