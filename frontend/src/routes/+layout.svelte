<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import { client } from '$lib/pocketbase';
	import { redirect } from '@sveltejs/kit';

	let { children } = $props();

	$effect.pre(() => {
		if (client.authStore.isValid) return;
		debugger;
		if (!page.url.pathname.startsWith('/auth')) {
			redirect(301, '/auth/login');
		}
	});
</script>

{@render children?.()}
