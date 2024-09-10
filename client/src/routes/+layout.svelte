<script>
    import { page } from '$app/stores';
</script>

<nav>
    <ul>
        <li><a href="/" aria-current={$page.url.pathname === '/'}>Home</a></li>
        <li><a href="/login" aria-current={$page.url.pathname === '/login'}>Login</a></li>
        <li><a href="/signup" aria-current={$page.url.pathname === '/signup'}>Signup</a></li>
        <li>
            <a href="/dashboard" aria-current={$page.url.pathname === '/dashboard'}>Dashboard</a>
        </li>
    </ul>
</nav>

<slot></slot>

<style>
    ul {
        list-style-type: none;
        margin: 0;
        padding: 0;
        overflow: hidden;
        background-color: #333;
    }

    li {
        float: left;
    }

    li a {
        display: block;
        color: white;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
    }

    li a:hover {
        background-color: #111;
    }
</style>
