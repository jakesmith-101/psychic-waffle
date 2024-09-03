export function login(username: string, password: string) {
	if (username === '') throw new Error('Missing Username');
}

export function signup(username: string, password: string, confirmPassword: string) {
	if (username === '') throw new Error('Missing Username');
	if (password !== confirmPassword) throw new Error('Password does not match');
}
