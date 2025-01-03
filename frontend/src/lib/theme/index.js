import { SetActivatedTheme } from '$lib/wailsjs/go/main/App'
import * as sass from 'sass'
import { setTheme } from '@fluentui/web-components';
import { webLightTheme, webDarkTheme } from '@fluentui/tokens';
import frieren from '$lib/images/frieren.jpg';
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
	let flg = false
	for (const styleElement of document.head.getElementsByTagName("style")) {
		if (styleElement.id === name) {
			styleElement.remove()
			styleElement.textContent = theme.css
			flg = true
		}
	}
	if (!flg) {
		const style = document.createElement("style")
		style.textContent = theme.css
		style.id = name
		document.head.appendChild(style)
	}
}

export async function loadDefaultExtTheme() {
	const orangeRequest = await fetch("/theme/orange.scss", {
		method: "GET"
	})
	const orangeSassStr = await orangeRequest.text()
	loadThemeCss("orange", orangeSassStr)
}

export const themeList = () => {
	return [
		{ className: 'default', label: '默认' },
		{ className: 'dark', label: '暗黑' },
		{ className: 'light', label: '明亮' },
		{ className: 'green', label: '清新' },
		{ className: 'red', label: '热情' },
		{ className: 'violet', label: '淡雅' },
		{ className: 'blue', label: '清幽' },
		{ className: 'orange', label: 'citrus' }
	]
}

/**
 * 
 * @param {String} theme 
 */
export function changeMainThemeEvent(theme) {
	SetActivatedTheme(theme);
	switch (theme) {
		case 'dark':
			setTheme(webDarkTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('dark-theme');
			break;
		case 'light':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('light-theme');
			break;
		case 'green':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('green-theme');
			break;
		case 'red':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('red-theme');
			break;
		case 'violet':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('violet-theme');
			break;

		case 'blue':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('blue-theme');
			break;
		case 'orange':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('orange-theme');
			break;
		default:
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('default-theme');
			break;
	}
}


export function loadBackgroundImage() {
	const style = document.createElement("style")
	style.textContent = `
	:root {
		--edit-opacity: 0.9;
		--background-image: url(${frieren});
	}
	`.trim()
	style.id = "background-image"
	document.head.appendChild(style)
}