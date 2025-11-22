<script>
  // No JS needed for static landing page, but could add interactivity if desired
  const screenshots = [
    '/assets/screenshot1.png',
    '/assets/screenshot2.png',
    '/assets/screenshot3.png'
  ];
</script>

<!-- Hero Section -->
<section class="bg-gradient-to-br from-blue-900 via-blue-700 to-blue-500 text-white min-h-[60vh] flex flex-col justify-center items-center px-4 py-20">
  <h1 class="text-5xl md:text-7xl font-extrabold mb-4 drop-shadow-lg">SpyWatch</h1>
  <p class="text-2xl md:text-3xl font-semibold mb-6">Stay ahead of your competitors, effortlessly.</p>
  <p class="max-w-2xl text-lg md:text-xl mb-8 text-blue-100">
    SpyWatch monitors your competitors' websites and notifies you of every change. Track pricing, new features and more, automatically.
  </p>
  <a href="/auth/signup" class="bg-white text-blue-800 font-bold px-8 py-3 rounded-full shadow-lg hover:bg-blue-100 transition text-lg">
    Get Started Free
  </a>
</section>

<!-- Features Section -->
<section class="py-16 bg-white text-gray-900">
  <div class="max-w-5xl mx-auto px-4">
    <h2 class="text-3xl md:text-4xl font-bold text-center mb-10">Why Choose SpyWatch?</h2>
    <div class="grid md:grid-cols-2 gap-10">
      <div class="flex items-start space-x-4">
        <span class="text-blue-700 text-3xl">🔍</span>
        <div>
          <h3 class="font-semibold text-xl mb-1">Automated Monitoring</h3>
          <p>Get instant alerts when your competitors update their websites, launch new products, or change prices.</p>
        </div>
      </div>
      <div class="flex items-start space-x-4">
        <span class="text-blue-700 text-3xl">📈</span>
        <div>
          <h3 class="font-semibold text-xl mb-1">Visual Change Tracking</h3>
          <p>See exactly what changed with side-by-side diffs and visual highlights for every update.</p>
        </div>
      </div>
      <div class="flex items-start space-x-4">
        <span class="text-blue-700 text-3xl">⚡</span>
        <div>
          <h3 class="font-semibold text-xl mb-1">Easy Setup</h3>
          <p>Start tracking any website in seconds—no coding or complex configuration required.</p>
        </div>
      </div>
      <div class="flex items-start space-x-4">
        <span class="text-blue-700 text-3xl">🔔</span>
        <div>
          <h3 class="font-semibold text-xl mb-1">Real-Time Notifications</h3>
          <p>Receive email or in-app notifications the moment something changes, so you never miss a beat.</p>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Screenshots Section -->
<section class="py-16 bg-gray-50">
  <div class="max-w-6xl mx-auto px-4">
    <h2 class="text-3xl md:text-4xl font-bold text-center mb-10">See SpyWatch in Action</h2>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      {#each screenshots as src}
        <div class="rounded-xl overflow-hidden shadow-lg bg-white">
          <img src={src} alt="SpyWatch screenshot" class="w-full h-64 object-cover object-top hover:scale-105 transition-transform duration-300" />
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Final Call to Action -->
<section class="py-16 bg-gradient-to-br from-blue-700 to-blue-500 text-white text-center">
  <h2 class="text-3xl md:text-4xl font-bold mb-6">Ready to get started?</h2>
  <a href="/auth/signup" class="bg-white text-blue-800 font-bold px-8 py-3 rounded-full shadow-lg hover:bg-blue-100 transition text-lg">
    Create Your Free Account
  </a>
</section>
