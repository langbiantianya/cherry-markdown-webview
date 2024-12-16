// 参考 https://github.com/Tencent/cherry-markdown/blob/963951e5013e50749c9c2a92f05089465f27af77/src/toolbars/hooks/Image.js#L21

import Cherry from "cherry-markdown"
import { SelectLocalFile } from "../../wailsjs/go/main/App"
import { file } from "../../wailsjs/go/models"

export function InsertMenu() {
	/**
	 * @type {Cherry}
	 */
	let cherryInstance
	return {
		/**
		 *
		 * @param {Cherry} cherry
		 */
		setCherry: (cherry) => {
			cherryInstance = cherry
		},
		localImageMenu: Cherry.createMenuHook("本地图片", {
			iconName: "image",
			onClick: async function (selection) {
				const doc = new file.File()
				doc.DisplayName = "Image file"
				doc.Pattern = "*apng;*.avif;*.gif;*.jpg;*.jpeg;*.jfif;*.pjpeg;*.pjp;*.png;*.svg;*.bmp;*.ico;*.cur;*.tif;*.tiff;*.webp"
				const localFile = await SelectLocalFile(doc)
				let fileName = localFile.Name
				let select = false
				if (selection) {
					fileName = this.getSelection(selection)
					select = true
				}
				if (localFile.Path) {
					cherryInstance.insert(`![${fileName}](${localFile.Path})\n`, select)
				}
			},
		})
	}
}
