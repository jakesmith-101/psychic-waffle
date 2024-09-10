<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<div>
    <p>{data.nickname}</p>
    <p>{data.username}</p>
    <p>{data.createdAt}</p>
    <p>{data.updatedAt}</p>
</div>
