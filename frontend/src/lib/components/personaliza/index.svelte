<script>
	import '@fluentui/web-components/button.js';
	import '@fluentui/web-components/tab.js';
	import '@fluentui/web-components/tabs.js';
	import '@fluentui/web-components/tab-panel.js';
	import { onMount } from 'svelte';
	import '@fluentui/web-components/tab.js';
	import '@fluentui/web-components/tabs.js';
	import '@fluentui/web-components/tab-panel.js';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';
	import DiyTheme from './diyTheme.svelte';
	/**
	 * @type {{heid:function():void}}
	 */
	let { heid } = $props();

	function goBack() {
		// TODO 重置主题为激活的主题
		heid();
	}

	onMount(async () => {});
</script>

<div class="bg-base m-0 h-full overflow-hidden p-0">
	<div part="header" class="bg-toolBar h-12 w-full px-4 shadow-sm">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div class="text-toolBar flex h-12 w-full flex-nowrap items-center justify-between">
			<h1 class="text-2xl">个性化</h1>
			<!-- svelte-ignore event_directive_deprecated -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<fluent-button
				class="hover:bg-buttonHover hover:text-buttonHover bg-button text-button h-9 rounded-none leading-9"
				on:click={goBack}>返回</fluent-button
			>
		</div>
	</div>
	<fluent-tabs orientation="vertical" size="large">
		<fluent-tab id="background-image">背景图片</fluent-tab>
		<fluent-tab-panel id="background-image-panel"> </fluent-tab-panel>

		<fluent-tab id="diy-theme">自定义主题</fluent-tab>
		<fluent-tab-panel id="diy-theme-panel">
			<DiyTheme />
		</fluent-tab-panel>
	</fluent-tabs>
</div>

<style>
	/* 重写 select 选项的样式 */
	.accordion-item-head {
		color: var(--color-text-base);
	}
	fluent-tab {
		color: var(--color-text-base);
	}
	fluent-label {
		color: var(--color-text-base);
	}
	fluent-tab {
		--colorCompoundBrandStroke: var(--color-text-base);
	}
	fluent-text-input {
		--colorNeutralForeground1: var(--color-text-base);
		--colorCompoundBrandStroke: var(--color-text-base);
		--colorCompoundBrandStrokePressed: var(--color-text-base);
		--colorNeutralBackground1: var(--color-bg-input);
	}

	.drawer-contentw-full {
		height: calc(100lvh +200px);
	}
	.edit {
		padding-bottom: 1rem;
		/* min-width: 37rem;
		max-width: 57rem; */
	}
	@media (max-width: 768px) {
		.theme-input {
			min-width: 100lvw;
		}
	}
	@media (min-width: 1080px) {
		.theme-input {
			justify-content: flex-start; /* 水平居中 */
			align-items: flex-start;
		}
		.drawer-contentw-full {
			/* align-items: center;  */
			justify-content: center; /* 水平居中 */
		}
		.edit {
			width: 47rem;
			height: 27rem;
		}
		.theme-input {
			max-width: 25rem;
			height: 27rem;
		}
	}
</style>
