<script lang="ts">
	import { goto } from "$app/navigation";
	import { client } from "$lib/pocketbase";
	import { onMount } from "svelte";

    onMount(() => {
        if (client.authStore.isValid) goto('/entities')
    })
</script>

<!-- SpyWatch Homepage -->
<div class="min-h-screen flex flex-col bg-gray-50">
	<header class="w-full py-8 bg-white shadow-sm">
		<div class="max-w-4xl mx-auto flex items-center justify-between px-4">
			<div class="flex items-center space-x-3">
				<!-- Logo (optional, fallback to text) -->
				<img src="/favicon.svg" alt="SpyWatch Logo" class="h-8 w-8" />
				<span class="text-2xl font-bold text-gray-800 tracking-tight">SpyWatch</span>
			</div>
            <div>
                <a href="/auth/login" class="px-5 py-2 my-2 rounded text-gray-500 font-medium hover:bg-gray-100 transition">Login</a>
                <a href="/auth/signup" class="px-5 py-2 rounded bg-blue-600 text-white font-medium hover:bg-blue-700 transition">Sign Up</a>
            </div>
		</div>
	</header>

	<main class="flex-1 flex flex-col items-center justify-center px-4">
		<section class="max-w-2xl text-center mt-16">
			<h1 class="text-4xl md:text-5xl font-extrabold text-gray-900 mb-4">Stay Ahead. Effortlessly.</h1>
			<p class="text-lg md:text-xl text-gray-600 mb-8">
				SpyWatch is your SaaS solution to monitor updates on targeted websites, competitors, or any pages of interest. Get notified instantly when something changes—no more manual checking.
			</p>
			<a href="/auth/signup" class="inline-block px-8 py-3 rounded bg-blue-600 text-white font-semibold text-lg shadow hover:bg-blue-700 transition">
				Get Started
			</a>
		</section>

		<!-- Optional: Feature highlights -->
		<section class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-8 max-w-4xl w-full">
			<div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
				<span class="text-blue-600 text-3xl mb-2">🔔</span>
				<h2 class="font-semibold text-lg mb-1">Instant Alerts</h2>
				<p class="text-gray-500 text-sm">Be the first to know when your tracked pages change.</p>
			</div>
			<div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
				<span class="text-blue-600 text-3xl mb-2">📊</span>
				<h2 class="font-semibold text-lg mb-1">Competitor Tracking</h2>
				<p class="text-gray-500 text-sm">Monitor competitors and industry trends with ease.</p>
			</div>
			<div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
				<span class="text-blue-600 text-3xl mb-2">⚡</span>
				<h2 class="font-semibold text-lg mb-1">Simple & Powerful</h2>
				<p class="text-gray-500 text-sm">Easy setup, powerful insights, and a clean interface.</p>
			</div>
		</section>

		<!-- Optional: Screenshot for visual interest -->
		<section class="mt-16 mb-8">
			<img src="/assets/screenshot1.png" alt="SpyWatch Dashboard Screenshot" class="rounded-lg shadow-lg border border-gray-200 max-w-full h-auto mx-auto" style="max-height: 400px;" />
		</section>
	</main>

	<footer class="w-full py-6 text-center text-gray-400 text-sm border-t mt-8">
		&copy; {new Date().getFullYear()} SpyWatch. All rights reserved.
	</footer>
</div>
