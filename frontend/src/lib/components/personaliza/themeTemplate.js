import { loadThemeCss } from "$lib/theme"

/**
 * 
 * @param {string} name 
 * @param {string[]} colors 
 * @returns 
 */
export async function applayDiyThemeScss(colors, name = "diy") {
	loadThemeCss(name, await generateScss(colors, name))
}

/**
 * 
 * @param {string[]} colors
 * @param {string} name  
 */
export async function generateScss(colors, name) {
	const baseStyleRequest = await fetch("/theme/editTemplate.scss", {
		method: "GET"
	})
	let baseStyle = await baseStyleRequest.text()
	baseStyle = baseStyle.replaceAll("{{DIY}}", name.toLocaleUpperCase())
	baseStyle = baseStyle.replaceAll("{{diy}}", name)

	colors.forEach((color, index) => {
		switch (index) {
			case 10:
				baseStyle = baseStyle.replaceAll("{{toolbarBg}}", color)
				break
			case 11:
				baseStyle = baseStyle.replaceAll("{{editorBg}}", color)
				break
			default:
				baseStyle = baseStyle.replaceAll(`{{DIY${index}}}`, color)
				break

		}
	})
	return baseStyle
}


/**
 * 
 * @param {string[]} colors
 * @param {string} name  
 */
export async function generateToolbarCss(colors, name) {
	const baseStyleRequest = await fetch("/theme/toolbarTemplate.css", {
		method: "GET"
	})
	let baseStyle = await baseStyleRequest.text()
	baseStyle = baseStyle.replaceAll("{{DIY}}", name.toLocaleUpperCase())
	baseStyle = baseStyle.replaceAll("{{diy}}", name)

	colors.forEach((color, index) => {
		switch (index) {
			case 10:
				baseStyle = baseStyle.replaceAll("{{toolbarBg}}", color)
				break
			case 11:
				baseStyle = baseStyle.replaceAll("{{editorBg}}", color)
				break
			default:
				baseStyle = baseStyle.replaceAll(`{{DIY${index}}}`, color)
				break

		}
	})
	return baseStyle
}