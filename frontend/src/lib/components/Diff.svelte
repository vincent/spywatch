<script lang="ts">
	import { Diff } from "diff";
	import { onMount } from "svelte";

    let { right, left } = $props();

    let display = $state<Element>();
    const fragment = document.createDocumentFragment();

    onMount(() => {
        var diff = new Diff<string>().diff(left, right, {
            ignoreCase: true,
        });

        diff.forEach(part => {
            const color = part.added  ? 'green' : part.removed ? 'red' : 'grey';
            const span = document.createElement('span');
            span.style.color = color;
            span.appendChild(document.createTextNode(part.value.toString()));
            fragment.appendChild(span);
        });

        display?.appendChild(fragment);
    });
</script>

<pre bind:this={display}></pre>

