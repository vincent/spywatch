<script lang="ts">
	import { client } from "$lib/pocketbase";
	import type { CompetitorsRecord, SnapshotsRecord } from "$lib/pocketbase/generated-types";

    type CompetitorWithChanges = CompetitorsRecord & { changes: any[] }

    const df = Intl.DateTimeFormat(navigator.language)
    let competitors = $state<CompetitorWithChanges[]>();
    let bounds = {
        min: new Date(new Date().getTime() - (1000 * 60 * 60 * 24 * 30)),
        max: new Date(),
    }

    const fetchSnapshots = (bounds: { min: Date, max: Date }) => {
		return client
            .collection('snapshots')
            .getFullList<SnapshotsRecord & { expand: any }>({
                fields: 'competitor,narrative,created,expand',
                filter: `narrative!='' && created>='${bounds.min.toISOString()}' && created<='${bounds.max.toISOString()}'`,
                expand: 'resource.competitor',
            })
            .then(snapshots => {
                const merged = snapshots.reduce((acc, { created, narrative, expand }) => {
                    const { id: cid, name, url, logo } = expand.resource.expand.competitor;
                    const { url: rurl } = expand.resource;
                    if (!acc[cid]) acc[cid] = { id: cid, name, url, logo };
                    acc[cid].changes = (acc[cid].changes || []).concat({
                        rurl, created, narrative
                    });
                    return acc;
                }, {} as any)
                competitors = Object.values(merged)
            })
	}

    $effect(() => {
        fetchSnapshots(bounds)
    })
</script>

<main class="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-100 py-12 px-4">
	<div class="max-w-6xl mx-auto">
		<header class="mb-12 text-center">
			<h1
				class="text-4xl sm:text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-700 via-purple-600 to-pink-500 mb-4 drop-shadow-lg"
			>
				Competitors Watch Report
			</h1>
			<p class="text-lg text-gray-700 max-w-2xl mx-auto">
				Stay ahead by tracking what your competitors are changing. See the latest updates, features,
				and announcements from leading website monitoring tools.
			</p>
		</header>
		<div class="flex flex-col space-y-10">
			{#each competitors as competitor}
				<section
					class="bg-white rounded-3xl shadow-2xl hover:shadow-pink-200 transition-shadow duration-300 p-7 flex items-center border-2 border-transparent hover:border-pink-400"
				>
					<div class="w-[25%] flex flex-col items-center">
						<img
							class="w-20 h-20 object-contain mb-4 rounded-full border-4 border-blue-200 shadow"
							src={competitor.logo}
							alt={competitor.name + ' logo'}
							on:error={() => {
								/* fallback if logo missing */
							}}
						/>
						<div class="text-2xl font-bold text-gray-800 mb-1 text-center">{competitor.name}</div>
						<a
							class="text-blue-600 hover:underline text-sm mb-4"
							href={competitor.url}
							target="_blank"
							rel="noopener noreferrer"
						>
							{competitor.url}
						</a>
					</div>
					<div class="w-[70%]">
						<ul class="space-y-4">
							{#each competitor.changes as change}
								<li class="relative pl-4 border-l-4">
									<div class="flex items-center gap-2 mb-1">
										<span class="text-xs text-gray-400">{df.format(change.date)}</span>
										<span class="text-xs text-gray-400"><a href={change.rurl} target="_blank">{change.rurl}</a></span>
									</div>
									<div class="text-gray-600 text-sm">{change.narrative}</div>
								</li>
							{/each}
						</ul>
					</div>
				</section>
			{/each}
		</div>
		<footer class="mt-16 text-center text-gray-400 text-xs">
			&copy; {new Date().getFullYear()} Competitors Watch Report &mdash; Demo UI
		</footer>
	</div>
</main>
