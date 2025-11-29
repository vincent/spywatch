<script lang="ts">
	import { page } from "$app/state";
	import SiteForm from "$lib/components/SiteForm.svelte";
	import { client } from "$lib/pocketbase";
	import type { BodiesRecord } from "$lib/pocketbase/generated-types";
	const body = () => client.collection('bodies').getOne<BodiesRecord>(page.params.id as string);
</script>
<main>
  {#await body()}
    Loading ...
  {:then site} 
    <SiteForm {site} />  
  {/await}
</main>
