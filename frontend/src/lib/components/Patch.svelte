<script lang="ts">
    import { onMount } from "svelte";
    import { parsePatch, type StructuredPatch } from "diff";
	import type { SnapshotsRecord } from "$lib/pocketbase/generated-types";

    let { event }: { event: SnapshotsRecord } = $props();

    let display = $state<Element>();
    let formattedPatch = $state<StructuredPatch[]>();
    // const fragment = document.createDocumentFragment();

    onMount(() => {
        if (!event?.diff) return;
        formattedPatch = parsePatch(event.diff as string);

        // diff.forEach(part => {
        //     const color = part.added  ? 'green' : part.removed ? 'red' : 'grey';
        //     const span = document.createElement('span');
        //     span.style.color = color;
        //     span.appendChild(document.createTextNode(part.value.toString()));
        //     fragment.appendChild(span);
        // });

        // display?.appendChild(fragment);
    });
</script>

<pre bind:this={display}>{formattedPatch}</pre>

