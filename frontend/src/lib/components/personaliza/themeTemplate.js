import { loadThemeCss } from "$lib/theme"

/**
 * 
 * @param {string} name 
 * @param {string[]} colors 
 * @returns 
 */
export async function ApplayDiyThemeScss(name, colors) {

	const baseStyleRequest = await fetch("/theme/template.scss", {
		method: "GET"
	})
	let baseStyle = await baseStyleRequest.text()
	baseStyle = baseStyle.replaceAll("{{DIY}}", name)

	colors.forEach((color, index) => {
		baseStyle = baseStyle.replaceAll(`{{DIY${index}}`, color)
	})

	baseStyle
	loadThemeCss("diy", baseStyle)
}