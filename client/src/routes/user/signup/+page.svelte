<script lang="ts">
    let showPassword = false;
    $: password = showPassword ? 'text' : 'password';
</script>

<h1>Welcome to signup</h1>

<form method="post">
    <div>
        <label for="username">Username</label>
        <input id="username" name="username" type="text" required placeholder="Username" />
    </div>
    <div>
        <label for="password">Password</label>
        <input id="password" name="password" type={password} required placeholder="Password" />
    </div>
    <div>
        <label for="confirmPassword">Confirm Password</label>
        <input
            id="confirmPassword"
            name="confirmPassword"
            type={password}
            required
            placeholder="Confirm Password"
        />
    </div>
    <div>
        <input type="checkbox" bind:checked={showPassword} id="showPassword" />
        <label for="showPassword"> Show Password </label>
    </div>
    <div>
        <input type="submit" />
    </div>
    <div>
        <a href="/user/login">Login</a>
    </div>
</form>

<style>
    /* Bordered form */
    form {
        border: 3px solid #f1f1f1;
    }

    /* Full-width inputs */
    input[type='text'],
    input[type='password'] {
        width: 100%;
        padding: 12px 20px;
        margin: 8px 0;
        display: inline-block;
        border: 1px solid #ccc;
        box-sizing: border-box;
    }

    /* Set a style for all buttons */
    input[type='submit'] {
        background-color: #04aa6d;
        color: white;
        padding: 14px 20px;
        margin: 8px 0;
        border: none;
        cursor: pointer;
        width: 100%;
    }

    /* Add a hover effect for buttons */
    input[type='submit']:hover {
        opacity: 0.8;
    }

    /* Add padding to containers */
    form > div {
        padding: 16px;
    }
</style>
