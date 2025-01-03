<script>
	import { goto } from '$app/navigation';
	import '../app.css';
	import { Circle2 } from 'svelte-loading-spinners';
	import { globalState } from '$lib/store';
	import { EventsOn, EventsOff } from '$lib/wailsjs/runtime';
	import '@fluentui/web-components/drawer.js';
	import '@fluentui/web-components/button.js';
	import '@fluentui/web-components/dialog.js';
	import Options from '$lib/components/options/index.svelte';
	import Personaliza from '$lib/components/personaliza/index.svelte';
	import About from '$lib/components/about/index.svelte';
	import { GetActivatedTheme } from '$lib/wailsjs/go/main/App';

	import { onMount } from 'svelte';
	import { changeMainThemeEvent, loadBackgroundImage } from '$lib/theme';

	/** @type {{children: import('svelte').Snippet}} */
	let { children } = $props();
	/** @type {import('@fluentui/web-components').Drawer} */
	let optionsDrawer;

	/** @type {import('@fluentui/web-components').Drawer} */
	let personalizaDrawer;

	/** @type {import('@fluentui/web-components').Dialog} */
	let aboutDialog;
	function heidOptionsDrawer() {
		optionsDrawer.hide();
	}

	function heidPersonalizaDrawer() {
		personalizaDrawer.hide();
	}
	try {
		EventsOff('optionsEvent', 'personalizaEvent', 'aboutEvent');
		EventsOn('aboutEvent', (event) => {
			aboutDialog.show();
		});
		EventsOn('optionsEvent', (event) => {
			optionsDrawer.show();
		});
		EventsOn('personalizaEvent', (event) => {
			personalizaDrawer.show();
		});
	} catch (error) {
		console.error(error);
	}

	onMount(async () => {
		loadBackgroundImage();
		changeMainThemeEvent(await GetActivatedTheme());
	});
</script>

<div class="app h-full">
	<div class="background-image"></div>
	<main class="h-full">
		{@render children()}
	</main>
	<footer></footer>
	<fluent-drawer class="m-0 p-0" type="model" position="end" size="full" bind:this={optionsDrawer}>
		<Options heid={heidOptionsDrawer} />
	</fluent-drawer>

	<fluent-drawer
		class="m-0 p-0"
		type="model"
		position="end"
		size="full"
		bind:this={personalizaDrawer}
	>
		<Personaliza heid={heidPersonalizaDrawer} />
	</fluent-drawer>
	<fluent-dialog bind:this={aboutDialog}>
		<About />
	</fluent-dialog>

	{#if $globalState.loading}
		<div class="background-image"></div>
		<div
			class="loading bg-base absolute left-0 right-0 top-0 z-50 flex h-lvh w-full items-center justify-center"
		>
			<Circle2 size="50" unit="lvh" />
		</div>
	{/if}
</div>

<style>
	main {
		opacity: var(--edit-opacity);
	}
	.loading {
		opacity: var(--edit-opacity);
	}
	.background-image {
		opacity: 1;
		position: absolute;
		top: 0;
		left: 0;
		width: 100lvw;
		height: 100lvh;
		background-image: var(--background-image); /* 设置背景图片路径 */
		background-size: cover; /* 调整背景图片大小以覆盖整个容器 */
		background-position: center; /* 居中背景图片 */
	}
</style>
