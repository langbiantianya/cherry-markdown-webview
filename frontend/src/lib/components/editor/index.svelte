<script>
	import { onMount } from 'svelte';
	import 'cherry-markdown/dist/cherry-markdown.css';
	import {
		changeMainThemeEvent,
		clearBackgroundImage,
		isDefault,
		loadBackgroundImage,
		loadDefaultExtTheme,
		loadThemeCss,
		themeList
	} from '$lib/theme';
	import Cherry from 'cherry-markdown';
	import { FileMenu } from '$lib/utils/fileMenu';
	import { ExportMenu } from '$lib/utils/exportMenu';
	import { base64ToString } from '$lib/utils/blob';
	import {
		AssociateOpen,
		SetSaved,
		GetSaved,
		GetActivatedTheme,
		LoadTheme
	} from '$lib/wailsjs/go/main/App';
	import { EventsOn, Quit, WindowSetTitle, EventsOff } from '$lib/wailsjs/runtime';
	import { BubbleExtInit, BubbleExtMenu } from '$lib/utils/rightClickBubble';
	import { hookbeforeImageMounted, hookFileUpload } from '$lib/utils/fileAnalysis';
	import { globalState } from '$lib/store';
	import { InsertMenu } from '$lib/utils/InsertMenu';
	import { file } from '$lib/wailsjs/go/models';

	const fileMenu = FileMenu();
	const exportMenu = ExportMenu();
	const bubbleExtMenu = BubbleExtMenu();
	const insertMenu = InsertMenu();
	/**DF
	 * 初始化 Cherry Markdown 编辑器实例
	 * @type {Cherry}
	 */
	let cherryInstance;
	/**
	 *
	 * @param mdStr
	 * @param {file.File|undefined|null} assciateOpenFile
	 * @param {import('cherry-markdown/types/cherry').EditorMode} model
	 */
	async function newCherry(mdStr = '', assciateOpenFile = null, model = 'edit&preview') {
		const cherryInstance = new Cherry({
			id: 'markdown-container',
			editor: {
				defaultModel: model
			},
			value: mdStr,
			toolbars: {
				// shortcutKeySettings: {
				// 	/** 是否替换已有的快捷键, true: 替换默认快捷键； false： 会追加到默认快捷键里，相同的shortcutKey会覆盖默认的 */
				// 	isReplace: false,
				// 	shortcutKeyMap: {}
				// },
				// 定义顶部工具栏
				toolbar: [
					'undo',
					'redo',
					// 把字体样式类按钮都放在加粗按钮下面
					{
						bold: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup']
					},
					'color',
					'size',
					'|',
					'quote',
					'detail',
					'header',
					'list',
					'justify',
					'|',
					// 把插入类按钮都放在插入按钮下面
					{
						insert: [
							'localImageMenu',
							'image',
							// "audio",
							// "video",
							// "pdf",
							// "word",
							// "file",
							'link',
							'hr',
							'br',
							'formula',
							'toc',
							'table'
						]
					},
					'panel'
				],
				// 定义侧边栏，默认为空
				sidebar: ['theme', 'mobilePreview', 'copy'],
				// 定义顶部右侧工具栏，默认为空
				toolbarRight: ['togglePreview', 'codeTheme', 'shortcutKey'],
				// 定义选中文字时弹出的“悬浮工具栏”，默认为 ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', '|', 'size', 'color']
				bubble: [
					'copyMenu',
					'pasteMenu',
					'|',
					'color',
					'size',
					'|',
					'bold',
					'italic',
					'underline',
					'strikethrough',
					'|',
					'quote',
					'header'
				],
				// 定义光标出现在行首位置时出现的“提示工具栏”，默认为 ['h1', 'h2', 'h3', '|', 'checklist', 'quote', 'table', 'code']
				float: ['table', 'code', 'graph'],
				customMenu: {
					copyMenu: bubbleExtMenu.copyMenu,
					pasteMenu: bubbleExtMenu.pasteMenu,
					localImageMenu: insertMenu.localImageMenu
				},
				toc: {
					updateLocationHash: false, // 要不要更新URL的hash
					defaultModel: 'pure' // pure: 精简模式/缩略模式，只有一排小点； full: 完整模式，会展示所有标题
				}
			},
			callback: {
				/**
				 *
				 * @param {string} srcProp
				 * @param {string} src
				 */
				beforeImageMounted: await hookbeforeImageMounted(assciateOpenFile),
				/**
				 * 文件上传逻辑（涉及到文件上传均会调用此处）
				 * @param {File} sourcesFile 文件对象
				 * @param {Function} callback 回调函数，用于回显图片
				 */
				fileUpload: hookFileUpload()
			},
			event: {
				changeMainTheme: changeMainThemeEvent
			},
			previewer: {},
			themeSettings: {
				// 主题列表，用于切换主题
				themeList: await themeList(),
				mainTheme: 'default',
				codeBlockTheme: 'default',
				inlineCodeTheme: 'black', // red or black
				toolbarTheme: 'dark' // light or dark 优先级低于mainTheme
			}
		});
		const activityThemeName = await GetActivatedTheme();
		const activietTheme = (await themeList())
			.filter((item) => {
				return item.label === activityThemeName;
			})
			.pop();
		cherryInstance.setTheme(activietTheme?.className || 'default');
		fileMenu.setCherry(cherryInstance);
		exportMenu.setCherry(cherryInstance);
		bubbleExtMenu.setCherry(cherryInstance);
		insertMenu.setCherry(cherryInstance);
		BubbleExtInit(cherryInstance);
		return cherryInstance;
	}
	/**
	 * 计算文本的行数
	 * @param {string} text - 要计算行数的文本
	 * @returns {number} - 文本的行数
	 */
	function countLines(text) {
		const lineCount = (text.match(/\n/g) || []).length + 1;
		return lineCount;
	}
	async function init() {
		let assciateOpenFile = await AssociateOpen();
		cherryInstance && cherryInstance.destroy();
		if (assciateOpenFile && assciateOpenFile.Path !== '' && assciateOpenFile.Bytes.length > 0) {
			if (typeof assciateOpenFile.Bytes === 'string') {
				WindowSetTitle(assciateOpenFile.Name);
				const mdStr = base64ToString(assciateOpenFile.Bytes);
				const lines = countLines(mdStr);
				if (lines > 50000) {
					alert('文件过大,使用异步加载会导致撤销变成空白！！！');
					if (confirm('文件过大，是否仅使用编辑模式？')) {
						cherryInstance = await newCherry('', null, 'editOnly');
					} else {
						cherryInstance = await newCherry();
					}
					cherryInstance.setMarkdown(mdStr);
				} else {
					cherryInstance = await newCherry(mdStr, assciateOpenFile);
				}
			}
		} else {
			cherryInstance = await newCherry();
		}
		try {
			EventsOff(
				'openFileEvent',
				'saveFileEvent',
				'saveAsFileEvent',
				'exportPdfEvent',
				'exportHtmlEvent',
				'quitEvent'
			);
			EventsOn('openFileEvent', (event) => {
				fileMenu.openFile();
				// assciateOpenFile = undefined;
			});
			EventsOn('saveFileEvent', async (event) => {
				const res = await fileMenu.saveFile(assciateOpenFile);
				res && (assciateOpenFile = res) && WindowSetTitle(assciateOpenFile.Name);
				await SetSaved(true);
			});
			EventsOn('saveAsFileEvent', async (event) => {
				await fileMenu.saveAsFile();
				await SetSaved(true);
			});
			EventsOn('exportPdfEvent', (event) => {
				exportMenu.exportPdf();
			});
			EventsOn('exportHtmlEvent', async (event) => {
				await exportMenu.exportHtml();
			});
			EventsOn('quitEvent', async (event) => {
				if (cherryInstance.getMarkdown() && !(await GetSaved()) && confirm('是否保存当前文件？')) {
					await fileMenu.saveFile(assciateOpenFile);
				} else {
					await SetSaved(true);
				}
				Quit();
			});
		} catch (error) {
			console.error(error);
		}
		$globalState.loading = false;
	}

	onMount(async () => {
		await loadDefaultExtTheme();
		await init();
	});
</script>

<svelte:head></svelte:head>

<div class="h-lvh">
	<div id="markdown-container" class="h-full"></div>
</div>

<style>
</style>
