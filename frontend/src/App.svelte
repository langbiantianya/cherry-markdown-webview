<script>
  import { onMount } from "svelte";
  import "cherry-markdown/dist/cherry-markdown.css";
  import Cherry from "cherry-markdown";
  import { FileMenu } from "./utils/fileMenu";
  import { ExportMenu } from "./utils/exportMenu";
  import { arrayToBlob, blobToString, base64ToString } from "./utils/blob";
  import { AssociateOpen } from "../wailsjs/go/main/App";

  import { Circle2 } from "svelte-loading-spinners";

  const fileMenu = FileMenu();
  const exportMenu = ExportMenu();
  let loding = $state(true);
  onMount(async () => {
    const cherryInstance = new Cherry({
      id: "markdown-container",
      value: "",
      toolbars: {
        // 定义顶部工具栏
        toolbar: [
          "undo",
          "redo",
          "|",
          { fileMenu: ["openFileMenu", "saveFileMenu", "saveAsFileMenu"] },
          {
            exportMenu: ["exportPdfMenu", "exportHtmlMenu"],
          },
          "|",
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
              // "image",
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
          "color",
          "size",
          "|",
          "bold",
          "italic",
          "underline",
          "strikethrough",
          "sub",
          "sup",
          "|",
          "quote",
          "header",
          "ul",
          "ol",
          "checklist",
        ],
        // 定义光标出现在行首位置时出现的“提示工具栏”，默认为 ['h1', 'h2', 'h3', '|', 'checklist', 'quote', 'table', 'code']
        float: ["table", "code", "graph"],
        customMenu: {
          fileMenu: fileMenu.fileMenu,
          openFileMenu: fileMenu.openFileMenu,
          saveFileMenu: fileMenu.saveFileMenu,
          saveAsFileMenu: fileMenu.saveAsFileMenu,
          exportMenu: exportMenu.exportMenu,
          exportPdfMenu: exportMenu.exportPdfMenu,
          exportHtmlMenu: exportMenu.exportHtmlMenu,
        },
        toc: {
          updateLocationHash: false, // 要不要更新URL的hash
          defaultModel: "pure", // pure: 精简模式/缩略模式，只有一排小点； full: 完整模式，会展示所有标题
        },
      },
    });
    fileMenu.setCherry(cherryInstance);
    exportMenu.setCherry(cherryInstance);
    const assciateOpenFile = await AssociateOpen();
    if (
      assciateOpenFile &&
      assciateOpenFile.Name !== "" &&
      assciateOpenFile.Bytes.length > 0
    ) {
      if (typeof assciateOpenFile.Bytes === "string") {
        const mdStr = base64ToString(assciateOpenFile.Bytes);
        cherryInstance.setMarkdown(mdStr);
      } else if (typeof assciateOpenFile.Bytes === "object") {
        const mdStr = await blobToString(arrayToBlob(assciateOpenFile.Bytes));
        cherryInstance.setMarkdown(mdStr);
        loding = false;
      }
    } else {
      loding = false;
    }
    loding = false;
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
