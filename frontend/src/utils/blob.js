
/**
 * 
 * @param {number[]} array 
 * @returns  {Blob}
 */
export function arrayToBlob(array) {
	// 将number[]转换为Uint8Array
	const uint8Array = new Uint8Array(array);

	// 创建ArrayBuffer
	const arrayBuffer = uint8Array.buffer;

	// 创建Blob
	return new Blob([arrayBuffer], { type: "text/plain" });
}
/**
 * 
 * @param {string} base64Str 
 * @returns {string}
 */
export function base64ToString(base64Str) {
	const decoder = new TextDecoder('utf-8');
	return decoder.decode(new Uint8Array(atob(base64Str).split('').map(char => char.charCodeAt(0))));
}
/**
 * @param {Blob} blob 
 */
export async function blobToString(blob) {
	return new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.onload = () => resolve(reader.result);
		reader.onerror = reject;
		reader.readAsText(blob);
	});
}
/**
 * 
 * @param {string} str 
 * 
 */
export function stringToBinaryArray(str) {
	const encoder = new TextEncoder();
	return encoder.encode(str);
}