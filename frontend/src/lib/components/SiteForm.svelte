<script lang="ts">
	import type { BodiesRecord, ResourcesRecord } from '$lib/pocketbase/generated-types';
	import { Button, Checkbox, Input, Label } from 'flowbite-svelte';
	import { CloseOutline } from 'flowbite-svelte-icons';
	import { DraftingCompass } from '@lucide/svelte';
	import { client } from '$lib/pocketbase';
	import { onMount } from 'svelte';
	import WorkspaceSelector from './WorkspaceSelector.svelte';

	let { open = $bindable(), afterUpdate = null, data = { id: '', name: '', url: '', logo: '' , workspace: '' } } = $props();

	let resources = $state<ResourcesRecord[]>([]);
	let selectedResources = $state<Record<string, boolean>>({});

	const db = {
		bodies: client.collection('bodies'),
		resources: client.collection('resources'),
	}

	onMount(async () => {
		if (!data.id) return;
		data = await db.bodies.getOne<BodiesRecord>(data.id as string);
		resources = (await db.resources.getFullList({ filter: `url!='' && body='${data.id}'` }))
		resources = resources.concat(data.found_resources.split('###').map((url: string) => resources.find((r) => r.url === url) ?? { id: '', url }))
		
		selectedResources = resources.reduce((acc, r) => ({ ...acc, [`${r.url}`]: !!r.id }), {});
	});

	function reset() {
		resources = [];
		selectedResources = {};
		handleSubmit();
	}

	function preview() {
		client
			.send('/api/preview', {
				method: 'POST',
				body: data
			})
			.then((p) => {
				if (p.logo && !data.logo) data = { ...data, logo: p.logo };
				if (p.title && !data.name) data = { ...data, name: p.title };

				resources = []
					.concat(p.ownLinks)
					.concat(p.socialNetworks)
					.map((url) => resources.find((r) => r.url === url) ?? { id: '', url });

				selectedResources = resources.reduce(
					(acc, r) => ({
						...acc,
						[r.url as string]: r.url?.replace(/\/$/, '') === data.url.replace(/\/$/, '')
					}),
					{}
				);
			});
	}

	async function handleSubmit(event?: Event) {
		event?.preventDefault();

		if (!resources.length) {
			preview()
			return;
		}

		if (!data.id) {
			data = await client.collection('bodies').create({
				owner: client.authStore.record?.id,
				found_resources: resources.map((r) => r.url).join('###'),
				...data
			});
		} else {
			data = await client.collection('bodies').update(data.id, {
				found_resources: resources.map((r) => r.url).join('###'),
				...data
			});
		}

		const toInsert = Object.entries(selectedResources)
			.filter(([url, selected]) => url && selected)
			.map(([u]) => resources.find(r => (r.url as string) === u));

		for (let index = 0; index < toInsert.length; index++) {
			const res = toInsert[index];
			await client.collection('resources').create<ResourcesRecord>({
				body: data.id,
				...res
			});
		}

		open = false
		afterUpdate?.()
	}
</script>

<form onsubmit={handleSubmit} class="space-y-6">
	<div class="space-y-4">
		<Label class="space-y-2">
			<span>URL</span>
			<Input
				name="url"
				class="border font-normal outline-none mt-2"
				placeholder="https://body.com"
				bind:value={data.url}
				onchange={reset}
				type="url"
				required
			/>
		</Label>
		{#if data.url}
			<Label class="space-y-2">
				<span>Workspace</span>
				<WorkspaceSelector bind:value={data.workspace} afterCreate={(workspace: string) => (data = {...data, workspace})} />
			</Label>
			<Label class="space-y-2">
				<span> Name </span>
				<div class="flex items-center mt-2">
					<Input
						name="name"
						class="border font-normal outline-none"
						placeholder="Competitor name"
						bind:value={data.name}
						type="text"
						required
					/>
					{#if data.logo}
						<img class="w-10 h-10 ms-2 object-contain" src={data.logo} alt={data.name + ' logo'} />
					{/if}
				</div>
			</Label>
		{/if}
		{#if resources.length}
			<Label class="space-y-2">
				<span> Watched resources </span>
			</Label>
			<div class="space-y-2 ps-4 mt-2">
				{#each resources as res}
					<div class="flex items-center space-x-2">
						<Label>
							<Checkbox
								bind:checked={selectedResources[res.url as string]}
							/>
							{res.url}
						</Label>
					</div>
				{/each}
			</div>
		{/if}
		<div class="bottom-0 left-0 flex w-full justify-center space-x-4 pb-4 md:px-4 mt-10">
			<Button type="button" onclick={preview} outline={!!resources.length}>
				<DraftingCompass />
				Preview
			</Button>
			{#if resources.length}
				<Button type="submit" class="w-full">
					{data && data.id
						? 'Update body'
						: 'Add body'}
				</Button>
			{/if}
			<Button color="alternative" class="w-full" onclick={() => (open = false)}>
				<CloseOutline />
				Cancel
			</Button>
		</div>
	</div>
</form>
