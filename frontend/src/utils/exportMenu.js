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
    exportPdfMenu: Cherry.createMenuHook("导出PDF", {
      onClick: () => {
        cherryInstance.export();
      },
    }),
    exportHtmlMenu: Cherry.createMenuHook("导出Html", {
      onClick: async () => {
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
    }),
  };
};
