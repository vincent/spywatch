<script lang="ts">
	import type { CompetitorWithChanges } from '$lib/components/Release.svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import { client } from '$lib/pocketbase';
	import Release from '$lib/components/Release.svelte';
	import MonthNavigator, { DEFAULT_BOUNDS } from '$lib/components/MonthNavigator.svelte';

	const dfm = Intl.DateTimeFormat(navigator.language, { month: 'long', year: 'numeric' });
	let competitors = $state<CompetitorWithChanges[]>();
	let bounds = $state(DEFAULT_BOUNDS)

	const fetchSnapshots = (bounds: { min: Date; max: Date }) => {
		return client
			.collection('snapshots')
			.getFullList<SnapshotsRecord & { expand: any }>({
				fields: 'competitor,narrative,created,expand',
				filter: `narrative!='' && created>='${bounds.min.toISOString()}' && created<='${bounds.max.toISOString()}'`,
				expand: 'resource.competitor'
			})
			.then((snapshots) => {
				const merged = snapshots.reduce((acc, { created, narrative, expand }) => {
					const { id: cid, name, url, logo } = expand.resource.expand.competitor;
					const { url: rurl } = expand.resource;
					if (!acc[cid]) acc[cid] = { id: cid, name, url, logo };
					acc[cid].changes = (acc[cid].changes || []).concat({
						rurl,
						created,
						narrative
					});
					return acc;
				}, {} as any);
				competitors = Object.values(merged);
			});
	};

	$effect(() => {
		fetchSnapshots(bounds);
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home>Home</BreadcrumbItem>
			<BreadcrumbItem>Releases</BreadcrumbItem>
		</Breadcrumb>
		<div class="flex justify-between">
			<Heading tag="h3" class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">
				Releases - {dfm.format(bounds.max)}
			</Heading>
			<MonthNavigator bind:bounds={bounds} />
		</div>
	</div>

	<Release {competitors} />
</main>
