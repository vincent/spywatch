<script lang="ts">
  import { DarkMode, NavBrand, Navbar, Search } from 'flowbite-svelte';
	import UserMenu from './UserMenu.svelte';

  let { drawerHidden = $bindable(false) } = $props();
</script>

<Navbar class="mx-10 sm:mx-0">
  <NavBrand href="/" class="mx-10">
    <img src="/favicon.svg" class="me-2.5 h-6 sm:h-8" alt="Logo" />
    <span class="ml-px self-center text-xl font-semibold whitespace-nowrap sm:text-2xl dark:text-white"> SpyWatch </span>
  </NavBrand>
  <form>
    <Search size="md" class="mt-1 w-96 border focus:outline-none" />
  </form>
  <div class="ms-auto me-4 flex items-center text-gray-500 sm:order-2 dark:text-gray-300">
    <DarkMode />
  </div>
  <UserMenu />
</Navbar>
