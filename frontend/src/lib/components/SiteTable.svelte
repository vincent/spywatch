<script lang="ts">
	import { Button, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
	import type { BodiesResponse } from '$lib/pocketbase/generated-types';
	import { EditOutline, TrashBinSolid } from 'flowbite-svelte-icons';
	import CompetitorLogo from './CompetitorLogo.svelte';
	
	type Props = {
		bodies: BodiesResponse[],
		edit: CallableFunction,
		remove: CallableFunction
	}
	let { bodies, edit, remove }: Props = $props();
</script>

<Table>
	<TableHead class="border-y border-gray-200 bg-gray-100 dark:border-gray-700">
		<TableHeadCell class="ps-4 font-normal w-30">Name</TableHeadCell>
		<TableHeadCell class="ps-4 font-normal">URL</TableHeadCell>
		<TableHeadCell class="ps-4 font-normal">Workspace</TableHeadCell>
		<TableHeadCell class="ps-4 font-normal"></TableHeadCell>
	</TableHead>
	<TableBody>
		{#each bodies as body}
			<TableBodyRow class="text-base">
				<TableBodyCell class="space-x-6 p-4 whitespace-nowrap">
					<div class="text-sm font-normal text-gray-500 dark:text-gray-300">
						<div class="flex text-base font-semibold text-gray-900 dark:text-white">
							<CompetitorLogo {body} />
							<a href={`/entities/${body.id}`}>{body.name}</a>
						</div>
					</div>
				</TableBodyCell>
				<TableBodyCell class="p-4"
					><a href={body.url} target="_blank">{body.url}</a></TableBodyCell
				>
				<TableBodyCell class="p-4">{(body.expand as any)?.workspace?.name}</TableBodyCell>
				<TableBodyCell class="space-x-2 text-right">
					<Button
						size="sm"
						class="gap-2 px-3"
						outline
						onclick={() => edit(body)}
					>
						<EditOutline size="sm" /> Edit
					</Button>
					<Button
						color="red"
						size="sm"
						class="gap-2 px-3"
						onclick={() => remove(body)}
					>
						<TrashBinSolid size="sm" />
					</Button>
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>

