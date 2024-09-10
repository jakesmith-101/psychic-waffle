<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<a href="/user/dashboard/settings">User Settings</a>
<a href={`/user/${data.UserID}`}>User Details</a>
