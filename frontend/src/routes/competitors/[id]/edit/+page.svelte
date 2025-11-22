<script lang="ts">
	import { page } from "$app/state";
	import SiteForm from "$lib/components/SiteForm.svelte";
	import { client } from "$lib/pocketbase";
	import type { CompetitorsRecord } from "$lib/pocketbase/generated-types";
	const competitor = () => client.collection('competitors').getOne<CompetitorsRecord>(page.params.id as string);
</script>
<main>
  {#await competitor()}
    Loading ...
  {:then site} 
    <SiteForm {site} />  
  {/await}
</main>
