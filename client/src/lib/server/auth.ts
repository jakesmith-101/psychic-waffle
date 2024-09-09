import { apiFetch } from './api';

export async function login(username: string, password: string) {
    return auth('login', username, password);
}

export async function signup(username: string, password: string) {
    return auth('signup', username, password);
}

interface tAuthReturn {
    Token: string,
    UserID: string,
    Message: string,
}

async function auth(path: 'signup' | 'login', Username: string, Password: string): Promise<tAuthReturn> {
    if (Username === '') throw new Error('Missing Username');
    if (Password === '') throw new Error('Missing Password');

    const data = await apiFetch(`/auth/${path}`, 'POST', { Username, Password });
    if (typeof data === "object") {
        if (
            typeof data?.token === "string" &&
            typeof data?.UserID === "string" &&
            typeof data?.Message === "string"
        )
            return data as tAuthReturn
    }
    throw new Error(`Auth failed: ${data?.message}`);
}