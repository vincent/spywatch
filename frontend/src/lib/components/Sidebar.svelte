<script lang="ts">
  import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper, SidebarButton, uiHelpers } from 'flowbite-svelte';
  import { CogOutline, GithubSolid, ChartPieOutline  } from 'flowbite-svelte-icons';
  import { afterNavigate } from '$app/navigation';
	import { Newspaper } from '@lucide/svelte';

  interface Props {
    drawerHidden: boolean;
  }
  let { drawerHidden = $bindable(false) }: Props = $props();
  const closeDrawer = () => {
    drawerHidden = true;
  };

  let iconClass = 'flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-300 dark:group-hover:text-white';
  let itemClass = 'flex items-center p-2 text-base text-gray-900 transition duration-75 rounded-lg hover:bg-gray-100 group dark:text-gray-200 dark:hover:bg-gray-700 w-full';
  let groupClass = 'pt-2 space-y-2 mb-3';

  const sidebarUi = uiHelpers();
  let isOpen = $state(false);
  const closeSidebar = sidebarUi.close;
  $effect(() => {
    isOpen = sidebarUi.isOpen;
  });

  afterNavigate(() => {
    // this fixes https://github.com/themesberg/flowbite-svelte/issues/364
    document.getElementById('svelte')?.scrollTo({ top: 0 });
    closeDrawer();
  });

  let items = [
    { name: 'Entities', Icon: ChartPieOutline, href: '/entities' },
    { name: 'Releases', Icon: Newspaper, href: '/releases' },
    { name: 'Settings', Icon: CogOutline, href: '/settings' },
  ];

  let links = [
    {
      label: 'Feedback & Issues',
      href: 'https://github.com/vincent/spywatch/issues',
      Icon: GithubSolid
    },
  ];
</script>

<SidebarButton breakpoint="lg" onclick={sidebarUi.toggle} class="fixed top-[22px] z-40 mb-2" />

<Sidebar
  breakpoint="lg"
  backdrop={false}
  {isOpen}
  {closeSidebar}
  params={{ x: -50, duration: 50 }}
  class="top-0 left-0 mt-[69px] h-screen w-64 bg-gray-50 transition-transform lg:block dark:bg-gray-800"
  classes={{ div: 'h-full py-1 overflow-y-auto bg-gray-50 dark:bg-gray-800', nonactive: 'p-2', active: 'p-2' }}
>
  <h4 class="sr-only">Main menu</h4>
  <SidebarWrapper class="scrolling-touch h-full max-w-2xs overflow-y-auto bg-white px-3 pt-20 lg:sticky lg:me-0 lg:block lg:h-[calc(100vh-4rem)] lg:pt-5 dark:bg-gray-800">
    <SidebarGroup class={groupClass}>
      {#each items as { name, Icon, href } (name)}
        <SidebarItem label={name} {href} spanClass="ml-3" class={itemClass} aClass="w-full p-0 py-2">
          {#snippet icon()}
            <Icon class={iconClass} />
          {/snippet}
        </SidebarItem>
      {/each}
    </SidebarGroup>
    <SidebarGroup class={groupClass}>
      {#each links as { label, href, Icon } (label)}
        <SidebarItem {label} {href} spanClass="ml-3" class={itemClass}>
          {#snippet icon()}
            <Icon class={iconClass} />
          {/snippet}
        </SidebarItem>
      {/each}
    </SidebarGroup>
  </SidebarWrapper>
</Sidebar>
