<!-- @wc-ignore-file -->
<script lang="ts">
	import { Button, Timeline, TimelineItem } from 'flowbite-svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { ArrowRightOutline } from 'flowbite-svelte-icons';

	let { events }: { events: SnapshotsRecord[] } = $props();
</script>

<Timeline>
	{#each events as e}
		<TimelineItem
			title=""
			date={Intl.DateTimeFormat('fr-FR').format(Date.parse(e.created as string))}
		>
			<p class="mb-4 text-base font-normal text-gray-500 dark:text-gray-400">
				{e.narrative}
			</p>
			<Button color="alternative" href={`/competitors/${e.resource}/diff/${e.id}`}
				>Learn more<ArrowRightOutline class="ms-2 h-5 w-5" /></Button
			>
		</TimelineItem>
	{/each}
</Timeline>
