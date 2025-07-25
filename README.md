# README

## About

一个功能丰富且支持`windows` `macos` `linux(龙芯也支持)`的 Markdown 编辑器，目前暂时只支持中文，待基础功能完善后会添加更多的语言支持,由于我没有 mac 所以有些问题是无法发现的，如有发现问题欢迎提 issues。
待后期想好名字后会迁移到新仓库并加一个图标。

## 已知问题

- [x] 链接跳转
- [ ] 图片查看
- [ ] 自定义界面难看
- [ ] 提示信息不正常
- [ ] 缺少自动保存
- [ ] 窗口最多只能打开2个
- [ ] 缺少仅查看

## 主题

目前添加了 2 个主题的示例[芙丽莲](https://github.com/langbiantianya/edit-theme-frieren)和[瑶瑶](https://github.com/langbiantianya/edit-theme-yaoyao)

### 自定义主题

#### 简单

的可以通过个性化选项中的可视化配置颜色与背景图片。

#### 复杂

可以参考[芙丽莲](https://github.com/langbiantianya/edit-theme-frieren)与[瑶瑶](https://github.com/langbiantianya/edit-theme-yaoyao)还有[cherryMarkdown 配置主题](https://github.com/Tencent/cherry-markdown/wiki/%E9%85%8D%E7%BD%AE%E4%B8%BB%E9%A2%98)的文档来进行更多的外观定制。

## todo 画饼个

- [x] cherry-markdown
- [x] 鼠标右键
- [ ] 文件关联
  - [x] windows
  - [ ] linux
- [x] 优化文件保存提示
- [ ] 跨平台支持(基础编辑器和文件 io 可用)
  - [x] windos
    - [x] amd64
    - [x] arm64
  - [ ] linux
    - [x] amd64
    - [x] arm64
    - [ ] armf
    - [ ] loong64
- [x] 大文件打开性能优化
- [x] webview 依赖缺少提示
- [x] 本地图片解析
- [x] 图片上传对象存储
- [x] 个性化选项

## download

[这里下载](https://github.com/langbiantianya/cherry-markdown-webview/releases)

## screenshot

![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/1.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/2.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/3.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/4.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/5.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/6.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/7.png)

## Depends

- [Tencent/cherry-markdown: ✨ A Markdown Editor](https://github.com/Tencent/cherry-markdown)
- [The Wails Project | Wails](https://wails.io/zh-Hans/)
- [bytedance/iconpark](https://github.com/bytedance/iconpark)
- [zimg](https://github.com/lkzc19/zimg)

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on <http://localhost:34115>. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
