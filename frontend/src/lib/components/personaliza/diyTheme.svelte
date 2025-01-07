<script>
	import '@fluentui/web-components/button.js';
	import { onMount } from 'svelte';
	import Cherry from 'cherry-markdown';
	import { previewText } from './previewText';
	import { themeList } from '$lib/theme';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';
	import iro from '@jaames/iro';
	import { TextInput } from '@fluentui/web-components';
	import { applayDiyThemeScss, generateScss, generateToolbarCss } from './themeTemplate';
	import { UpsertThemeItem } from '$lib/wailsjs/go/main/App';
	import { config } from '$lib/wailsjs/go/models';
	/**
	 * @type {string[]}
	 */
	const colors = [
		'#ebfbee',
		'#d3f9d8',
		'#b2f2bb',
		'#8ce99a',
		'#69db7c',
		'#51cf66',
		'#40c057',
		'#37b24d',
		'#2f9e44',
		'#2b8a3e',
		'#ffffff',
		'#ffffff'
	];
	function resertColors() {
		[
			'#ebfbee',
			'#d3f9d8',
			'#b2f2bb',
			'#8ce99a',
			'#69db7c',
			'#51cf66',
			'#40c057',
			'#37b24d',
			'#2f9e44',
			'#2b8a3e',
			'#ffffff',
			'#ffffff'
		].forEach((element, index) => {
			colors[index] = element;
		});
		applayDiyThemeScss(colors);
	}
	/**
	 * @param {HTMLElement} element
	 */
	function pickerClick(element) {
		element.addEventListener('click', (event) => {
			picker.style.top = event.pageY + 'px';
			picker.style.left = event.pageX - 60 + 'px';
			colorPicker.off('color:change', () => {});
			colorPicker.on(
				'color:change',
				/**
				 *
				 * @param {iro.Color} color
				 */
				(color) => {
					//将当前颜色记录为十六进制字符串
					if (event.target instanceof TextInput) {
						event.target['currentValue'] = color.hexString;
						event.target.dispatchEvent(new Event('change'));
					}
				}
			);
			picker.hidden = false;
		});
	}
	async function saveTheme() {
		let regex = /^[A-Za-z0-9]*$/;
		if (!diyLabelThemeName.currentValue) {
			alert('请输入主题名称');
			return;
		} else if (!diyClassThemeName.currentValue) {
			alert('请输入主题类名');
			return;
		} else if (!regex.test(diyClassThemeName.currentValue)) {
			alert('主题类名只能使用英文和数字');

			return;
		}

		const theme = new config.ThemeItem({
			name: new config.ExtTheme({
				className: diyClassThemeName.currentValue,
				label: diyLabelThemeName.currentValue
			}),
			colors: colors,
			toolBar: await generateToolbarCss(colors, diyClassThemeName.currentValue),
			scss: await generateScss(colors, diyClassThemeName.currentValue),
			backgroundImage: null
		});
		UpsertThemeItem(theme);
	}
	/**
	 * @param {Event} event
	 */
	function verificationDiyClassThemeName(event) {
		let regex = /^[A-Za-z0-9]*$/;
		if (event.target instanceof TextInput) {
			event.target.currentValueChanged = (prev, next) => {
				if (!regex.test(next) && event.target instanceof TextInput) {
					alert('主题类名只能使用英文和数字');
				}
			};
		}
	}
	/**
	 * @type {TextInput}
	 */
	let diyLabelThemeName;
	/**
	 * @type {TextInput}
	 */
	let diyClassThemeName;
	/**
	 * @type {HTMLElement}
	 */
	let picker;
	/**
	 * @type {HTMLElement}
	 */
	let themeInput;
	/**
	 * @type {iro.ColorPicker}
	 */
	let colorPicker;
	/**
	 * @type {HTMLElement}
	 */
	let editor;
	/**
	 * @type {Cherry}
	 */
	let cherrymarkdown;
	function hiddenColorPicker() {
		picker.hidden = true;
	}
	/**
	 * @param {Event}event
	 */
	async function inputFildChange(event) {
		if (event.target instanceof TextInput) {
			switch (event.target.id) {
				case 'toolbarbg':
					colors[10] = event.target.currentValue;
					break;
				case 'editorBg':
					colors[11] = event.target.currentValue;
					break;
				case 'color0':
					colors[0] = event.target.currentValue;
					break;
				case 'color1':
					colors[1] = event.target.currentValue;
					break;
				case 'color2':
					colors[2] = event.target.currentValue;
					break;
				case 'color3':
					colors[3] = event.target.currentValue;
					break;
				case 'color4':
					colors[4] = event.target.currentValue;
					break;
				case 'color5':
					colors[5] = event.target.currentValue;
					break;
				case 'color6':
					colors[6] = event.target.currentValue;
					break;
				case 'color7':
					colors[7] = event.target.currentValue;
					break;
				case 'color8':
					colors[8] = event.target.currentValue;
					break;
				case 'color9':
					colors[9] = event.target.currentValue;
					break;
			}
			await applayDiyThemeScss(colors);
			cherrymarkdown && cherrymarkdown.setTheme('diy');
		}
	}
	onMount(async () => {
		colorPicker = iro.ColorPicker(picker, {
			width: 120
		});
		for (const element of themeInput.getElementsByTagName('fluent-text-input')) {
			if (!element.id.includes('theme-name') && element instanceof HTMLElement) {
				pickerClick(element);
			}
		}

		cherrymarkdown = new Cherry({
			value: previewText,
			id: 'markdown-personaliza-preview',
			editor: {
				defaultModel: 'edit&preview'
			},
			toolbars: {
				shortcutKeySettings: {
					isReplace: true
				},
				float: false,
				bubble: false
			},
			themeSettings: {
				// 主题列表，用于切换主题
				themeList: [...(await themeList()), { className: 'diy', label: 'diy' }],
				mainTheme: 'default',
				codeBlockTheme: 'default',
				inlineCodeTheme: 'black', // red or black
				toolbarTheme: 'dark' // light or dark 优先级低于mainTheme
			}
		});
		await applayDiyThemeScss(colors);
		cherrymarkdown.setTheme('diy');
	});
