<script lang="ts">
	import { page } from '$app/state';
	import { client } from '$lib/pocketbase';
	import EventsTimeline from '$lib/components/EventsTimeline.svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';
	const history = () => client.collection('snapshots').getFullList<SnapshotsRecord>({
		filter: `resource.competitor='${page.params.id}'`
	});
</script>

<main>
	{#await history() }
		<p>Loading history...</p>
	{:then versions} 
		{#if versions.length === 0}
			<p>No history available.</p>
		{:else if versions.length === 1}
			<h3>First content</h3>
			<pre>{versions[0].content}</pre>
		{:else}
			{@const [left, right] = versions}
			<h3>Change History:</h3>
			<!-- <Diff {left} {right} /> -->
			<EventsTimeline events={versions} />
		{/if}
	{/await}
</main>
