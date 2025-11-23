<script lang="ts" module>
	const ONE8DAY = 1000 * 60 * 60 * 24
	const ONE_MONTH = ONE8DAY * 30
    export const DEFAULT_BOUNDS = {
		min: new Date(new Date().getTime() - ONE_MONTH),
		max: new Date()
	}
</script>
<script lang="ts">
    let { bounds = $bindable(DEFAULT_BOUNDS) } = $props()
	let current = $derived(bounds.max.getTime() > new Date().getTime() - ONE8DAY);

    const nav = {
		previous: () => bounds = {
			min: new Date(bounds.min.getTime() - ONE_MONTH),
			max: new Date(bounds.max.getTime() - ONE_MONTH),
		},
		next: () => bounds = {
			min: new Date(bounds.min.getTime() + ONE_MONTH),
			max: new Date(bounds.max.getTime() + ONE_MONTH),
		},
	}
</script>

<div class="text-gray-900 sm:text-2xl dark:text-white">
	<a href="#" onclick={(e) => (e.preventDefault(), nav.previous())}>⬅</a>&nbsp;<a
		class={current ? 'invisible' : ''}
		href="#"
		onclick={(e) => (e.preventDefault(), nav.next())}>⮕</a
	>
</div>
