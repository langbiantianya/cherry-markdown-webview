<script>
  import { Router, createRouter } from "@roxi/routify";
  import routes from "../.routify/routes.default.js";
  export const router = createRouter({ routes });
  import { Circle2 } from "svelte-loading-spinners";
  import { globalState } from "./store.js";
  $: loading = $globalState.loading;
</script>

<Router {router} />
{#if loading}
  <div
    class="absolute top-0 left-0 right-0 z-20 bg-opacity-75 bg-white dark:bg-gray-800 w-full h-full flex items-center justify-center"
  >
    <Circle2 size="50" unit="vh" />
  </div>
{/if}
<style>
</style>
