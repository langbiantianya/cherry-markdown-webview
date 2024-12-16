import Cherry from "cherry-markdown";
import { GetWebServerPort } from "../../wailsjs/go/main/App";
import { isRelativePath, isRootDirectory } from "./fileMenu";
import { file } from "../../wailsjs/go/models";

/**
 * 
 * @param {Cherry} cherryInstance 
 * @param {file.File} assciateOpenFile 
 */
export async function hookbeforeImageMounted(assciateOpenFile, cherryInstance) {
	let webServerPort = await GetWebServerPort();
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
			try {
				let url = new URL(src);
				console.log("url", src);
				if (url.protocol === "file:") {
					return {
						srcProp,
						src: `http://127.0.0.1:${webServerPort}/file?uri=${url.href}`,
					};
				}
			} catch (e) {
				if (e instanceof TypeError && isRelativePath(src)) {
					return {
						srcProp,
						src: `http://127.0.0.1:${webServerPort}/file?absPath=${src}&docPath=${assciateOpenFile.Path}`,
					};
				} else if (e instanceof TypeError && isRootDirectory(src)) {
					return {
						srcProp,
						src: `http://127.0.0.1:${webServerPort}/file?uri=file:///${src}`,
					};
				}
			}

			return { srcProp, src };
		}
	);
}

/**
 * 
 * @param {Cherry} cherryInstance 
 */
export function hookFileUpload(cherryInstance) {
	cherryInstance.on('fileUpload',
		/**
		* 文件上传逻辑（涉及到文件上传均会调用此处）
		* @param {File} file 具体文件
		* @param {(arg0: string, arg1: any? ,arg2: File?)=>string} callback
		*/
		function (file, callback) {
			if (/image/i.test(file.type)) {
				if(confirm("是否上传到oss")){
					alert("TODO")
				}
				// 如果上传的是图片，则默认回显base64内容（因为没有图床）
				// 创建 FileReader 对象
				const reader = new FileReader();
				// 读取文件内容
				reader.onload = (event) => {
					// 获取 base64 内容
					const base64Content = event.target.result;
					callback(base64Content, {
						name: `${file.name.replace(/\.[^.]+$/, '')}`,
						isShadow: true,
						isRadius: true,
						// width: '30%',
						width: 'auto',
						height: 'auto',
					});
				};
				reader.readAsDataURL(file);
			} else {
				callback('images/demo-dog.png');
			}
		})
}