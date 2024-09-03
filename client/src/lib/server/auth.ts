import { redirect } from '@sveltejs/kit';
import { apiFetch } from './api';

export async function login(username: string, password: string) {
	if (username === '') throw new Error('Missing Username');
	if (password === '') throw new Error('Missing Password');

	const data = await apiFetch('/auth/login', 'POST', { username, password });

	throw redirect(303, `/dashboard`);
}

export async function signup(username: string, password: string, confirmPassword: string) {
	if (username === '') throw new Error('Missing Username');
	if (password === '') throw new Error('Missing Password');
	if (password !== confirmPassword) throw new Error('Password does not match');

	// TODO: password strength
	const data = await apiFetch('/auth/signup', 'POST', { username, password, confirmPassword });

	throw redirect(303, `/dashboard`);
}
