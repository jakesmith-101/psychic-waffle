<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<h1>Welcome {data.username}!</h1>
<slot></slot>
