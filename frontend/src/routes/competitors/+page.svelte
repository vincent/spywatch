<script lang="ts">
  import { onMount } from "svelte";
	import { client } from "$lib/pocketbase";
	import type { CompetitorsRecord } from "$lib/pocketbase/generated-types";

  let competitors = $state<CompetitorsRecord[]>();
  onMount(async () => competitors = 
    await client.collection('competitors').getFullList<CompetitorsRecord>()
  )
</script>

<main class="min-h-screen bg-gray-50 py-10 px-4">
  <div class="max-w-4xl mx-auto">
    <header class="mb-10 text-center">
      <h2 class="text-3xl font-extrabold text-gray-900 mb-2">Competitors to Watch</h2>
      <p class="text-lg text-gray-600">Here is a selection of leading website change monitoring tools to keep an eye on:</p>
    </header>
    <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
      {#each competitors as competitor}
        <div class="bg-white rounded-2xl shadow-md hover:shadow-xl transition-shadow duration-200 p-6 flex flex-col items-center">
          {#if competitor.logo}
            <img
              class="w-16 h-16 object-contain mb-4"
              src={competitor.logo}
              alt={competitor.name + ' logo'}
              on:error={() => { /* fallback if logo missing */ }}
            />
          {/if}
          <div class="text-xl font-semibold text-gray-800 mb-2 text-center"><a href={competitor.url} target="_blank">{competitor.name}</a></div>
          <div class="text-gray-600 text-base mb-4 text-center"><a href={competitor.url} target="_blank">{competitor.url}</a></div>
          <div class="flex">
            <a
              class="mt-auto inline-block px-4 py-2 bg-blue-600 text-white font-medium rounded-lg shadow hover:bg-blue-700 transition-colors duration-150"
              href={`/competitors/${competitor.id}/edit`}
              rel="noopener noreferrer"
            >
              Edit
            </a>
            <a
              class="mt-auto inline-block px-4 py-2 bg-blue-600 text-white font-medium rounded-lg shadow hover:bg-blue-700 transition-colors duration-150"
              href={`/competitors/${competitor.id}`}
              rel="noopener noreferrer"
            >
              History
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
</main>
