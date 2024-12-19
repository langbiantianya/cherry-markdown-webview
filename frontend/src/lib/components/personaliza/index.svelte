<script>
	import { globalState } from '$lib/store';
	import { goto } from '$app/navigation';
	$globalState.loading = false;

	function goBack() {
		// goto('/', { replaceState: true });
	}
</script>

<div class="settings_page h-full">
	<h1>个性化</h1>
	<button on:click={goBack}>返回</button>
</div>

<style>
	.settings_page {
		background-color: white;
	}
</style>
