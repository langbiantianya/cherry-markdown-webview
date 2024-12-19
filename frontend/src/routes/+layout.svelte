<script>
	import { goto } from '$app/navigation';
	import '../app.css';
	import { Circle2 } from 'svelte-loading-spinners';
	import { globalState } from '$lib/store';
	import { EventsOn, EventsOff } from '$lib/wailsjs/runtime';
	import { setTheme } from '@fluentui/web-components';
	import { webLightTheme } from '@fluentui/tokens';
	import '@fluentui/web-components/drawer.js';
	// import '@fluentui/web-components/drawer-body.js';
	import '@fluentui/web-components/button.js';
	import Options from '$lib/components/options/index.svelte';
	setTheme(webLightTheme);
	/** @type {{children: import('svelte').Snippet}} */
	let { children } = $props();
	/** @type {import('@fluentui/web-components').Drawer} */
	let drawer;
	function heidDrawer() {
		drawer.hide();
	}
	try {
		EventsOff('optionsEvent', 'personalizaEvent');
		EventsOn('optionsEvent', (event) => {
			drawer.show();
			// goto('/settings/options');
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
		<fluent-drawer class="m-0 p-0" type="model" position="end" size="full" bind:this={drawer}>
			<Options heid={heidDrawer} />
		</fluent-drawer>
	</main>

	<footer></footer>

	{#if $globalState.loading}
		<div
			class="absolute left-0 right-0 top-0 z-50 flex h-lvh w-full items-center justify-center bg-white bg-opacity-75 dark:bg-gray-800"
		>
			<Circle2 size="50" unit="vh" />
		</div>
	{/if}
</div>

<style>
</style>
