<script>
	import { globalState } from '$lib/store';
	import '@fluentui/web-components/button.js';
	import '@fluentui/web-components/tab.js';
	import '@fluentui/web-components/tabs.js';
	import '@fluentui/web-components/tab-panel.js';
	import '@fluentui/web-components/accordion.js';
	import '@fluentui/web-components/accordion-item.js';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';
	import '@fluentui/web-components/drawer.js';
	// import '@fluentui/web-components/drawer-body.js';
	import help from '$lib/icon/help.svg';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime';
	/**
	 * @type {{heid:function():void}}
	 */
	let { heid } = $props();
	function goBack() {
		heid();
	}
	function saveConfig() {}
	function tencentCosHelp() {
		BrowserOpenURL('https://cloud.tencent.com/document/product/436/56390');
	}
	function aliOssHelp() {
		BrowserOpenURL('https://cloud.tencent.com/document/product/436/56390');
	}
	$globalState.loading = false;
</script>

<div class="m-0 p-0">
	<div part="header" class="bg-defaultToolBar h-12 w-full px-4">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div class="flex h-12 w-full flex-nowrap items-center justify-between text-white">
			<h1 class="text-2xl">首选项</h1>
			<!-- svelte-ignore event_directive_deprecated -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<fluent-button
				class="hover:bg-defaultButtonHover bg-defaultToolBar h-9 rounded-none leading-9 text-white"
				on:click={goBack}>返回</fluent-button
			>
		</div>
	</div>
	<div class="drawer-content overflow-y-scroll px-1">
		<fluent-tabs orientation="vertical" size="large">
			<fluent-tab id="oss">图床</fluent-tab>
			<fluent-tab-panel id="ossPanel">
				<fluent-accordion>
					<fluent-accordion-item class="w-full" size="large">
						<span slot="heading">默认配置</span>
						<div class="panel">
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="DefaultOssPath">路径</fluent-label>
							<fluent-text-input
								id="DefaultOssPath"
								appearance="outline"
								placeholder="/cherrymarkdown"
							></fluent-text-input>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="DefaultOss">配置</fluent-label>
							<select
								id="DefaultOss"
								class="border-coolGray-300 h-8 w-full rounded border px-1 leading-8"
							>
								<option value="base64">base64</option>
								<option value="cos">腾讯COS</option>
								<option value="oss">阿里OSS</option>
							</select>
						</div>
					</fluent-accordion-item>

					<fluent-accordion-item class="" size="large">
						<span slot="heading">腾讯COS </span>

						<div class="panel">
							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={tencentCosHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretID"
								>SecretID <img
									slot="start"
									class="rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
									src={help}
									alt="help"
								/></fluent-label
							>
							<fluent-text-input id="SecretID" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={tencentCosHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretKey">SecretKey</fluent-label>
							<fluent-text-input id="SecretKey" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={tencentCosHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="BucketURL">BucketURL</fluent-label>
							<fluent-text-input
								id="BucketURL"
								appearance="outline"
								placeholder="https://examplebucket-1250000000.cos.<Region>.myqcloud.com"
							></fluent-text-input>
						</div>
					</fluent-accordion-item>
					<fluent-accordion-item class="" size="large">
						<span slot="heading">阿里OSS</span>
						<div class="panel">
							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={aliOssHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeyID">AccessKeyID</fluent-label>
							<fluent-text-input id="AccessKeyID" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={aliOssHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeySecret">AccessKeySecret</fluent-label>
							<fluent-text-input id="AccessKeySecret" appearance="outline"></fluent-text-input>

							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={aliOssHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="Region">Region</fluent-label>
							<fluent-text-input id="Region" appearance="outline" placeholder="cn-hangzhou"
							></fluent-text-input>

							<!-- svelte-ignore a11y_click_events_have_key_events -->
							<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
							<img
								on:click={aliOssHelp}
								class="inline rounded p-0.5 text-gray-500 hover:bg-slate-200 hover:text-black"
								src={help}
								alt="help"
							/>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="BucketName">BucketName</fluent-label>
							<fluent-text-input
								id="BucketName"
								appearance="outline"
								placeholder="ExampleBucketName"
							></fluent-text-input>
						</div>
					</fluent-accordion-item>
				</fluent-accordion>
			</fluent-tab-panel>
		</fluent-tabs>
	</div>
</div>

<style>
	/* 重写 select 选项的样式 */
	select option {
		border-radius: 4px;
	}

	.drawer-content {
		height: calc(100vh - 3.2rem);
	}
</style>
