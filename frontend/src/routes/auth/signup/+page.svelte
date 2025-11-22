<script lang="ts">
  import { button_class, input_class, shadow_hover } from '$lib/styles.svelte';

  let name = "";
  let email = "";
  let password = "";
  let confirmPassword = "";
  let acceptTerms = false;

  let error = "";
  let success = "";

  function validateEmail(email: string) {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  }

  function validatePassword(password: string) {
    // At least 8 chars, one letter, one number
    return /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(password);
  }

  function handleSignup(event: Event) {
    event.preventDefault();
    error = "";
    success = "";

    if (!name.trim()) {
      error = "Full name is required.";
      return;
    }
    if (!validateEmail(email)) {
      error = "Please enter a valid email address.";
      return;
    }
    if (!validatePassword(password)) {
      error = "Password must be at least 8 characters and contain a letter and a number.";
      return;
    }
    if (password !== confirmPassword) {
      error = "Passwords do not match.";
      return;
    }
    if (!acceptTerms) {
      error = "You must accept the terms and privacy policy.";
      return;
    }

    // TODO: Implement actual signup logic (API call)
    success = "Signup successful! (Not yet implemented)";
  }
</script>

<main class="flex items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900">
  <div class="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-6 text-center text-indigo-700 dark:text-indigo-300">Create your account</h2>
    <form on:submit|preventDefault={handleSignup} class="space-y-4">
      {#if error}
        <div class="bg-red-100 text-red-700 px-4 py-2 rounded">{error}</div>
      {/if}
      {#if success}
        <div class="bg-green-100 text-green-700 px-4 py-2 rounded">{success}</div>
      {/if}
      <label class="block">
        <span class="text-gray-700 dark:text-gray-200">Full Name</span>
        <input
          type="text"
          bind:value={name}
          class={input_class + " w-full"}
          placeholder="Your full name"
          required
        />
      </label>
      <label class="block">
        <span class="text-gray-700 dark:text-gray-200">Email</span>
        <input
          type="email"
          bind:value={email}
          class={input_class + " w-full"}
          placeholder="you@example.com"
          required
        />
      </label>
      <label class="block">
        <span class="text-gray-700 dark:text-gray-200">Password</span>
        <input
          type="password"
          bind:value={password}
          class={input_class + " w-full"}
          placeholder="Password"
          required
        />
        <span class="text-xs text-gray-400">At least 8 characters, 1 letter, 1 number</span>
      </label>
      <label class="block">
        <span class="text-gray-700 dark:text-gray-200">Confirm Password</span>
        <input
          type="password"
          bind:value={confirmPassword}
          class={input_class + " w-full"}
          placeholder="Confirm password"
          required
        />
      </label>
      <label class="flex items-center space-x-2">
        <input
          type="checkbox"
          bind:checked={acceptTerms}
          required
        />
        <span class="text-sm text-gray-600 dark:text-gray-300">
          I accept the <a href="/terms" class="underline text-indigo-600 dark:text-indigo-300">Terms</a> and <a href="/privacy" class="underline text-indigo-600 dark:text-indigo-300">Privacy Policy</a>
        </span>
      </label>
      <button type="submit" class={button_class + " w-full font-semibold mt-2 " + shadow_hover}>
        Sign Up
      </button>
    </form>
    <div class="mt-6 text-center text-sm text-gray-600 dark:text-gray-300">
      Already have an account?
      <a href="/auth/login" class="text-indigo-600 dark:text-indigo-300 underline ml-1">Login</a>
    </div>
    <div class="mt-6">
      <div class="flex items-center justify-center space-x-2 mb-2">
        <span class="text-gray-400 text-xs">or sign up with</span>
      </div>
      <div class="flex justify-center space-x-4">
        <button class="bg-gray-100 dark:bg-gray-700 p-2 rounded shadow hover:shadow-md" disabled>
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><!-- Google icon --></svg>
        </button>
        <button class="bg-gray-100 dark:bg-gray-700 p-2 rounded shadow hover:shadow-md" disabled>
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><!-- GitHub icon --></svg>
        </button>
      </div>
    </div>
  </div>
</main>
