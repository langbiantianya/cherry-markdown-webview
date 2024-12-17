/**
 *
 * @param {FileHandle} fileHandle
 * @param {string|Blob|File} contents
 */
export async function writeFile(fileHandle, contents) {
  const writable = await fileHandle.createWritable()
  await writable.write(contents)
  await writable.close()
}

export async function getOpenFileHandle(
  option = {
    types: [
      {
        description: "Markdown file",
        accept: {
          "text/plain": [".md", ".mdx"],
        },
      },
    ],
    multiple: false,
  }
) {
  const [fileHandle] = await window.showOpenFilePicker(option)
  return fileHandle
}

export async function getSaveFileHandle(
  option = {
    types: [
      {
        description: "Markdown file",
        accept: {
          "text/plain": [".md", ".mdx"],
        },
      },
    ],
  }
) {
  return await window.showSaveFilePicker(option)
}
