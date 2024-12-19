<script>
	import { globalState } from '../../../lib/store';
	import '@fluentui/web-components/button.js';
	import '@fluentui/web-components/tab.js';
	import '@fluentui/web-components/tabs.js';
	import '@fluentui/web-components/tab-panel.js';
	import '@fluentui/web-components/accordion.js';
	import '@fluentui/web-components/accordion-item.js';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';

	$globalState.loading = false;
	function goBack() {
		window.history.back();
	}
	function saveConfig() {}
</script>

<div class="">
	<header class="bg-defaultToolBar h-12 w-full px-4">
		<div class="flex h-12 w-full flex-nowrap items-center justify-between text-white">
			<h1 class="text-2xl">首选项</h1>
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<fluent-button
				class="hover:bg-defaultButtonHover bg-defaultToolBar h-9 rounded-none leading-9 text-white"
				on:click={goBack}>返回</fluent-button
			>
		</div>
	</header>
	<div class="px-1">
		<fluent-tabs orientation="vertical" size="large">
			<fluent-tab id="oss">图床</fluent-tab>
			<fluent-tab-panel id="ossPanel">
				<fluent-accordion>
					<fluent-accordion-item class="w-full">
						<span slot="heading">默认配置</span>
						<div class="panel w-full">
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
					<fluent-accordion-item class="w-full">
						<span slot="heading">腾讯COS</span>
						<div class="panel w-full">
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretID">SecretID</fluent-label>
							<fluent-text-input id="SecretID" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretKey">SecretKey</fluent-label>
							<fluent-text-input id="SecretKey" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="BucketURL">BucketURL</fluent-label>
							<fluent-text-input
								id="BucketURL"
								appearance="outline"
								placeholder="https://examplebucket-1250000000.cos.<Region>.myqcloud.com"
							></fluent-text-input>
						</div>
					</fluent-accordion-item>
					<fluent-accordion-item class="w-full">
						<span slot="heading">阿里OSS</span>
						<div class="panel w-full">
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeyID">AccessKeyID</fluent-label>
							<fluent-text-input id="AccessKeyID" appearance="outline"></fluent-text-input>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeySecret">AccessKeySecret</fluent-label>
							<fluent-text-input id="AccessKeySecret" appearance="outline"></fluent-text-input>

							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="Region">Region</fluent-label>
							<fluent-text-input id="Region" appearance="outline" placeholder="cn-hangzhou"
							></fluent-text-input>

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
</style>
