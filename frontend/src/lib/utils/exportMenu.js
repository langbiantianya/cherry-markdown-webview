// eslint-disable-next-line no-unused-vars
import Cherry from "cherry-markdown"
import { SaveFile } from "../wailsjs/go/main/App"
import { file } from "../wailsjs/go/models"

/**
 * 
 * @param {string} htmlString 
 */
async function convertImgToBase64(htmlString) {
  const parser = new DOMParser();
  const doc = parser.parseFromString(htmlString, 'text/html');
  const imgElements = doc.querySelectorAll('img');

  for (const img of imgElements) {
    const imgUrl = img.getAttribute('src');

    if (imgUrl && (new URL(imgUrl)).hostname.includes("127.0.0.1")) {
      const response = await fetch(imgUrl, {
        method: 'GET',
        mode: 'cors',
      });
      const blob = await response.blob();
      const reader = new FileReader();
      reader.readAsDataURL(blob);
      await new Promise((resolve) => {
        reader.onloadend = () => {
          img.setAttribute('src', reader.result);
          resolve();
        };
      });
    }
  }

  return doc.documentElement.outerHTML;
}


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
      const mdHtml = await convertImgToBase64(cherryInstance.getHtml())
      const doc = new file.File()
      doc.StrData = mdHtml
      doc.DisplayName = "Html file"
      doc.Pattern = "*.html"
      await SaveFile(doc)
    }
  }
}
