<script lang="ts">
    import type { PageData } from './$types';

    export let data: PageData;
</script>

<div>
    <div>
        <label for="username">Username</label>
        <input
            id="username"
            name="username"
            type="text"
            placeholder="Username"
            value={data.username}
            disabled
        />
    </div>
    <div>
        <label for="nickname">Nickname</label>
        <input
            id="nickname"
            name="nickname"
            type="text"
            placeholder="Nickname"
            value={data.nickname}
            disabled
        />
    </div>
    <div>
        <label for="createdAt">Created</label>
        <input
            id="createdAt"
            name="createdAt"
            type="text"
            placeholder="Created"
            value={data.createdAt}
            disabled
        />
    </div>
    <div>
        <label for="updatedAt">Updated</label>
        <input
            id="updatedAt"
            name="updatedAt"
            type="text"
            placeholder="Updated"
            value={data.updatedAt}
            disabled
        />
    </div>
</div>

<style>
    input[type='text'] {
        width: 100%;
        padding: 12px 20px;
        margin: 8px 0;
        display: inline-block;
        border: 1px solid #ccc;
        box-sizing: border-box;
    }

    /* Add padding to containers */
    div > div {
        padding: 16px;
    }
</style>
