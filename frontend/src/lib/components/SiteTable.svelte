<script lang="ts">
	import { Button, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
	import type { CompetitorsRecord } from '$lib/pocketbase/generated-types';
	import { EditOutline, TrashBinSolid } from 'flowbite-svelte-icons';
	import CompetitorLogo from './CompetitorLogo.svelte';
	
	type Props = {
		competitors: CompetitorsRecord[],
		edit: CallableFunction,
		remove: CallableFunction
	}
	let { competitors, edit, remove }: Props = $props();
</script>

<Table>
	<TableHead class="border-y border-gray-200 bg-gray-100 dark:border-gray-700">
		<TableHeadCell class="ps-4 font-normal w-30">Name</TableHeadCell>
		<TableHeadCell class="ps-4 font-normal">URL</TableHeadCell>
		<TableHeadCell class="ps-4 font-normal"></TableHeadCell>
	</TableHead>
	<TableBody>
		{#each competitors as competitor}
			<TableBodyRow class="text-base">
				<TableBodyCell class="space-x-6 p-4 whitespace-nowrap">
					<div class="text-sm font-normal text-gray-500 dark:text-gray-300">
						<div class="flex text-base font-semibold text-gray-900 dark:text-white">
							<CompetitorLogo {competitor} />
							<a href={`/competitors/${competitor.id}`}>{competitor.name}</a>
						</div>
					</div>
				</TableBodyCell>
				<TableBodyCell class="p-4"
					><a href={competitor.url} target="_blank">{competitor.url}</a></TableBodyCell
				>
				<TableBodyCell class="space-x-2 text-right">
					<Button
						size="sm"
						class="gap-2 px-3"
						outline
						onclick={() => edit(competitor)}
					>
						<EditOutline size="sm" /> Update
					</Button>
					<Button
						color="red"
						size="sm"
						class="gap-2 px-3"
						onclick={() => remove(competitor)}
					>
						<TrashBinSolid size="sm" />
					</Button>
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>

