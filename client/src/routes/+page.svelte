<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
    console.log(data.posts);
</script>

<h1>Welcome to the Site</h1>
<content></content>
{#each data.posts as post}
    <article>
        <header>
            <p>{post.postTitle}</p>
        </header>
        <p>{post.postDescription}</p>
        <footer>
            <a href="/user/{post.author.userID}">{post.author.nickname}</a>
        </footer>
    </article>
{/each}