</script>

<div
	class="drawer-contentw-full flex h-lvh flex-wrap space-y-4 overflow-y-scroll px-1 pb-10 align-top"
>
	<div bind:this={themeInput} class="theme-input flex flex-wrap space-x-1">
		<div class="ml-1 w-52">
			<fluent-label>主题名称</fluent-label>
			<fluent-text-input
				bind:this={diyLabelThemeName}
				on:blur={hiddenColorPicker}
				id="diy-label-theme-name"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>主题类名</fluent-label>
			<fluent-text-input
				on:input={verificationDiyClassThemeName}
				bind:this={diyClassThemeName}
				on:blur={hiddenColorPicker}
				id="diy-class-theme-name"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>toolbarbg</fluent-label>
			<fluent-text-input
				initialValue={colors[10]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				id="toolbarbg"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>editorBg</fluent-label>
			<fluent-text-input
				initialValue={colors[11]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				id="editorBg"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color0</fluent-label>
			<fluent-text-input
				initialValue={colors[0]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				id="color0"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color1</fluent-label>
			<fluent-text-input
				initialValue={colors[1]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				id="color1"
				appearance="outline"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color2</fluent-label>
			<fluent-text-input
				initialValue={colors[2]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color2"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color3</fluent-label>
			<fluent-text-input
				initialValue={colors[3]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color3"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color4</fluent-label>
			<fluent-text-input
				initialValue={colors[4]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color4"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color5</fluent-label>
			<fluent-text-input
				initialValue={colors[5]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color5"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color6</fluent-label>
			<fluent-text-input
				initialValue={colors[6]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color6"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color7</fluent-label>
			<fluent-text-input
				initialValue={colors[7]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color7"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color8</fluent-label>
			<fluent-text-input
				initialValue={colors[8]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color8"
			></fluent-text-input>
		</div>
		<div class="w-52">
			<fluent-label>color9</fluent-label>
			<fluent-text-input
				initialValue={colors[9]}
				on:blur={hiddenColorPicker}
				on:change={inputFildChange}
				appearance="outline"
				id="color9"
			></fluent-text-input>
		</div>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="w-52 space-x-3">
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<fluent-button
				class="hover:bg-buttonHover hover:text-buttonHover bg-button text-button mt-5"
				size="medium"
				type="reset"
				on:click={resertColors}>重置</fluent-button
			>
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<fluent-button
				class="hover:bg-buttonHover hover:text-buttonHover bg-button text-button mt-5"
				size="medium"
				on:click={saveTheme}>保存</fluent-button
			>
		</div>
	</div>
	<div class="edit h-96">
		<div
			bind:this={editor}
			class="mx-auto h-96 w-full"
			tabindex="-1"
			id="markdown-personaliza-preview"
		></div>
	</div>
</div>
<div bind:this={picker} hidden class=" fixed right-8 top-10 z-10 w-32"></div>

<style>
	/* 重写 select 选项的样式 */
	.accordion-item-head {
		color: var(--color-text-base);
	}
	fluent-tab {
		color: var(--color-text-base);
	}
	fluent-label {
		color: var(--color-text-base);
	}
	fluent-tab {
		--colorCompoundBrandStroke: var(--color-text-base);
	}
	fluent-text-input {
		--colorNeutralForeground1: var(--color-text-base);
		--colorCompoundBrandStroke: var(--color-text-base);
		--colorCompoundBrandStrokePressed: var(--color-text-base);
		--colorNeutralBackground1: var(--color-bg-input);
	}

	.drawer-contentw-full {
		height: calc(100lvh +200px);
	}
	.edit {
		padding-bottom: 4rem;
		/* min-width: 37rem;
		max-width: 57rem; */
	}
	@media (max-width: 768px) {
		.theme-input {
			min-width: 100lvw;
		}
	}
	@media (min-width: 1280px) {
		.theme-input {
			min-width: 27rem;
			justify-content: flex-start; /* 水平居中 */
			align-items: flex-start;
		}
		.drawer-contentw-full {
			/* align-items: center;  */
			justify-content: center; /* 水平居中 */
		}
		.edit {
			width: 47rem;
			height: 27rem;
			padding-bottom: 1.5rem;
		}
		.theme-input {
			max-width: 25rem;
			height: 27rem;
		}
	}
</style>
