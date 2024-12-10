import Cherry from "cherry-markdown";
import { getOpenFileHandle, getSaveFileHandle, writeFile } from "./io";

export const ExportMenu = function () {
  //   let fileHandle;
  let cherryInstance;
  return {
    /**
     *
     * @param {Cherry} cherry
     */
    setCherry: (cherry) => {
      cherryInstance = cherry;
    },
    exportMenu: Cherry.createMenuHook("导出", {}),
    exportPdf: () => {
      cherryInstance.export();
    },
    exportPdfMenu: Cherry.createMenuHook("导出PDF", {
      onClick: () => {
        this.exportPdf()
      },
    }),
    exportHtml: async () => {
      const mdHtml = cherryInstance.getHtml();
      const fileHandle = await getSaveFileHandle({
        types: [
          {
            description: "Html file",
            // @ts-ignore
            accept: {
              "text/html": [".html"],
            },
          },
        ],
      });
      await writeFile(fileHandle, mdHtml);
    },
    exportHtmlMenu: Cherry.createMenuHook("导出Html", {
      onClick: async () => {
        await this.exportHtml()
      },
    }),
  };
};
