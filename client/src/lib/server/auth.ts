import { apiFetch } from './api';

export async function login(username: string, password: string) {
    return auth('login', username, password);
}

export async function signup(username: string, password: string) {
    return auth('signup', username, password);
}

interface tAuthReturn {
    token: string;
    userID: string;
    message: string;
}

async function auth(
    path: 'signup' | 'login',
    username: string,
    password: string
): Promise<tAuthReturn> {
    if (username === '') throw new Error('Missing Username');
    if (password === '') throw new Error('Missing Password');

    const data = await apiFetch(`/auth/${path}`, 'POST', { username, password });
    if (
        typeof data?.token === 'string' &&
        typeof data?.userID === 'string' &&
        typeof data?.message === 'string'
    )
        return data as tAuthReturn;
    throw new Error(`Auth failed: ${data?.message}`);
}
