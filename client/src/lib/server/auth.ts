import { apiFetch } from './api';

export async function login(username: string, password: string) {
    return auth('login', username, password);
}

export async function signup(username: string, password: string) {
    return auth('signup', username, password);
}

interface tPostAuth {
    message: string;
}

async function auth(
    path: 'signup' | 'login',
    username: string,
    password: string
): Promise<tPostAuth> {
    if (username === '') throw new Error('Missing Username');
    if (password === '') throw new Error('Missing Password');

    const data = await apiFetch<tPostAuth>(`/auth/${path}`, 'POST', { username, password }); // possible API error response message
    if (
        typeof data?.message === 'string'
    )
        return data as tPostAuth;
    throw new Error(`Auth failed: ${data?.message}`);
}
