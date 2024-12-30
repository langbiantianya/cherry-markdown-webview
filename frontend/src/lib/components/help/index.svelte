<script>
	import { onMount } from 'svelte';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime';

	/**
	 * @type {HTMLSpanElement}
	 */
	let svgSpan;
	/**
	 * @type {{url:string}}
	 */
	let { url } = $props();

	function helpOpenUrl() {
		BrowserOpenURL(url);
	}

	onMount(() => {

	});
</script>

<!-- svelte-ignore event_directive_deprecated -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_mouse_events_have_key_events -->
<span
	on:click={helpOpenUrl}
	class="hover:bg-buttonHover rounded p-0.5 pt-0 text-base"
	bind:this={svgSpan}
>
	<svg
		class="inline h-5 pb-0.5"
		width="16"
		height="16"
		viewBox="0 0 48 48"
		fill="none"
		xmlns="http://www.w3.org/2000/svg"
	>
		<path
			d="M24 44C29.5228 44 34.5228 41.7614 38.1421 38.1421C41.7614 34.5228 44 29.5228 44 24C44 18.4772 41.7614 13.4772 38.1421 9.85786C34.5228 6.23858 29.5228 4 24 4C18.4772 4 13.4772 6.23858 9.85786 9.85786C6.23858 13.4772 4 18.4772 4 24C4 29.5228 6.23858 34.5228 9.85786 38.1421C13.4772 41.7614 18.4772 44 24 44Z"
			fill="none"
			stroke="#333"
			stroke-width="4"
			stroke-linejoin="round"
		/>
		<path
			d="M24 28.6248V24.6248C27.3137 24.6248 30 21.9385 30 18.6248C30 15.3111 27.3137 12.6248 24 12.6248C20.6863 12.6248 18 15.3111 18 18.6248"
			stroke="#333"
			stroke-width="4"
			stroke-linecap="round"
			stroke-linejoin="round"
		/>
		<path
			class="third-path"
			fill-rule="evenodd"
			clip-rule="evenodd"
			d="M24 37.6248C25.3807 37.6248 26.5 36.5055 26.5 35.1248C26.5 33.7441 25.3807 32.6248 24 32.6248C22.6193 32.6248 21.5 33.7441 21.5 35.1248C21.5 36.5055 22.6193 37.6248 24 37.6248Z"
			fill="#333"
		/>
	</svg>
</span>

<style>
	svg path {
		stroke: var(--color-text-base);
	}
	.third-path {
		fill: var(--color-text-base);
	}
	span:hover svg .third-path{
		fill: var(--color-text-button-hover);

	}
	span:hover svg path {
		stroke: var(--color-text-button-hover);
	}
</style>
