<script>
	import { file } from '$lib/wailsjs/go/models';
	import uploadPicture from '$lib/icon/upload-picture.svg';
	import { loadBackgroundImage } from '$lib/theme';

	/**
	 * @type {file.File}
	 */
	let backgroundImage = new file.File();
	/**
	 *
	 * @param {Event} event
	 */
	function handleFileSelect(event) {
		if (event.target instanceof HTMLInputElement && event.target.files) {
			const bgFile = event.target.files[0];

			const reader = new FileReader();
			reader.onload = function (e) {
				if (e.target) {
					const base64Image = e.target.result;
					backgroundImage.Mime = bgFile.type;
					backgroundImage.Name = bgFile.name;
					backgroundImage.Bytes = base64Image?.toString().split(',')[1];
					// TODO 保存到配置目录中
					loadBackgroundImage(base64Image);
				}
			};
			reader.readAsDataURL(bgFile);
		}
	}
	function openFileDialog() {
		const fileInput = document.createElement('input');
		fileInput.type = 'file';
		fileInput.accept = 'image/*';
		fileInput.multiple = false;
		fileInput.style.display = 'none';
		fileInput.onchange = handleFileSelect;
		document.body.appendChild(fileInput);
		fileInput.click();
		document.body.removeChild(fileInput);
	}
</script>

<div class="w-2/3">
	<fluent-tooltip id="tooltip" anchor="diy-background-name" delay="100" auto-update-mode="anchor">
		点击上传背景图片
	</fluent-tooltip>

	<fluent-image class="" id="diy-background-name" fit="cover" shape="rounded">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
		<img
			class=""
			on:click={openFileDialog}
			src={backgroundImage.Bytes
				? `data:data:${backgroundImage.Mime};base64,${backgroundImage.Bytes}`
				: uploadPicture}
			alt="点击选择图片"
		/></fluent-image
	>
</div>

<style>
</style>
