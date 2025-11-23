<script lang="ts">
	import { Breadcrumb, BreadcrumbItem, Button, Heading } from 'flowbite-svelte';
	import { onMount, type Component } from 'svelte';
	import { Plus } from '@lucide/svelte';
	import type { CompetitorsRecord } from '$lib/pocketbase/generated-types';
	import CompetitorDrawer from '$lib/components/CompetitorDrawer.svelte';
	import DeleteDrawer from '$lib/components/DeleteDrawer.svelte';
	import SiteTable from '$lib/components/SiteTable.svelte';
	import { client } from '$lib/pocketbase';
	
	let open: boolean = $state(false);
	let current_competitor: any = $state({});
	let DrawerComponent: Component = $state(CompetitorDrawer);

	const toggle = (component: Component) => {
		DrawerComponent = component;
		open = !open;
	};

	const competitors = client.collection('competitors');
	let list = $state<CompetitorsRecord[]>([]);
	async function afterUpdate() {
		list = await competitors.getFullList<CompetitorsRecord>();
	}
	onMount(async () => {
		afterUpdate()
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home>Home</BreadcrumbItem>
			<BreadcrumbItem>Competitors</BreadcrumbItem>
		</Breadcrumb>
		<div class="flex justify-between">
			<Heading tag="h3" class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white"
				>All competitors</Heading
			>
			<Button onclick={() => ((current_competitor = {}), toggle(DrawerComponent))}><Plus /> Add new competitor watch</Button>
		</div>
	</div>

	<SiteTable
		competitors={list}
		edit={(c: CompetitorsRecord) => ((current_competitor = c), toggle(DrawerComponent))}
		remove={() => toggle(DeleteDrawer)}
	/>
</main>

<DrawerComponent bind:open data={current_competitor} {afterUpdate} />