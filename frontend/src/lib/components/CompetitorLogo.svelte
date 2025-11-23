<script lang="ts">
    let { competitor, size = 5 } = $props();
</script>

{#if competitor.logo}
    <img
        class="w-{size} h-{size} me-2 object-contain"
        src={competitor.logo}
        alt={competitor.name + ' logo'}
    />
{/if}
