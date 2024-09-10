<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<ul>
    <li><a href="/user/dashboard/settings">User Settings</a></li>
    <li><a href={`/user/${data.UserID}`}>User Details</a></li>
</ul>

<style>
    ul {
        list-style-type: none;
        margin: 0;
        padding: 0;
        width: 200px;
        background-color: #f1f1f1;
    }

    li a {
        display: block;
        color: #000;
        padding: 8px 16px;
        text-decoration: none;
    }

    /* Change the link color on hover */
    li a:hover {
        background-color: #555;
        color: white;
    }
</style>
