import Cherry from "cherry-markdown";
import { getOpenFileHandle, getSaveFileHandle, writeFile } from "./io";
import { file } from "../../wailsjs/go/models";
import { SaveFile, OpenFile } from "../../wailsjs/go/main/App";
import { stringToBinaryArray } from "./blob";

export const FileMenu = function () {
  let fileHandle
  let cherryInstance
  return {
    /**
     *
     * @param {Cherry} cherry
     */
    setCherry: (cherry) => {
      cherryInstance = cherry
    },
    // fileMenu: Cherry.createMenuHook("文件", {}),
    openFile: async () => {
      // if (window.showOpenFilePicker) {
      //   fileHandle = await getOpenFileHandle()
      //   const fileData = await fileHandle.getFile()
      //   cherryInstance.setMarkdown(await fileData.text())
      // } else {
        await OpenFile()
      // }
    },
    // openFileMenu: Cherry.createMenuHook("打开", {
    //   onClick: async () => {
    //     await this.openFile()
    //   },
    // }),
    saveAsFile: async () => {
      // if (window.showSaveFilePicker) {
      //   const savefileHandle = await getSaveFileHandle()
      //   writeFile(savefileHandle, cherryInstance.getMarkdown())
      //   alert("保存成功")
      // } else {
        const mdBase64 = btoa(String.fromCharCode(...stringToBinaryArray(cherryInstance.getMarkdown())))

        const doc = new file.File()
        doc.Bytes = mdBase64

        await SaveFile(doc)
      // }
    },
    // saveAsFileMenu: Cherry.createMenuHook("另存为", {
    //   onClick: async () => {
    //     await this.saveAsFile()
    //   },
    // }),
    /**
     * 
     * @param {file.File|undefined|null} doc 
     */
    saveFile: async (doc) => {
      // if ((doc && doc.Path && doc.Bytes.length > 0) || !window.showSaveFilePicker) {
        const mdBase64 = btoa(String.fromCharCode(...stringToBinaryArray(cherryInstance.getMarkdown())))
        if (doc) {
          doc.Bytes = mdBase64
        } else {
          doc = new file.File()
          doc.Bytes = mdBase64
        }
        await SaveFile(doc)
      // } else {
      //   if (fileHandle) {
      //     await writeFile(fileHandle, cherryInstance.getMarkdown())
      //   } else {
      //     const savefileHandle = await getSaveFileHandle()
      //     await writeFile(savefileHandle, cherryInstance.getMarkdown())
      //   }
      //   alert("保存成功")
      // }
    },
    // saveFileMenu: Cherry.createMenuHook("保存", {
    //   onClick: async () => {
    //     await this.saveFile()
    //   },
    // }),
  }
}

/**
 * 
 * @param {string} path 
 * 
 */
export function isRelativePath(path) {
  return /^\.\/|^\.\.\//.test(path);
}