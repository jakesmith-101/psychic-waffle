<script lang="ts">
	let showPassword = false;
	$: password = showPassword ? 'text' : 'password';
</script>

<h1>Welcome to login</h1>

<form method="post">
	<div>
		<label for="username" class="sr-only">Username</label>
		<input id="username" name="username" type="text" required placeholder="Username" />
	</div>
	<div>
		<label for="password" class="sr-only">Password</label>
		<input id="password" name="password" type={password} required placeholder="Password" />
	</div>
</form>
