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
</form>
