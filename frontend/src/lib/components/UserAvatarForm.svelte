<script lang="ts">
	import { client } from '$lib/pocketbase';
	import { Avatar, Fileupload, Helper, Label } from 'flowbite-svelte';

	let { user } = $props();
    let avatarSrc = $derived(client.files.getURL(user as any, user.avatar))

    function upload(e: Event & { currentTarget: EventTarget & HTMLInputElement; }) {
        const avatar = e.currentTarget.files?.[0]
        client.collection('users')
            .update(user?.id as string, { avatar })
            .then(u => (user = u))
    }
</script>

<div class="flex w-[500px]">
    <Avatar size="xl" src={avatarSrc} />
    <div class="ms-5">
        <Label for="with_helper" class="pb-2">Upload file</Label>
        <Fileupload id="with_helper" class="mb-2" onchange={upload} />
        <Helper>SVG, PNG, JPG or GIF (MAX. 500x500px).</Helper>
    </div>
</div>
