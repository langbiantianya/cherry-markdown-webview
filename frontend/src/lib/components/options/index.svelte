<script>
	import '@fluentui/web-components/button.js';
	import '@fluentui/web-components/tab.js';
	import '@fluentui/web-components/tabs.js';
	import '@fluentui/web-components/tab-panel.js';
	import '@fluentui/web-components/accordion.js';
	import '@fluentui/web-components/accordion-item.js';
	import '@fluentui/web-components/text-input.js';
	import '@fluentui/web-components/label.js';
	import '@fluentui/web-components/drawer.js';
	import Help from '$lib/components/help/index.svelte';
	import { onMount } from 'svelte';
	import { GetPicBed, UpsertPicBed } from '$lib/wailsjs/go/main/App';
	import { config } from '$lib/wailsjs/go/models';
	import { TextInput } from '@fluentui/web-components';
	/**
	 * @type {{heid:function():void}}
	 */
	let { heid } = $props();
	function goBack() {
		heid();
	}
	/**
	 * @type {config.PicBed}
	 */
	let picBedConf = $state(new config.PicBed());
	/**
	 * @type {config.PicBed}
	 */
	let picBedConfBackup;

	const tencentCosHelp = 'https://cloud.tencent.com/document/product/436/56390';
	const aliOssHelp =
		'https://help.aliyun.com/zh/oss/developer-reference/quick-start-for-oss-go-sdk-v2?spm=a2c4g.11186623.0.0.363e78b3LyiJ7Z';

	async function updatePicConf() {
		await UpsertPicBed(config.PicBed.createFrom($state.snapshot(picBedConf)));
	}
	/**
	 *
	 * @param {Event} event
	 */
	function activatedChange(event) {
		picBedConf.activated =
			(event.target && event.target instanceof HTMLSelectElement && event.target.value) ||
			picBedConf.activated;
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function basePathInput(event) {
		picBedConf.basePath = (
			(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
			picBedConf.basePath
		).trim();
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function cosSecretIDInput(event) {
		picBedConf.cos = {
			secretKey: '',
			bucketURL: '',
			...picBedConf.cos,
			secretID: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function cosSecretKeyInput(event) {
		picBedConf.cos = {
			secretID: '',
			bucketURL: '',
			...picBedConf.cos,
			secretKey: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function cosBucketURLInput(event) {
		picBedConf.cos = {
			secretID: '',
			secretKey: '',
			...picBedConf.cos,
			bucketURL: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}

	/**
	 *
	 * @param {Event} event
	 */
	function ossAccessKeyIDInput(event) {
		picBedConf.oss = {
			accessKeySecret: '',
			region: '',
			bucketName: '',
			...picBedConf.oss,
			accessKeyID: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function ossAccessKeySecretInput(event) {
		picBedConf.oss = {
			accessKeyID: '',
			region: '',
			bucketName: '',
			...picBedConf.oss,
			accessKeySecret: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function ossRegionInput(event) {
		picBedConf.oss = {
			accessKeySecret: '',
			accessKeyID: '',
			bucketName: '',
			...picBedConf.oss,
			region: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	/**
	 *
	 * @param {Event} event
	 */
	function ossBucketNameInput(event) {
		picBedConf.oss = {
			region: '',
			accessKeySecret: '',
			accessKeyID: '',
			...picBedConf.oss,
			bucketName: (
				(event.target && event.target instanceof TextInput && event.target['currentValue']) ||
				''
			).trim()
		};
		updatePicConf();
	}
	onMount(async () => {
		picBedConf = await GetPicBed();
		picBedConfBackup = JSON.parse(JSON.stringify(picBedConf));
	});
</script>

<div class="bg-base m-0 p-0">
	<div part="header" class="bg-toolBar h-12 w-full px-4 shadow-sm">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div class="text-toolBar flex h-12 w-full flex-nowrap items-center justify-between">
			<h1 class="text-2xl">首选项</h1>
			<!-- svelte-ignore event_directive_deprecated -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<fluent-button
				class="hover:bg-buttonHover hover:text-buttonHover bg-button text-button h-9 rounded-none leading-9"
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
						<span slot="heading" class="accordion-item-head">默认配置</span>
						<div class="panel">
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="basePath" hidden={picBedConf.activated === 'Base64'}
								>路径</fluent-label
							>
							<fluent-text-input
								hidden={picBedConf.activated === 'Base64'}
								on:input={basePathInput}
								initialValue={picBedConf.basePath}
								id="basePath"
								appearance="outline"
								placeholder="/cherrymarkdown"
							></fluent-text-input>
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="activated">启用配置</fluent-label>
							<select
								on:change={activatedChange}
								value={picBedConf.activated}
								name="activated"
								id="activated"
								class="bg-input border-0.1 border-select text-input h-8 w-full appearance-none rounded px-1 leading-8 outline-0"
							>
								<option value="Base64">base64</option>
								<option value="COS">腾讯COS</option>
								<option value="OSS">阿里OSS</option>
							</select>
						</div>
					</fluent-accordion-item>

					<fluent-accordion-item class="" size="large">
						<span slot="heading" class="accordion-item-head">腾讯COS </span>

						<div class="panel">
							<Help url={tencentCosHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretID">SecretID</fluent-label>
							<fluent-text-input
								on:input={cosSecretIDInput}
								initialValue={picBedConf.cos?.secretID}
								id="SecretID"
								appearance="outline"
							></fluent-text-input>
							<Help url={tencentCosHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="SecretKey">SecretKey</fluent-label>
							<fluent-text-input
								on:input={cosSecretKeyInput}
								initialValue={picBedConf.cos?.secretKey}
								id="SecretKey"
								appearance="outline"
							></fluent-text-input>
							<Help url={tencentCosHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="BucketURL">BucketURL</fluent-label>
							<fluent-text-input
								on:input={cosBucketURLInput}
								initialValue={picBedConf.cos?.bucketURL}
								id="BucketURL"
								appearance="outline"
								placeholder="https://examplebucket-1250000000.cos.<Region>.myqcloud.com"
							></fluent-text-input>
						</div>
					</fluent-accordion-item>
					<fluent-accordion-item class="" size="large">
						<span slot="heading" class="accordion-item-head">阿里OSS</span>
						<div class="panel">
							<Help url={aliOssHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeyID">AccessKeyID</fluent-label>
							<fluent-text-input
								on:input={ossAccessKeyIDInput}
								initialValue={picBedConf.oss?.accessKeyID}
								id="AccessKeyID"
								appearance="outline"
							></fluent-text-input>
							<Help url={aliOssHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="AccessKeySecret">AccessKeySecret</fluent-label>
							<fluent-text-input
								on:input={ossAccessKeySecretInput}
								initialValue={picBedConf.oss?.accessKeySecret}
								id="AccessKeySecret"
								appearance="outline"
							></fluent-text-input>

							<Help url={aliOssHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="Region">Region</fluent-label>
							<fluent-text-input
								on:input={ossRegionInput}
								initialValue={picBedConf.oss?.region}
								id="Region"
								appearance="outline"
								placeholder="cn-hangzhou"
							></fluent-text-input>

							<Help url={aliOssHelp} />
							<!-- svelte-ignore attribute_invalid_property_name -->
							<fluent-label htmlFor="BucketName">BucketName</fluent-label>
							<fluent-text-input
								on:input={ossBucketNameInput}
								initialValue={picBedConf.oss?.bucketName}
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
	/* .control {
		color: var(--color-text-base);
	} */
	.drawer-content {
		height: calc(100vh - 3.2rem);
	}
</style>
