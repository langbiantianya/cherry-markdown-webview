# README

## About
一个功能丰富且支持`windows` `macos` `linux`的Markdown编辑器，目前暂时只支持中文，待基础功能完善后会添加更多的语言支持,由于我没有mac所以有些问题是无法发现的，如有发现问题欢迎提issues。
待后期想好名字后会迁移到新仓库并加一个图标。
## todo 画饼个
- [x] cherry-markdown
- [x] 鼠标右键
- [ ] 文件关联
	- [x] windows
	- [ ] linux
	- [ ] mac 
- [ ] 优化文件保存提示
- [ ] 跨平台支持(基础编辑器和文件io可用)
	- [x]  windos
		- [x] amd64
		- [x] arm64  
	- [ ]  linux
		- [x] amd64
		- [x] arm64
		- [ ] armf
		- [ ] loong64
	- [x]  mac
		- [x] arm64
		- [x] amd64
- [ ] 大文件打开性能优化
- [ ] webview依赖缺少提示 
- [ ] 本地图片、文件的解析
- [ ] 工作区
- [ ] 个性化选项
## download
[这里]([Releases · Tencent/cherry-markdown](https://github.com/Tencent/cherry-markdown/releases)) 
## screenshot
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/1.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/2.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/3.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/4.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/5.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/6.png)
![截图](https://github.com/langbiantianya/cherry-markdown-webview/raw/refs/heads/main/screenshot/7.png)
## Depends
- [cherry-markdown]([Tencent/cherry-markdown: ✨ A Markdown Editor](https://github.com/Tencent/cherry-markdown))
- [Wails]([The Wails Project | Wails](https://wails.io/zh-Hans/))

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
