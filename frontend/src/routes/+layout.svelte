<script lang="ts">
	import '../app.css';
  import { loadLocale } from 'wuchale/load-utils';
  
  // Layout logic (e.g., load user session, provide context)
  export let data;

  loadLocale('fr')
</script>

<slot />
