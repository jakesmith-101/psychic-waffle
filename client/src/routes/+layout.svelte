<script>
	import { page } from '$app/stores';
</script>

<nav>
	<a href="/" aria-current={$page.url.pathname === '/'}>home</a>
	<a href="/login" aria-current={$page.url.pathname === '/login'}>login</a>
	<a href="/signup" aria-current={$page.url.pathname === '/signup'}>signup</a>
</nav>

<slot></slot>
