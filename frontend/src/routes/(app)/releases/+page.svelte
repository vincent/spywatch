<script lang="ts">
	import type { CompetitorWithChanges } from '$lib/components/Release.svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
	import { client } from '$lib/pocketbase';
	import Release from '$lib/components/Release.svelte';
	import MonthNavigator, { DEFAULT_BOUNDS } from '$lib/components/MonthNavigator.svelte';
	import WorkspaceSelector from '$lib/components/WorkspaceSelector.svelte';

	const dfm = Intl.DateTimeFormat(navigator.language, { month: 'long', year: 'numeric' });
	let bodies = $state<CompetitorWithChanges[]>();
	let bounds = $state(DEFAULT_BOUNDS)
	let workspace = $state('')

	const fetchSnapshots = (workspace: string, bounds: { min: Date; max: Date }) => {
		return client
			.collection('snapshots')
			.getFullList<SnapshotsRecord & { expand: any }>({
				fields: 'body,narrative,created,expand',
				filter: `narrative!='' && created>='${bounds.min.toISOString()}' && created<='${bounds.max.toISOString()}' ${workspace ? `&& resource.body.workspace='${workspace}'` : '' }`,
				expand: 'resource.body'
			})
			.then((snapshots) => {
				const merged = snapshots.reduce((acc, { created, narrative, expand }) => {
					const { id: cid, name, url, logo } = expand.resource.expand.body;
					const { url: rurl } = expand.resource;
					if (!acc[cid]) acc[cid] = { id: cid, name, url, logo };
					acc[cid].changes = (acc[cid].changes || []).concat({
						rurl,
						created,
						narrative
					});
					return acc;
				}, {} as any);
				bodies = Object.values(merged);
			});
	};

	$effect(() => {
		fetchSnapshots(workspace, bounds);
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home>Home</BreadcrumbItem>
			<BreadcrumbItem>Releases</BreadcrumbItem>
		</Breadcrumb>
		<div class="flex justify-between items-center">
			<Heading tag="h3" class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">
				Releases - {dfm.format(bounds.max)}
			</Heading>
			<WorkspaceSelector
				className="w-50 ms-5 me-auto pt-6"
				bind:value={workspace}
				none={"Any workspace"}
			/>
			<MonthNavigator bind:bounds={bounds} />
		</div>
	</div>

	<Release {bodies} />
</main>
