<script>
  // import {Greet} from '../wailsjs/go/main/App.js'
  import { onMount } from "svelte";
  import "cherry-markdown/dist/cherry-markdown.css";
  import Cherry from "cherry-markdown";
  import { FileMenu } from "./utils/fileMenu";
  import { ExportMenu } from "./utils/exportMenu";

  const fileMenu = FileMenu();
  const exportMenu = ExportMenu();
  onMount(() => {
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
              "ruby",
            ],
          },
          "color",
          "size",
          "|",
          "quote",
          "detail",
          "header",
          "ul",
          "ol",
          "checklist",
          "list",
          "justify",
          "panel",
          "|",
          "code",
          // 把插入类按钮都放在插入按钮下面
          {
            insert: [
              "image",
              "audio",
              "video",
              "pdf",
              "word",
              "file",
              "link",
              "hr",
              "br",
              "formula",
              "toc",
              "table",
              "drawIo",
            ],
          },
          "graph",
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
      },
    });
    fileMenu.setCherry(cherryInstance);
    exportMenu.setCherry(cherryInstance);
  });
</script>

<div id="markdown-container" class="h-full"></div>

<style>
</style>
