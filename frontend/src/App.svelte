<script>
  // import {Greet} from '../wailsjs/go/main/App.js'
  import { onMount } from "svelte";
  import "cherry-markdown/dist/cherry-markdown.css";
  import Cherry from "cherry-markdown";

  // fileHandle is an instance of FileSystemFileHandle..
  async function writeFile(fileHandle, contents) {
    // Create a FileSystemWritableFileStream to write to.
    const writable = await fileHandle.createWritable();
    // Write the contents of the file to the stream.
    await writable.write(contents);
    // Close the file and write the contents to disk.
    await writable.close();
  }

  let fileHandle;
  onMount(() => {
    var fileMenu = Cherry.createMenuHook("文件", {});

    var openFileMenu = Cherry.createMenuHook("打开", {
      onClick: async () => {
        [fileHandle] = await window.showOpenFilePicker({
          types: [
            {
              accept: {
                "text/plain": [".md", ".mdx"],
              },
            },
          ],
          multiple: false,
        });
        const fileData = await fileHandle.getFile();
        console.log(await fileData.text());

        cherryInstance.setMarkdown(await fileData.text());
      },
    });
    var saveAsFileMenu = Cherry.createMenuHook("另存为", {
      onClick: async () => {
        const savefileHandle = await window.showSaveFilePicker({
          types: [
            {
              description: "Markdown file",
              accept: {
                "text/plain": [".md", ".mdx"],
              },
            },
          ],
        });
        writeFile(savefileHandle, cherryInstance.getMarkdown());
      },
    });
    var saveFileMenu = Cherry.createMenuHook("保存", {
      onClick: async () => {
        if (fileHandle) {
          writeFile(fileHandle, cherryInstance.getMarkdown());
        } else {
          const savefileHandle = await window.showSaveFilePicker({
            types: [
              {
                description: "Markdown file",
                accept: {
                  "text/plain": [".md", ".mdx"],
                },
              },
            ],
          });
          writeFile(savefileHandle, cherryInstance.getMarkdown());
        }
      },
    });

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
          "export",
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
          fileMenu,
          openFileMenu,
          saveFileMenu,
          saveAsFileMenu,
        },
      },
    });

    document.getElementsByClassName("cherry-toolbar-文件")[0];

    document.getElementsByName("fileMenu").forEach((item) => {
      item.style.left = "65";
    });
  });
</script>

<div id="markdown-container" class="h-full"></div>

<style>
</style>
