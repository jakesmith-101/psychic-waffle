import { apiFetch } from "./api";

interface tAuthReturn {
    message: string,
}

export async function updateUser(token: string, nickname?: string, passwordHash?: string, roleID?: string): Promise<tAuthReturn> {
    if (token === '') throw new Error('Missing Token');

    const data = await apiFetch(`/user/update`, 'POST', { token, nickname, passwordHash, roleID });
    if (typeof data?.message === "string") return data as tAuthReturn
    throw new Error(`Auth failed: ${data?.message}`);
}