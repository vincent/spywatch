<script lang="ts">
	import { Heading, Drawer } from 'flowbite-svelte';
	import SiteForm from './SiteForm.svelte';

	let { open = $bindable(false), afterUpdate = null, data = {} } = $props();
	let title = $derived(data.id ? 'Edit competitor' : 'Add new competitor');
</script>

<Drawer width="half" placement="right" bind:open>
	<Heading tag="h5" class="mb-6 text-sm font-semibold uppercase">
		{title || (data.id ? 'Edit competitor' : 'Add new competitor')}
	</Heading>

    <SiteForm {data} {afterUpdate} bind:open />
</Drawer>
