import { redirect } from '@sveltejs/kit';
import { apiFetch } from './api';

export async function login(username: string, password: string) {
    return auth('login', username, password);
}

export async function signup(username: string, password: string) {
    return auth('signup', username, password);
}

async function auth(path: 'signup' | 'login', username: string, password: string) {
    if (username === '') throw new Error('Missing Username');
    if (password === '') throw new Error('Missing Password');

    const data = await apiFetch(`/auth/${path}`, 'POST', { username, password });
    // FIXME: data holds user id and jwt token

    throw redirect(303, `/dashboard`);
}