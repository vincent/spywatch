<script lang="ts">
	import { Plus } from '@lucide/svelte';
	import { type Component } from 'svelte';
	import { client } from '$lib/pocketbase';
	import { Breadcrumb, BreadcrumbItem, Button, Heading } from 'flowbite-svelte';
	import EntityDeleteDrawer from '$lib/components/EntityDeleteDrawer.svelte';
	import WorkspaceSelector from '$lib/components/WorkspaceSelector.svelte';
	import EntityFormDrawer from '$lib/components/EntityFormDrawer.svelte';
	import type { BodiesResponse } from '$lib/pocketbase/generated-types';
	import SiteTable from '$lib/components/SiteTable.svelte';
	
	let DrawerComponent: Component = $state(EntityFormDrawer);
	let open: boolean = $state(false);
	let workspace = $state('');
	let entity = $state({});

	const toggle = (component: Component) => {
		DrawerComponent = component;
		open = !open;
	};

	const bodies = client.collection('bodies');
	let list = $state<BodiesResponse[]>([]);
	async function afterUpdate() {
		list = await bodies.getFullList<BodiesResponse>({
			filter: workspace ? `workspace='${workspace}'` : undefined,
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
				bind:value={workspace}
				none={"Any workspace"}
			/>
			<Button onclick={() => ((entity = {}), toggle(EntityFormDrawer))}><Plus /> Add new entity watch</Button>
		</div>
	</div>

	<SiteTable
		bodies={list}
		edit={(c: BodiesResponse) => ((entity = c), toggle(EntityFormDrawer))}
		remove={(c: BodiesResponse) => ((entity = c), toggle(EntityDeleteDrawer))}
	/>
</main>

<DrawerComponent bind:open data={entity} callback={afterUpdate} />