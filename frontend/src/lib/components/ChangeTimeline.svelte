<script lang="ts">
  // Props: changes (array of change objects)
  export let changes = [];
</script>

<section>
  <h3>Change Timeline</h3>
  {#if changes.length === 0}
    <p>No changes detected yet.</p>
  {:else}
    <ul>
      {#each changes as change}
        <li>
          <strong>{change.timestamp}</strong>: {change.diff_summary}
          {#if change.diff_url}
            [<a href={change.diff_url} target="_blank">diff</a>]
          {/if}
          {#if change.snapshot_url}
            [<a href={change.snapshot_url} target="_blank">snapshot</a>]
          {/if}
        </li>
      {/each}
    </ul>
  {/if}
</section>
