<script>
	import '@fluentui/web-components/button.js';
	import { onMount } from 'svelte';
	import Cherry from 'cherry-markdown';
	import { previewText } from './previewText';
	import { themeList } from '$lib/theme';
	import { GetActivatedTheme } from '$lib/wailsjs/go/main/App';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';
	import iro from '@jaames/iro';
	import { TextInput } from '@fluentui/web-components';
	/**
	 * @type {{heid:function():void}}
	 */
	let { heid } = $props();

	function goBack() {
		// TODO 重置主题为激活的主题
		heid();
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

	function hiddenColorPicker() {
		picker.hidden = true;
	}
	/**
	 * @param {Element}event
	 */
	function inputFildChange(event) {
		console.log(event);
	}
	onMount(async () => {
		colorPicker = iro.ColorPicker(picker, {
			width: 120
		});
		for (const element of themeInput.getElementsByTagName('fluent-text-input')) {
			if (element.id !== 'diy-theme-name' && element instanceof HTMLElement) {
				pickerClick(element);
			}
		}

		const cherrymarkdown = new Cherry({
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
				themeList: [...themeList(), { className: 'diy', label: 'diy' }],
				mainTheme: 'default',
				codeBlockTheme: 'default',
				inlineCodeTheme: 'black', // red or black
				toolbarTheme: 'dark' // light or dark 优先级低于mainTheme
			}
		});
		cherrymarkdown.setTheme((await GetActivatedTheme()) || 'default');

		for (const element of editor.getElementsByClassName('cherry-toolbar')) {
			const mask = document.createElement('div');
			mask.style.position = 'absolute';
			mask.style.height = '3rem';
			mask.style.width = '100%';
			mask.style.zIndex = '2';
			element.append(mask);
		}
		for (const code of editor.getElementsByClassName('CodeMirror')) {
			for (const textarea of code.getElementsByTagName('textarea')) {
				textarea.readOnly = true;
				textarea.tabIndex = -1;
				textarea.disabled = true;
			}
		}
	});
</script>

<div class="bg-base m-0 h-full overflow-hidden p-0">
	<div part="header" class="bg-toolBar h-12 w-full px-4 shadow-sm">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div class="text-toolBar flex h-12 w-full flex-nowrap items-center justify-between">
			<h1 class="text-2xl">个性化</h1>
			<!-- svelte-ignore event_directive_deprecated -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<fluent-button
				class="hover:bg-buttonHover hover:text-buttonHover bg-button text-button h-9 rounded-none leading-9"
				on:click={goBack}>返回</fluent-button
			>
		</div>
	</div>
	<div
		class="drawer-contentw-full flex h-lvh flex-wrap space-y-4 overflow-y-scroll px-1 pb-10 align-top"
	>
		<div bind:this={themeInput} class="theme-input flex flex-wrap space-x-1">
			<div class="pl-1">
				<fluent-label>主题名称</fluent-label>
				<fluent-text-input on:blur={hiddenColorPicker} id="diy-theme-name" appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color0</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					id="color0"
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color1</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color2</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color3</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color4</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color5</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color6</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color7</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color8</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
			</div>
			<div>
				<fluent-label>color9</fluent-label>
				<fluent-text-input
					on:blur={hiddenColorPicker}
					on:change={inputFildChange}
					appearance="outline"
				></fluent-text-input>
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
</div>

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
		padding-bottom: 1rem;
		/* min-width: 37rem;
		max-width: 57rem; */
	}
	@media (max-width: 768px) {
		.theme-input {
			min-width: 100lvw;
		}
	}
	@media (min-width: 1080px) {
		.theme-input {
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
		}
		.theme-input {
			max-width: 25rem;
			height: 27rem;
		}
	}
</style>
