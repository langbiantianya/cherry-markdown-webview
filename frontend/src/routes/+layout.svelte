<script>
	import { goto } from '$app/navigation';
	import '../app.css';
	import { Circle2 } from 'svelte-loading-spinners';
	import { globalState } from '../lib/store';
	import { EventsOn } from '../lib/wailsjs/runtime';
	import { setTheme } from '@fluentui/web-components';
	import { webLightTheme } from '@fluentui/tokens';

	setTheme(webLightTheme);
	/** @type {{children: import('svelte').Snippet}} */
	let { children } = $props();
	try {
		EventsOn('optionsEvent', (event) => {
			goto('/settings/options');
		});
		EventsOn('personalizaEvent', (event) => {
			goto('/settings/personaliza');
		});
	} catch (error) {
		console.error(error);
	}
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
