import { GetThemeList, LoadTheme, SetActivatedTheme } from '$lib/wailsjs/go/main/App'
import * as sass from 'sass'
import { setTheme } from '@fluentui/web-components';
import { webLightTheme, webDarkTheme } from '@fluentui/tokens';
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
	for (const styleElement of document.head.getElementsByTagName("style")) {
		if (styleElement.id === name) {
			styleElement.remove()
		}
	}

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

export const defaultThemeList = [
	{ className: 'default', label: '默认' },
	{ className: 'dark', label: '暗黑' },
	{ className: 'light', label: '明亮' },
	{ className: 'green', label: '清新' },
	{ className: 'red', label: '热情' },
	{ className: 'violet', label: '淡雅' },
	{ className: 'blue', label: '清幽' },
	{ className: 'orange', label: 'citrus' }
]
/**
 * 
 * @param {string} className 
 * @returns 
 */
export function isDefault(className) {
	return defaultThemeList.map(item => item.className).includes(className)
}

export const themeList = async () => {
	return [...defaultThemeList, ...(await GetThemeList())]
}

/**
 * 
 * @param {String} theme 
 */
export function changeMainThemeEvent(theme) {
	SetActivatedTheme(theme);
	clearBackgroundImage();
	switch (theme) {
		case 'default':
			setTheme(webLightTheme);
			document.body.classList.remove(...document.body.classList.values());
			document.body.classList.toggle('default-theme');
			break;
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
			themeList().then(thems => {
				return thems.filter(item => { return item.className === theme }).pop()
			}).then(res => {
				if (res) {
					return LoadTheme(res.label)
				}
				return null
			}).then(res => {
				if (res) {
					loadThemeCss(res.name.className, res.scss);
					loadToolBarCss(res.name.className, res.toolBar)
					document.body.classList.toggle(`${res.name.className}-theme`);
					if (res.backgroundImage) {
						loadBackgroundImage(
							`data:${res.backgroundImage?.Mime};base64,${res.backgroundImage?.Bytes}`
						);
					}
				}
			})
			break;
	}
}

/**
 * 
 * @param {string} bgimg 
 */
export function loadBackgroundImage(bgimg) {

	for (const styleElement of document.head.getElementsByTagName("style")) {
		if (styleElement.id === "background-image") {
			styleElement.remove()
		}
	}

	const style = document.createElement("style")
	style.textContent = `
	:root {
		--edit-opacity: 0.9;
		--background-image: url(${bgimg});
	}
	`.trim()
	style.id = "background-image"
	document.head.appendChild(style)
}

export function clearBackgroundImage() {
	for (const styleElement of document.head.getElementsByTagName("style")) {
		if (styleElement.id === "background-image") {
			styleElement.remove()
		}
	}
}

/**
 * 
 * @param {string} name 
 * @param {string} cssStr 
 */
export function loadToolBarCss(name, cssStr) {

	for (const styleElement of document.head.getElementsByTagName("style")) {
		if (styleElement.id === name + "-theme") {
			styleElement.remove()
		}
	}

	const style = document.createElement("style")
	style.textContent = cssStr
	style.id = name + "-theme"
	document.head.appendChild(style)
}