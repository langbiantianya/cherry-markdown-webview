<script>
	import { BrowserOpenURL } from '$lib/wailsjs/runtime';
	function openSourceCode() {
		BrowserOpenURL('https://github.com/langbiantianya/cherry-markdown-webview');
	}
	function openGithubPage() {
		BrowserOpenURL('https://github.com/langbiantianya');
	}

	function openDepends() {
		BrowserOpenURL(
			'https://github.com/langbiantianya/cherry-markdown-webview?tab=readme-ov-file#depends'
		);
	}
</script>

<div class="bg-base flex h-40 flex-wrap content-between justify-center">
	<h2 class="w-full text-center text-base text-xl font-bold">Cherry Markdown On WebView</h2>

	<p class="w-full text-center">
		<span class="text-sm">
			<a on:click={openDepends}>Open Source</a>
			|
			<a class=" " on:click={openSourceCode}>源码地址</a>
		</span>
		<br />
		<span class="text-xs"
			>Copyright © 2024 <a on:click={openGithubPage}>langbiantianya</a>. All rights reserved.</span
		>
	</p>
</div>

<style></style>
