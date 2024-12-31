import * as sass from 'sass'

/**
 * @param {String} name 
 * @param {String} sassStr 
 */
export async function loadThemeCss(name, sassStr) {
	const baseStyleRequest = await fetch("/theme/variable.scss", {
		method: "GET"
	})
	const baseStyle = await baseStyleRequest.text()
	const theme = sass.compileString(baseStyle + "\n" + sassStr)
	const style = document.createElement("style")
	style.textContent = theme.css
	style.id = name
	document.head.appendChild(style)
}

export async function loadDefaultExtTheme() {
	const orangeRequest = await fetch("/theme/orange.scss", {
		method: "GET"
	})
	const orangeSassStr = await orangeRequest.text()
	loadThemeCss("orange", orangeSassStr)
}