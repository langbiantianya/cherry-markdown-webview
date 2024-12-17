<script>
	import { goto } from '$app/navigation';
	import '../app.css';
	import { Circle2 } from 'svelte-loading-spinners';
	import { globalState } from '../lib/store';
	import { EventsOn } from '../lib/wailsjs/runtime';

	/** @type {{children: import('svelte').Snippet}} */
	let { children } = $props();
	EventsOn('optionsEvent', (event) => {
		goto('/settings/options', { replaceState: true });
	});
	EventsOn('personalizaEvent', (event) => {
		goto('/settings/personaliza', { replaceState: true });
	});
</script>

<div class="app h-full">
	<main class="h-full">
		{@render children()}
	</main>

	<footer></footer>
	{#if $globalState.loading}
		<div
			class="absolute left-0 right-0 top-0 z-20 flex h-full w-full items-center justify-center bg-white bg-opacity-75 dark:bg-gray-800"
		>
			<Circle2 size="50" unit="vh" />
		</div>
	{/if}
</div>

<style>
</style>
