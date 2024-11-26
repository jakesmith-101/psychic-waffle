<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<h1>Welcome to the Site</h1>
{#each data.posts as post}
    <article>
        <header>
            <p><a href="/posts/{post.slug}">{post.postTitle}</a></p>
        </header>
        <p>{post.postDescription}</p>
        <footer>
            <a href="/user/{post.author.userID}">{post.author.nickname}</a>
        </footer>
    </article>
{/each}
