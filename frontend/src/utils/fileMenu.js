import Cherry from "cherry-markdown";
import { getOpenFileHandle, getSaveFileHandle, writeFile } from "./io";

export const FileMenu = function () {
  let fileHandle;
  let cherryInstance;
  return {
    /**
     *
     * @param {Cherry} cherry
     */
    setCherry: (cherry) => {
      cherryInstance = cherry;
    },
    fileMenu: Cherry.createMenuHook("文件", {}),
    openFileMenu: Cherry.createMenuHook("打开", {
      onClick: async () => {
        fileHandle = await getOpenFileHandle();
        const fileData = await fileHandle.getFile();
        cherryInstance.setMarkdown(await fileData.text());
      },
    }),
    saveAsFile: async () => {
      const savefileHandle = await getSaveFileHandle();
      writeFile(savefileHandle, cherryInstance.getMarkdown());
    },
    saveAsFileMenu: Cherry.createMenuHook("另存为", {
      onClick: async () => {
        await this.saveAsFile()
      },
    }),
    saveFile: async () => {
      if (fileHandle) {
        writeFile(fileHandle, cherryInstance.getMarkdown());
      } else {
        const savefileHandle = await getSaveFileHandle();
        writeFile(savefileHandle, cherryInstance.getMarkdown());
      }
    },
    saveFileMenu: Cherry.createMenuHook("保存", {
      onClick: async () => {
        await this.saveFile()
      },
    }),

  };
};

