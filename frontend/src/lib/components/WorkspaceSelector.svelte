<script lang="ts">
	import { client } from '$lib/pocketbase';
	import type { WorkspacesRecord } from '$lib/pocketbase/generated-types';
	import { Button, Input, Label, Modal, Select } from 'flowbite-svelte';
	import { onMount } from 'svelte';

    type Option = { name: string, value: string }

    let { value = $bindable(), open = $bindable(false), afterCreate } = $props();
    let workspaces = $state<Option[]>([])
	let select = $state<any>();
    let error = $state('');

    onMount(() => {
        loadWorkspaces()
    })

    async function loadWorkspaces() {
        workspaces = await client.collection('workspaces').getFullList<WorkspacesRecord>()
            .then(ws => ws.map(({ id, name }) => ({ name: name as string, value: id })))
    }

	function onaction({ data }: { data: any }) {
		error = '';
		client
			.collection('workspaces')
			.create(data)
			.then(({ id }) => {
                const ws = (client.authStore.record?.workspaces || [])
                    .filter((w: string) => workspaces.find(ow => ow.value === w))
                    .concat(id);
                client.collection('users').update(
                    client.authStore.record?.id as string,
                    { workspaces: ws }
                )
                .then(() => {
                    loadWorkspaces()
                    open = false
                    afterCreate?.(id)
                })
            })
	}
</script>

<Select
    bind:this={select}
    bind:value={value}
    underline
    size="lg"
    items={workspaces}
    class="mb-6"
/>
<a href="" onclick={() => (open = true)}>or create a new one</a>

<Modal form bind:open size="xs" {onaction}>
	<div class="flex flex-col space-y-6">
		{#if error}
			<Label color="red">{error}</Label>
		{/if}
		<Label class="space-y-2">
			<span>Name</span>
			<Input name="name" placeholder="New workspace" required />
		</Label>
		<Button type="submit" value="create">Create</Button>
	</div>
</Modal>
