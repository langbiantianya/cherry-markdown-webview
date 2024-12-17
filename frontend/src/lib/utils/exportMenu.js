// eslint-disable-next-line no-unused-vars
import Cherry from "cherry-markdown"
import { SaveFile } from "../wailsjs/go/main/App"
import { file } from "../wailsjs/go/models"
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
    
    exportPdf: () => {
      cherryInstance.export()
    },

    exportHtml: async () => {
      const mdHtml = cherryInstance.getHtml()
      const htmlBase64 = btoa(String.fromCharCode(...stringToBinaryArray(mdHtml)))

      const doc = new file.File()
      doc.Bytes = htmlBase64
      doc.DisplayName = "Html file"
      doc.Pattern = "*.html"
      await SaveFile(doc)
      // }

    }
  }
}
