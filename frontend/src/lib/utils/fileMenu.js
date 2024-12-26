// eslint-disable-next-line no-unused-vars
import Cherry from "cherry-markdown"
import { file } from "../wailsjs/go/models"
import { SaveFile, OpenFile } from "../wailsjs/go/main/App"

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
        const _doc = new file.File()
        _doc.DisplayName = doc.DisplayName
        _doc.Path = doc.Path
        _doc.Mime = doc.Mime
        _doc.Name = doc.Name
        _doc.StrData = cherryInstance.getMarkdown()
        doc = _doc
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