<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import { Rss } from '@lucide/svelte';
	import { client } from '$lib/pocketbase';
	import type { BodiesRecord, SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import EventsTimeline from '$lib/components/EventsTimeline.svelte';
	import CompetitorLogo from '$lib/components/CompetitorLogo.svelte';

	let body = $state<BodiesRecord>({ name: page.params.id as string } as BodiesRecord)
	onMount(() => {
		client.collection('bodies').getOne<BodiesRecord>(page.params.id as string).then(c => body = c);
	})
	const history = () => client.collection('snapshots').getFullList<SnapshotsRecord>({
		filter: `resource.body='${page.params.id}' && narrative!=''`
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home> Home </BreadcrumbItem>
			<BreadcrumbItem href={`/entities`}> Entities </BreadcrumbItem>
			<BreadcrumbItem> {body.name} </BreadcrumbItem>
		</Breadcrumb>
		<Heading tag="h1" class="flex items-center text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">
			<CompetitorLogo {body} />
			{body.name}
			<a class="ms-3" target="_blank" href={`/api/rss/${client.authStore.record?.api_key}/${body.id}`}><Rss/></a>
		</Heading>
	</div>
	{#await history() }
		<p>Loading history...</p>
	{:then events} 
		{#if events.length === 0}
			<p>No history available.</p>
		{:else if events.length === 1}
			<h3>First content</h3>
			<pre>{events[0].content}</pre>
		{:else}
			<EventsTimeline {events} />
		{/if}
	{/await}
</main>
