<script>
	import { onMount } from 'svelte';
	import 'cherry-markdown/dist/cherry-markdown.css';
	import Cherry from 'cherry-markdown';
	import { FileMenu } from '$lib/utils/fileMenu';
	import { ExportMenu } from '$lib/utils/exportMenu';
	import { base64ToString } from '$lib/utils/blob';
	import { AssociateOpen, SetSaved, GetSaved } from '$lib/wailsjs/go/main/App';
	import { EventsOn, Quit, WindowSetTitle, EventsOff } from '$lib/wailsjs/runtime';
	import { BubbleExtInit, BubbleExtMenu } from '$lib/utils/rightClickBubble';
	import { hookbeforeImageMounted, hookFileUpload } from '$lib/utils/fileAnalysis';
	import { globalState } from '$lib/store';
	import { InsertMenu } from '$lib/utils/InsertMenu';

	const fileMenu = FileMenu();
	const exportMenu = ExportMenu();
	const bubbleExtMenu = BubbleExtMenu();
	const insertMenu = InsertMenu();
	/**
	 * 初始化 Cherry Markdown 编辑器实例
	 * @type {Cherry}
	 */
	let cherryInstance;
	function newCherry(mdStr = '') {
		const cherryInstance = new Cherry({
			id: 'markdown-container',
			value: mdStr,
			toolbars: {
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
				toolbarRight: ['togglePreview'],
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
			}
		});
		fileMenu.setCherry(cherryInstance);
		exportMenu.setCherry(cherryInstance);
		bubbleExtMenu.setCherry(cherryInstance);
		insertMenu.setCherry(cherryInstance);
		BubbleExtInit(cherryInstance);
		return cherryInstance;
	}

	async function init() {
		let assciateOpenFile = await AssociateOpen();
		if (!cherryInstance) {
			cherryInstance = newCherry();
			hookbeforeImageMounted(assciateOpenFile, cherryInstance);
			hookFileUpload(cherryInstance);
			if (assciateOpenFile && assciateOpenFile.Path !== '' && assciateOpenFile.Bytes.length > 0) {
				if (typeof assciateOpenFile.Bytes === 'string') {
					WindowSetTitle(assciateOpenFile.Name);
					const mdStr = base64ToString(assciateOpenFile.Bytes);
					cherryInstance.setMarkdown(mdStr);
				}
			}
			$globalState.loading = false;
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
					if (
						cherryInstance.getMarkdown() &&
						!(await GetSaved()) &&
						confirm('是否保存当前文件？')
					) {
						await fileMenu.saveFile(assciateOpenFile);
					} else {
						await SetSaved(true);
					}
					Quit();
				});
			} catch (error) {
				console.error(error);
			}
		}
	}

	onMount(() => {
		init();
	});
</script>

<svelte:head></svelte:head>

<div class="h-lvh">
	<div id="markdown-container" class="h-full"></div>
</div>

<style>
</style>
