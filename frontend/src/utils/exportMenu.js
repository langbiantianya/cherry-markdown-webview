import Cherry from "cherry-markdown"
import { SaveFile } from "../../wailsjs/go/main/App"
import { file } from "../../wailsjs/go/models"
import { stringToBinaryArray } from "./blob"

export const ExportMenu = function () {
  //   let fileHandle
  /**
   * @type {Cherry}
   */
  let cherryInstance
  return {
    /**
     *
     * @param {Cherry} cherry
     */
    setCherry: (cherry) => {
      cherryInstance = cherry
    },
    // exportMenu: Cherry.createMenuHook("导出", {}),
    exportPdf: () => {
      cherryInstance.export()
    },
    // exportPdfMenu: Cherry.createMenuHook("导出PDF", {
    //   onClick: () => {
    //     this.exportPdf()
    //   },
    // }),
    exportHtml: async () => {
      const mdHtml = cherryInstance.getHtml()
      // if (window.showSaveFilePicker) {
      //   const fileHandle = await getSaveFileHandle({
      //     types: [
      //       {
      //         description: "Html file",
      //         // @ts-ignore
      //         accept: {
      //           "text/html": [".html"],
      //         },
      //       },
      //     ],
      //   })
      //   await writeFile(fileHandle, mdHtml)
      // } else {
      debugger
      const htmlBase64 = btoa(String.fromCharCode(...stringToBinaryArray(mdHtml)))

      const doc = new file.File()
      doc.Bytes = htmlBase64
      doc.DisplayName = "Html file"
      doc.Pattern = "*.html"
      await SaveFile(doc)
      // }

    },
    // exportHtmlMenu: Cherry.createMenuHook("导出Html", {
    //   onClick: async () => {
    //     await this.exportHtml()
    //   },
    // }),
  }
}
