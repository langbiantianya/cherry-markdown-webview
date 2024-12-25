// eslint-disable-next-line no-unused-vars
import Cherry from "cherry-markdown"
import { GetPicBed, GetWebServerPort, UploadPicbed } from "../wailsjs/go/main/App"
import { isRelativePath, isRootDirectory } from "./fileMenu"
// eslint-disable-next-line no-unused-vars
import { config, file } from "../wailsjs/go/models"

/**
 * 
 * @param {Cherry} cherryInstance 
 * @param {file.File} assciateOpenFile 
 */
export async function hookbeforeImageMounted(assciateOpenFile, cherryInstance) {
	let webServerPort = await GetWebServerPort()
	/**
 * 将渲染的img的src指向本地文件服务
 */
	cherryInstance.on(
		"beforeImageMounted",
		/**
		 *
		 * @param {string} srcProp
		 * @param {string} src
		 */
		function (srcProp, src) {
			if (!/^data:/.test(src)) {
				try {
					let url = new URL(src)
					console.log("url", src)
					if (url.protocol === "file:") {
						return {
							srcProp,
							src: `http://127.0.0.1:${webServerPort}/file?uri=${url.href}`,
						}
					}
				} catch (e) {
					if (e instanceof TypeError && isRelativePath(src)) {
						return {
							srcProp,
							src: `http://127.0.0.1:${webServerPort}/file?absPath=${src}&docPath=${assciateOpenFile.Path}`,
						}
					} else if (e instanceof TypeError && isRootDirectory(src)) {
						return {
							srcProp,
							src: `http://127.0.0.1:${webServerPort}/file?uri=file:///${src}`,
						}
					}
				}
			}
			return { srcProp, src }
		}
	)
}

/**
 * 
 * @param {Cherry} cherryInstance 
 */
export function hookFileUpload(cherryInstance) {
	cherryInstance.on('fileUpload',
		/**
		* 文件上传逻辑（涉及到文件上传均会调用此处）
		* @param {File} sourcesFile 文件对象
		* @param {Function} callback 回调函数，用于回显图片
		*/
		function (sourcesFile, callback) {
			if (/image/i.test(sourcesFile.type)) {
				// 如果上传的是图片，则默认回显base64内容（因为没有图床）
				// 创建 FileReader 对象
				const reader = new FileReader();
				// 读取文件内容
				reader.onload = () => {
					// 获取 base64 内容
					const base64 = reader.result;
					GetPicBed().then((conf) => {
						if (conf.activated === "Base64") {
							callback(base64, {
								name: `${sourcesFile.name.replace(/\.[^.]+$/, '')}`,
								isShadow: true,
								isRadius: false,
							});
						} else {

							return UploadPicbed(file.File.createFrom({
								Mime: sourcesFile.type,
								Name: decodeURIComponent(sourcesFile.name),
								Bytes: base64.split(',')[1],
							}))
						}

					}).then((fileUrl) => {
						callback(fileUrl, {
							name: `${sourcesFile.name.replace(/\.[^.]+$/, '')}`,
							isShadow: true,
							isRadius: false,
						})
					})

				};
				reader.readAsDataURL(sourcesFile);
			}
		})
}