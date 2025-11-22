<script lang="ts">
	import { client } from "$lib/pocketbase";
	import type { ResourcesRecord } from "$lib/pocketbase/generated-types";
	import type { RecordModel } from "pocketbase";
	import { onMount } from "svelte";

	// Props: site (optional, for editing)
	let { site = { id: "", name: "", url: "", logo: "" } } = $props();

	let resources = $state<ResourcesRecord[]>([]);
	let selectedResources = $state<Record<string, boolean>>({});

	onMount(async () => {
		if (!site.name) return;
		selectedResources = (await client.collection("resources").getFullList({ filter: `competitor='${site.id}'` })).reduce((acc, r) => ({ ...acc, [`${r.url}`]: true }), {});
		resources = site.found_resources.split('###').map((url: string) => ({ url }))
	});

	function reset() {
		resources = []
		selectedResources = {}
		handleSubmit()
	}

	async function handleSubmit(event?: Event) {
		event?.preventDefault();

		if (resources.length) {
			if (!site.id) {
				site = await client.collection("competitors").create({
					owner: client.authStore.record?.id,
					found_resources: resources.map(r => r.url).join('###'),
					...site
				});
			}

			const rInsert = Object.entries(selectedResources).filter(([url, selected]) => url && selected);
			for (let index = 0; index < rInsert.length; index++) {
				const [url] = rInsert[index];
				await client.collection("resources").create<ResourcesRecord>({
					competitor: site.id,
					url,
				});
			}

		} else {
			client.send("/api/preview", {
				method: "POST",
				body: site,
			}).then(p => {
				if (p.logo && !site.logo) site = { ...site, logo: p.logo }
				if (p.title && !site.name) site = { ...site, name: p.title }
				resources = (p.ownLinks as string[]).map(url => resources.find(r => r.url === url) ?? { id: '', url });
				selectedResources = resources.reduce((acc, r) => ({ ...acc, [r.url as string]: r.url?.replace(/\/$/, '') === site.url.replace(/\/$/, '') }), {});
			});
		}
	}
</script>

<div class="flex justify-center items-center min-h-screen bg-gray-50 py-8">
	<div class="w-full max-w-lg bg-white shadow-lg rounded-lg p-8">
		<h2 class="text-2xl font-bold mb-2 text-gray-800">Competitor Form</h2>
		<p class="mb-6 text-gray-500">Add or edit a competitor and manage their watched resources.</p>
		<form onsubmit={handleSubmit} class="space-y-6">
			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1" for="url">URL</label>
				<input
					id="url"
					type="url"
					bind:value={site.url}
					onchange={reset}
					required
					class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
					placeholder="https://competitor.com"
				/>
			</div>
			{#if site.url}
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1" for="name">Name</label>
					<div class="flex items-center">
						<input
							id="name"
							type="text"
							bind:value={site.name}
							required
							class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
							placeholder="Competitor name"
						/>
						{#if site.logo}
							<img
								class="w-10 h-10 ms-2 object-contain"
								src={site.logo}
								alt={site.name + ' logo'}
							/>
						{/if}
					</div>
				</div>
			{/if}
			{#if resources.length}
				<div>
					<span class="block text-sm font-medium text-gray-700 mb-2">Watched resources</span>
					<div class="space-y-2">
						{#each resources as res}
							<label class="flex items-center space-x-2">
								<input
									type="checkbox"
									bind:checked={selectedResources[res.url as string]}
									class="form-checkbox h-4 w-4 text-blue-600 border-gray-300 rounded"
								/>
								<span class="text-gray-700">{res.url}</span>
							</label>
						{/each}
					</div>
				</div>
			{/if}
			<div class="flex justify-end">
				<button
					type="submit"
					class="inline-flex items-center px-6 py-2 bg-blue-600 text-white font-semibold rounded-md shadow hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
				>
					{site && site.id
						? resources.length
							? "Update competitor"
							: "Update competitor"
						: resources.length
							? "Add competitor"
							: "Preview"}
				</button>
			</div>
		</form>
	</div>
</div>
