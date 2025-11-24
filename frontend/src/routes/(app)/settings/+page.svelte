<script lang="ts">
	import UserAvatarForm from '$lib/components/UserAvatarForm.svelte';
	import { client } from '$lib/pocketbase';
  import { Breadcrumb, BreadcrumbItem, Heading } from 'flowbite-svelte';
</script>

<main class="p-4">
  <div class="grid grid-cols-1 space-y-2 xl:grid-cols-3 xl:gap-3.5">
    <div class="col-span-full xl:mb-0">
      <Breadcrumb class="mb-6">
        <BreadcrumbItem home>Home</BreadcrumbItem>
        <BreadcrumbItem>User settings</BreadcrumbItem>
        <BreadcrumbItem>{client.authStore.record?.email}</BreadcrumbItem>
      </Breadcrumb>

      <Heading tag="h1" class="text-xl font-semibold text-gray-900 sm:text-2xl dark:text-white">User settings</Heading>
    </div>
    <div class="col-span-full space-y-4 xl:col-auto">
      <UserAvatarForm user={client.authStore.record} />
    </div>
    <div class="col-span-2 space-y-4">
    </div>
  </div>
  <div class="mt-4 grid grid-cols-1 gap-4 xl:grid-cols-2 xl:gap-4">
  </div>
</main>
