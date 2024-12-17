<script>
  import { onMount } from "svelte";
  import "cherry-markdown/dist/cherry-markdown.css";
  import Cherry from "cherry-markdown";
  import { FileMenu, isRelativePath, isRootDirectory } from "./utils/fileMenu";
  import { ExportMenu } from "./utils/exportMenu";
  import { base64ToString } from "./utils/blob";
  import {
    AssociateOpen,
    GetWebServerPort,
    SetSaved,
    GetSaved,
  } from "../wailsjs/go/main/App";
  import { Circle2 } from "svelte-loading-spinners";
  import { EventsOn, Quit } from "../wailsjs/runtime";
  import { BubbleExtInit, BubbleExtMenu } from "./utils/rightClickBubble";
  import { hookbeforeImageMounted, hookFileUpload } from "./utils/fileAnalysis";
  import { InsertMenu } from "./utils/InsertMenu";
  const fileMenu = FileMenu();
  const exportMenu = ExportMenu();
  const bubbleExtMenu = BubbleExtMenu();
  const insertMenu = InsertMenu();
  let loding = $state(true);
  /**
   * 初始化 Cherry Markdown 编辑器实例
   * @type {Cherry}
   */
  let cherryInstance;
  function newCherry(mdStr = "") {
    const cherryInstance = new Cherry({
      id: "markdown-container",
      value: mdStr,
      toolbars: {
        // 定义顶部工具栏
        toolbar: [
          "undo",
          "redo",
          // "|",
          // { fileMenu: ["openFileMenu", "saveFileMenu", "saveAsFileMenu"] },
          // {
          //   exportMenu: ["exportPdfMenu", "exportHtmlMenu"],
          // },
          // "|",
          // 把字体样式类按钮都放在加粗按钮下面
          {
            bold: [
              "bold",
              "italic",
              "underline",
              "strikethrough",
              "sub",
              "sup",
            ],
          },
          "color",
          "size",
          "|",
          "quote",
          "detail",
          "header",
          "list",
          "justify",
          "|",
          // 把插入类按钮都放在插入按钮下面
          {
            insert: [
              "localImageMenu",
              "image",
              // "audio",
              // "video",
              // "pdf",
              // "word",
              // "file",
              "link",
              "hr",
              "br",
              "formula",
              "toc",
              "table",
            ],
          },
          "panel",
        ],
        // 定义侧边栏，默认为空
        sidebar: ["theme", "mobilePreview", "copy"],
        // 定义顶部右侧工具栏，默认为空
        toolbarRight: ["togglePreview"],
        // 定义选中文字时弹出的“悬浮工具栏”，默认为 ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', '|', 'size', 'color']
        bubble: [
          "copyMenu",
          "pasteMenu",
          "|",
          "color",
          "size",
          "|",
          "bold",
          "italic",
          "underline",
          "strikethrough",
          "|",
          "quote",
          "header",
        ],
        // 定义光标出现在行首位置时出现的“提示工具栏”，默认为 ['h1', 'h2', 'h3', '|', 'checklist', 'quote', 'table', 'code']
        float: ["table", "code", "graph"],
        customMenu: {
          copyMenu: bubbleExtMenu.copyMenu,
          pasteMenu: bubbleExtMenu.pasteMenu,
          localImageMenu: insertMenu.localImageMenu,
          // fileMenu: fileMenu.fileMenu,
          // openFileMenu: fileMenu.openFileMenu,
          // saveFileMenu: fileMenu.saveFileMenu,
          // saveAsFileMenu: fileMenu.saveAsFileMenu,
          // exportMenu: exportMenu.exportMenu,
          // exportPdfMenu: exportMenu.exportPdfMenu,
          // exportHtmlMenu: exportMenu.exportHtmlMenu,
        },
        toc: {
          updateLocationHash: false, // 要不要更新URL的hash
          defaultModel: "pure", // pure: 精简模式/缩略模式，只有一排小点； full: 完整模式，会展示所有标题
        },
      },
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
    if (cherryInstance) {
      cherryInstance.destroy();
    }
    cherryInstance = newCherry();
    hookbeforeImageMounted(assciateOpenFile, cherryInstance);
    hookFileUpload(cherryInstance);
    if (
      assciateOpenFile &&
      assciateOpenFile.Path !== "" &&
      assciateOpenFile.Bytes.length > 0
    ) {
      if (typeof assciateOpenFile.Bytes === "string") {
        const mdStr = base64ToString(assciateOpenFile.Bytes);
        cherryInstance.setMarkdown(mdStr);
      }
    }

    loding = false;
    EventsOn("openFileEvent", (event) => {
      fileMenu.openFile();
      // assciateOpenFile = undefined;
    });
    EventsOn("saveFileEvent", async (event) => {
      await fileMenu.saveFile(assciateOpenFile);
    });
    EventsOn("saveAsFileEvent", (event) => {
      fileMenu.saveAsFile();
    });
    EventsOn("exportPdfEvent", (event) => {
      exportMenu.exportPdf();
    });
    EventsOn("exportHtmlEvent", async (event) => {
      await exportMenu.exportHtml();
    });
    EventsOn("quitEvent", async (event) => {
      if (
        cherryInstance.getMarkdown() &&
        !(await GetSaved()) &&
        confirm("是否保存当前文件？")
      ) {
        await fileMenu.saveFile(assciateOpenFile);
      } else {
        await SetSaved(true);
      }
      Quit();
    });
  }

  onMount(() => {
    init();
  });
</script>

<div id="markdown-container" class="h-full"></div>
{#if loding}
  <div
    class="absolute top-0 left-0 right-0 z-20 bg-opacity-75 bg-white dark:bg-gray-800 w-full h-full flex items-center justify-center"
  >
    <Circle2 size="50" unit="vh" />
  </div>
{/if}

<style>
</style>
