<script lang="ts">
  let { websites } = $props();
</script>

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>URL</th>
      <th>Last Checked</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    {#each websites as site}
      <tr>
        <td>{site.name}</td>
        <td>{site.url}</td>
        <td>{site.last_checked}</td>
        <td>
          <!-- Example action: view details -->
          <a href={`/dashboard/${site.id}`}>View</a>
        </td>
      </tr>
    {/each}
  </tbody>
</table>
