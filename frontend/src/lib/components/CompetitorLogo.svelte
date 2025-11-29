<script lang="ts">
    let { body, size = 5 } = $props();
</script>

{#if body.logo}
    <img
        class="w-{size} h-{size} me-2 object-contain"
        src={body.logo}
        alt={body.name + ' logo'}
    />
{/if}
