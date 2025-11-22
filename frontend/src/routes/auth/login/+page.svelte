<script lang="ts">
  import { goto } from "$app/navigation";
  import { client } from "$lib/pocketbase";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let error: string | null = null;
  let loading = false;

  async function handleLogin(event: Event) {
    event.preventDefault();
    error = null;
    loading = true;
    try {
      await client.collection('users').authWithPassword(email, password);
      goto('/dashboard');
    } catch (err: any) {
      error = err?.message || "Login failed. Please check your credentials.";
    } finally {
      loading = false;
    }
  }
</script>

<style>
  main {
    max-width: 400px;
    margin: 3rem auto;
    padding: 2rem;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 16px rgba(0,0,0,0.08);
    font-family: system-ui, sans-serif;
  }
  h2 {
    text-align: center;
    margin-bottom: 1.5rem;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  label {
    display: flex;
    flex-direction: column;
    font-weight: 500;
  }
  input[type="email"],
  input[type="password"] {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
    margin-top: 0.25rem;
    font-size: 1rem;
  }
  button[type="submit"] {
    padding: 0.75rem;
    background: #2d72d9;
    color: #fff;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.2s;
  }
  button[type="submit"]:hover,
  button[type="submit"]:focus {
    background: #1a4e96;
  }
  .links {
    display: flex;
    justify-content: space-between;
    margin-top: 1rem;
    font-size: 0.95rem;
  }
  .error {
    color: #d32f2f;
    background: #ffeaea;
    border: 1px solid #f5c2c7;
    padding: 0.75rem;
    border-radius: 4px;
    margin-bottom: 1rem;
    text-align: center;
  }
  @media (max-width: 500px) {
    main {
      margin: 1rem;
      padding: 1rem;
    }
  }
</style>

<main>
  <h2>Login</h2>
  {#if error}
    <div class="error" role="alert">{error}</div>
  {/if}
  <form on:submit|preventDefault={handleLogin} aria-describedby={error ? "login-error" : undefined}>
    <label for="email">
      Email
      <input
        id="email"
        type="email"
        bind:value={email}
        required
        autocomplete="username"
        placeholder="you@example.com"
      />
    </label>
    <label for="password">
      Password
      <input
        id="password"
        type="password"
        bind:value={password}
        required
        autocomplete="current-password"
        placeholder="Your password"
      />
    </label>
    <button type="submit" disabled={loading}>
      {#if loading}Logging in...{/if}
      {#if !loading}Login{/if}
    </button>
  </form>
  <div class="links">
    <a href="/auth/signup">Sign up</a>
    <a href="/forgot-password">Forgot password?</a>
  </div>
</main>
