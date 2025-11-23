<script lang="ts" module>
	import type { CompetitorsRecord } from "$lib/pocketbase/generated-types";
	export type CompetitorWithChanges = CompetitorsRecord & { changes: any[] };
</script>
<script lang="ts">
	const dfd = Intl.DateTimeFormat(navigator.language);
    let { competitors } = $props();
</script>

<div class="flex flex-col space-y-10">
	{#each competitors as competitor}
		<section class="rounded-3xl p-7 flex items-top">
			<div class="w-[20%] min-w-[200px] flex flex-col items-center text-center">
				<img
					class="w-20 h-20 object-contain mb-4"
					src={competitor.logo}
					alt={competitor.name + ' logo'}
					onerror={() => {
						/* fallback if logo missing */
					}}
				/>
				<div class="text-2xl font-bold text-gray-500 dark:text-white mb-1 capitalize">
					{competitor.name}
				</div>
				<a
					class="text-gray-500 hover:underline text-sm mb-4 shadow"
					href={competitor.url}
					target="_blank"
					rel="noopener noreferrer"
				>
					{competitor.url}
				</a>
			</div>
			<div class="w-[80%]">
				<ul class="space-y-4">
					{#each competitor.changes as change}
						<li class="relative pl-4 border-l-4 dark:border-gray-600">
							<div class="text-gray-500 dark:text-white max-w-[800px]">{change.narrative}</div>
							<div class="flex items-center gap-2 mt-1 text-xs text-gray-600 dark:text-gray-500">
								<span>{dfd.format(change.date)}</span>
								<a href={change.rurl} target="_blank">{change.rurl}</a>
							</div>
						</li>
					{/each}
				</ul>
			</div>
		</section>
	{/each}
</div>
