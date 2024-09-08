import { redirect } from '@sveltejs/kit';
import { apiFetch } from './api';

export async function login(username: string, passwordHash: string) {
    if (username === '') throw new Error('Missing Username');
    if (passwordHash === '') throw new Error('Missing Password');

    const data = await apiFetch('/auth/login', 'POST', { username, passwordHash });
    // FIXME: data holds user id and jwt token

    throw redirect(303, `/dashboard`);
}

export async function signup(username: string, passwordHash: string) {
    if (username === '') throw new Error('Missing Username');
    if (passwordHash === '') throw new Error('Missing Password');

    const data = await apiFetch('/auth/signup', 'POST', { username, passwordHash });
    // FIXME: data holds user id and jwt token

    throw redirect(303, `/dashboard`);
}
