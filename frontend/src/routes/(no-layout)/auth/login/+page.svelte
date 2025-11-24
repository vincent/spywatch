<script lang="ts">
  import MetaTag from '$lib/components/MetaTag.svelte';
	import SignIn from '$lib/components/SignIn.svelte';
  import { Label, Input } from 'flowbite-svelte';
	import { client } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

  let title = 'Sign in to platform';
  let site = {
    link: '/',
    name: 'SpyWatch',
    img: '/favicon.svg',
    imgAlt: 'SpyWatch Logo',
  };
  let rememberMe = true;
  let lostPassword = true;
  let createAccount = true;
  let lostPasswordLink = 'forgot-password';
  let loginTitle = 'Login to your account';
  let registerLink = '/auth/signup';
  let createAccountTitle = 'Create account';

  const onSubmit = async (e: Event) => {
    e.preventDefault();
    const formData = new FormData(e.target as HTMLFormElement);
    const email = formData.get('email')?.toString() || ''
    const password = formData.get('password')?.toString() || ''

    await client.collection('users').authWithPassword(email, password);
    goto('/releases');
  };

  const path: string = '/authentication/sign-in';
  const description: string = 'SpyWatch - Sign in';
  const metaTitle: string = 'SpyWatch - Sign in';
  const subtitle: string = 'Sign in';
</script>

<MetaTag {path} {description} title={metaTitle} {subtitle} />

<div class="overflow-hidden lg:flex bg-white dark:bg-gray-800">
  <SignIn {title} {site} {rememberMe} {lostPassword} {createAccount} {lostPasswordLink} {loginTitle} {registerLink} {createAccountTitle} onsubmit={onSubmit}>
    <div>
      <Label for="email" class="mb-2 dark:text-white">Your email</Label>
      <Input type="email" name="email" id="email" placeholder="name@company.com" required class="border outline-none dark:border-gray-600 dark:bg-gray-700" />
    </div>
    <div>
      <Label for="password" class="mb-2 dark:text-white">Your password</Label>
      <Input type="password" name="password" id="password" placeholder="••••••••" required class="border outline-none dark:border-gray-600 dark:bg-gray-700" />
    </div>
  </SignIn>
</div>