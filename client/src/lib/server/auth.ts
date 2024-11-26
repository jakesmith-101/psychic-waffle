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
): Promise<tPostAuth & { cookies: [string, string][] | undefined }> {
    if (username === '') throw new Error('Missing Username');
    if (password === '') throw new Error('Missing Password');

    const [headers, data] = await apiFetch<tPostAuth>(`/auth/${path}`, 'POST', { username, password }); // possible API error response message
    const cookies = headers.getSetCookie(); // ["name1=value1", "name2=value2"]
    let authCookies: [string, string][] | undefined = undefined;
    if (cookies.length === 2)
        authCookies = cookies.map(c => c.split("=") as [string, string]);
    if (
        typeof data?.message === 'string'
    )
        return {
            ...data as tPostAuth,
            cookies: authCookies
        };
    console.log(data);
    throw new Error(`Auth failed: ${data?.message}`);
}
