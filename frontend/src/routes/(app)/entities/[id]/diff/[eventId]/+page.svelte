<script lang="ts">
	import { page } from '$app/state';
	import { client } from '$lib/pocketbase';
	import Patch from '$lib/components/Patch.svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';

	const fetchEvent = () => {
		if (!page.params.eventId) return;
		return client.collection('snapshots').getOne<SnapshotsRecord>(page.params.eventId)
	}
</script>

<main>
	{#await fetchEvent() }
		<p>Loading history...</p>
	{:then event}
		<h3>Change History:</h3>
		{#if event}
			<Patch {event} />
		{:else}
			<p>No diff data</p>
		{/if}
	{/await}
</main>
