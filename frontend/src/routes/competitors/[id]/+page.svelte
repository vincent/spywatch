<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import { client } from '$lib/pocketbase';
	import type { CompetitorsRecord, SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import EventsTimeline from '$lib/components/EventsTimeline.svelte';
	import CompetitorLogo from '$lib/components/CompetitorLogo.svelte';

	let competitor = $state<CompetitorsRecord>({ name: page.params.id as string } as CompetitorsRecord)
	onMount(() => {
		client.collection('competitors').getOne<CompetitorsRecord>(page.params.id as string).then(c => competitor = c);
	})
	const history = () => client.collection('snapshots').getFullList<SnapshotsRecord>({
		filter: `resource.competitor='${page.params.id}' && narrative!=''`
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home> Home </BreadcrumbItem>
			<BreadcrumbItem href={`/competitors`}> Competitors </BreadcrumbItem>
			<BreadcrumbItem> {competitor.name} </BreadcrumbItem>
		</Breadcrumb>
		<Heading tag="h1" class="flex items-center text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">
			<CompetitorLogo {competitor} />
			{competitor.name}
		</Heading>
	</div>

	{#await history() }
		<p>Loading history...</p>
	{:then versions} 
		{#if versions.length === 0}
			<p>No history available.</p>
		{:else if versions.length === 1}
			<h3>First content</h3>
			<pre>{versions[0].content}</pre>
		{:else}
			<EventsTimeline events={versions} />
		{/if}
	{/await}
</main>
