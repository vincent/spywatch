<script lang="ts">
	import { Avatar, NavHamburger, Dropdown, DropdownHeader, DropdownGroup, DropdownItem } from "flowbite-svelte";
	import type { AuthRecord } from "pocketbase";
	import { UserIcon } from "@lucide/svelte";
	import { client } from "$lib/pocketbase";
	import { goto } from "$app/navigation";

    let user = client.authStore.record as AuthRecord;
    let avatarSrc = $derived(client.files.getURL(user as any, user?.avatar))

    function signout() {
        client.authStore.clear()
        goto('/auth/login')
    }
</script>

<div class="flex items-center md:order-2">
    {#if avatarSrc}
        <Avatar id="avatar-menu" src={avatarSrc} />
    {:else}
        <UserIcon/>
    {/if}
    <NavHamburger />
</div>
<Dropdown placement="bottom" triggeredBy="#avatar-menu">
	<DropdownHeader>
		<span class="block text-sm">{client.authStore.record?.name}</span>
		<span class="block truncate text-sm font-medium">{client.authStore.record?.email}</span>
	</DropdownHeader>
	<DropdownGroup>
		<DropdownItem href="/settings">Settings</DropdownItem>
		<DropdownItem>Earnings</DropdownItem>
	</DropdownGroup>
	<DropdownGroup>
    	<DropdownItem class="w-full text-start" onclick={signout}>Sign out</DropdownItem>
    </DropdownGroup>
</Dropdown>
