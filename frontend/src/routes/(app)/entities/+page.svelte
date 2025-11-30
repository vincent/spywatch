<script lang="ts">
	import { Breadcrumb, BreadcrumbItem, Button, Heading } from 'flowbite-svelte';
	import { onMount, type Component } from 'svelte';
	import { Plus } from '@lucide/svelte';
	import type { BodiesResponse } from '$lib/pocketbase/generated-types';
	import CompetitorDrawer from '$lib/components/CompetitorDrawer.svelte';
	import DeleteDrawer from '$lib/components/DeleteDrawer.svelte';
	import SiteTable from '$lib/components/SiteTable.svelte';
	import { client } from '$lib/pocketbase';
	import WorkspaceSelector from '$lib/components/WorkspaceSelector.svelte';
	
	let open: boolean = $state(false);
	let current_workspace = $state('');
	let current_competitor = $state({});
	let DrawerComponent: Component = $state(CompetitorDrawer);

	const toggle = (component: Component) => {
		DrawerComponent = component;
		open = !open;
	};

	const bodies = client.collection('bodies');
	let list = $state<BodiesResponse[]>([]);
	async function afterUpdate() {
		list = await bodies.getFullList<BodiesResponse>({
			filter: current_workspace ? `workspace='${current_workspace}'` : undefined,
			expand: 'workspace'
		});
	}
	$effect(() => {
		afterUpdate()
	});
</script>

<main class="relative h-full w-full overflow-y-auto bg-white dark:bg-gray-800">
	<div class="p-4">
		<Breadcrumb class="mb-5">
			<BreadcrumbItem home>Home</BreadcrumbItem>
			<BreadcrumbItem>Entities</BreadcrumbItem>
		</Breadcrumb>
		<div class="flex justify-between items-center">
			<Heading tag="h3" class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white"
				>All entities</Heading
			>
			<WorkspaceSelector
				className="w-50 ms-5 me-auto pt-6"
				bind:value={current_workspace}
				none={"Any workspace"}
			/>
			<Button onclick={() => ((current_competitor = {}), toggle(CompetitorDrawer))}><Plus /> Add new entity watch</Button>
		</div>
	</div>

	<SiteTable
		bodies={list}
		edit={(c: BodiesResponse) => ((current_competitor = c), toggle(CompetitorDrawer))}
		remove={() => toggle(DeleteDrawer)}
	/>
</main>

<DrawerComponent bind:open data={current_competitor} {afterUpdate} />