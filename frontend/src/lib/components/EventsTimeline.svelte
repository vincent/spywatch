<!-- @wc-ignore-file -->
<script lang="ts">
	import { Button, Timeline, TimelineItem } from 'flowbite-svelte';
	import type { SnapshotsRecord } from '$lib/pocketbase/generated-types';
	import { ArrowRightOutline } from 'flowbite-svelte-icons';

	let { events }: { events: SnapshotsRecord[] } = $props();
	let details = $state<Record<string, string>>({})
</script>

<Timeline>
	{#each events as e}
		<TimelineItem
			title=""
			date={Intl.DateTimeFormat('fr-FR').format(Date.parse(e.created as string))}
		>
			<p class="mb-4 text-base font-normal text-gray-500 dark:text-white">
				{e.narrative}
			</p>
			<div class="flex items-center">
				<Button color="alternative" onclick={() => (details[e.id] = 'diff')}> Show diff <ArrowRightOutline class="ms-2 h-5 w-5" /></Button>
				<Button color="alternative" onclick={() => (details[e.id] = 'content')}> Show content <ArrowRightOutline class="ms-2 h-5 w-5" /></Button>

			</div>
			{#if details[e.id] === 'diff'}
				<pre class="border bg-gray-500">{e.diff}</pre>
			{:else if details[e.id] === 'content'}
				<pre class="border bg-gray-500">{e.content}</pre>
			{/if}
		</TimelineItem>
	{/each}
</Timeline>
