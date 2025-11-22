<!-- @wc-ignore-file -->
<script lang="ts">
	import type { SnapshotsRecord } from "$lib/pocketbase/generated-types";

    let { events }: { events: SnapshotsRecord[] } = $props();
</script>

<ol class="relative border-s border-default">
    {#each events as e}
        <li class="mb-10 ms-4">
            <div class="absolute w-3 h-3 bg-neutral-quaternary rounded-full mt-1.5 -start-1.5 border border-buffer"></div>
            <time class="text-sm font-normal leading-none text-body">{Intl.DateTimeFormat('fr-FR').format(Date.parse(e.created as string))}</time>
            <p class="mb-4 text-base font-normal text-body">{e.narrative}</p>
            <a href={`/dashboard/${e.resource}/diff/${e.id}`} class="inline-flex items-center text-body bg-neutral-secondary-medium box-border border border-default-medium hover:bg-neutral-tertiary-medium hover:text-heading focus:ring-4 focus:ring-neutral-tertiary shadow-xs font-medium leading-5 rounded-base text-sm px-4 py-2.5 focus:outline-none">
                Learn more
                <svg class="w-4 h-4 ms-1.5 -me-0.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12H5m14 0-4 4m4-4-4-4"/></svg>
            </a>
        </li>        
    {/each}
</ol>
