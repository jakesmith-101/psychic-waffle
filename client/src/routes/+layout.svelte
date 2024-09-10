<script lang="ts">
    import { page } from '$app/stores';
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<nav>
    <ul>
        <li><a href="/" class={$page.url.pathname === '/' ? 'active' : 'inactive'}>Home</a></li>
        {#if data.Token !== undefined}
            <li>
                <a
                    href="/user/dashboard"
                    class={$page.url.pathname === '/dashboard' ? 'active' : 'inactive'}>Dashboard</a
                >
            </li>
        {/if}

        <li style="float:right">
            {#if data.Token !== undefined}
                <a
                    href="/user/logout"
                    class={$page.url.pathname === '/logout' ? 'active' : 'inactive'}>Logout</a
                >
            {:else}
                <a
                    href="/user/login"
                    class={$page.url.pathname === '/login' ? 'active' : 'inactive'}>Login</a
                >
            {/if}
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
