<script lang="ts">
	import { Heading, Drawer } from 'flowbite-svelte';
	import SiteForm from './SiteForm.svelte';

	let { open = $bindable(false), callback = null, data = {} } = $props();
	let title = $derived(data.id ? 'Edit body' : 'Add new entity');
</script>

<Drawer width="half" placement="right" bind:open>
	<Heading tag="h5" class="mb-6 text-sm font-semibold uppercase">
		{title || (data.id ? 'Edit body' : 'Add new entity')}
	</Heading>

    <SiteForm {data} afterUpdate={callback} bind:open />
</Drawer>
