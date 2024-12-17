// eslint-disable-next-line no-unused-vars
import Cherry from "cherry-markdown"
import { file } from "../wailsjs/go/models"
import { SaveFile, OpenFile } from "../wailsjs/go/main/App"
import { stringToBinaryArray } from "./blob"

export const FileMenu = function () {
  /**
   * @type {Cherry} cherryInstance 
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

    openFile: async () => {
      
      await OpenFile()

    },
  
    saveAsFile: async () => {

      const mdBase64 = btoa(String.fromCharCode(...stringToBinaryArray(cherryInstance.getMarkdown())))

      const doc = new file.File()
      doc.Bytes = mdBase64
      await SaveFile(doc)

    },
    /**
     * 
     * @param {file.File|undefined|null} doc 
     */
    saveFile: async (doc) => {

      const mdBase64 = btoa(String.fromCharCode(...stringToBinaryArray(cherryInstance.getMarkdown())))
      if (doc) {
        doc.Bytes = mdBase64
      } else {
        doc = new file.File()
        doc.Bytes = mdBase64
      }
      return await SaveFile(doc)
     
    },
  }
}

/**
 * 
 * @param {string} path 
 * 
 */
export function isRelativePath(path) {
  return /^\.\/|^\.\.\//.test(path)
}

/**
 * 
 * @param {string} path 
 * 
 */
export function isRootDirectory(path) {
  const unixRegex = /^\//
  const windowsRegex = /^([a-zA-Z]:\\)/
  return unixRegex.test(path) || windowsRegex.test(path)
}