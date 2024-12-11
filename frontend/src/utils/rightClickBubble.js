import Cherry from "cherry-markdown";
import { ClipboardSetText, ClipboardGetText } from "../../wailsjs/runtime"
/**
 * 
 * @param {Cherry} cherryInstance 
 */
export const BubbleExtInit = function (cherryInstance) {
	function listenEvent() {

		document.getElementById("markdown-container").addEventListener("mousemove", async function (event) {
			const hooks = cherryInstance.bubble.menus.hooks
			if (cherryInstance.getCodeMirror().getSelection()) {
				hooks["copyMenu"].show()
			} else {
				hooks["copyMenu"].hide()
			}
			if (await ClipboardGetText()) {
				hooks["pasteMenu"].show()
			} else {
				hooks["pasteMenu"].hide()
			}
		})
		// document.getElementById("markdown-container").addEventListener("mouseup", async function (event) {
		// 	const hooks = cherryInstance.bubble.menus.hooks
		// 	if (cherryInstance.getCodeMirror().getSelection()) {
		// 		hooks["copyMenu"].show()
		// 	} else {
		// 		hooks["copyMenu"].hide()
		// 	}
		// 	if (await ClipboardGetText()) {
		// 		hooks["pasteMenu"].show()
		// 	} else {
		// 		hooks["pasteMenu"].hide()
		// 	}
		// })
		cherryInstance.getCodeMirror().on("mousedown", async (instance, event) => {
			const hooks = cherryInstance.bubble.menus.hooks
			if (cherryInstance.getCodeMirror().getSelection()) {
				hooks["copyMenu"].show()
			} else {
				hooks["copyMenu"].hide()
			}
			if (await ClipboardGetText()) {
				hooks["pasteMenu"].show()
			} else {
				hooks["pasteMenu"].hide()
			}
			if (event.button === 2) {
				// 鼠标右键事件
				cherryInstance.bubble.showBubble(
					event.pageY < 57 * 2 ? event.pageY - 57 : event.pageY - 57,
					event.pageX
				);
			}
		});
	}
	(function init() {
		listenEvent()
	})()
}

export const BubbleExtMenu = function () {
	/**
	 * 
	 * @type {Cherry}
	 */
	let cherryInstance

	async function copy() {
		let selectStr = cherryInstance.getCodeMirror().getSelection()
		selectStr && await ClipboardSetText(selectStr)
	}
	async function paste() {
		let str = await ClipboardGetText()
		if (str) {
			cherryInstance.insert(str)
		} else {
			throw new Error("剪贴板为空")
		}
	}
	return {
		/**
		 * 
		 * @param {Cherry} cherry 
		 */
		setCherry: (cherry) => {
			cherryInstance = cherry
		},

		copyMenu: Cherry.createMenuHook("复制", {
			onClick: async () => {
				await copy()
			}
		}),

		pasteMenu: Cherry.createMenuHook("粘贴", {
			onClick: async () => {
				await paste()
			}
		}),
	}

}