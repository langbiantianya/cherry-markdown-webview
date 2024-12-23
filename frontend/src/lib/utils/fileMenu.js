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
      const doc = new file.File()
      doc.StrData = cherryInstance.getMarkdown()
      await SaveFile(doc)

    },
    /**
     * 
     * @param {file.File|undefined|null} doc 
     */
    saveFile: async (doc) => {
      if (doc) {
        doc.StrData = cherryInstance.getMarkdown()
      } else {
        doc = new file.File()
        doc.StrData = cherryInstance.getMarkdown()
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