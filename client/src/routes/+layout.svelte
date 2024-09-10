<script>
    import { page } from '$app/stores';
</script>

<nav>
    <ul>
        <li><a href="/" class={$page.url.pathname === '/' ? 'active' : 'inactive'}>Home</a></li>
        <li>
            <a href="/login" class={$page.url.pathname === '/login' ? 'active' : 'inactive'}
                >Login</a
            >
        </li>
        <li>
            <a href="/signup" class={$page.url.pathname === '/signup' ? 'active' : 'inactive'}
                >Signup</a
            >
        </li>

        <li style="float:right">
            <a href="/dashboard" class={$page.url.pathname === '/dashboard' ? 'active' : 'inactive'}
                >Dashboard</a
            >
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

    .active {
        background-color: #04aa6d;
    }
</style>
