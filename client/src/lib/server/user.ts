import { apiFetch } from "./api";

interface tUpdateReturn {
    message: string,
}

export async function updateUser(token: string, nickname?: string, password?: string, roleID?: string): Promise<tUpdateReturn> {
    if (token === '') throw new Error('Missing Token');

    const data = await apiFetch(`/user/update`, 'POST', { token, nickname, password, roleID });
    if (typeof data?.message === "string") return data as tUpdateReturn
    throw new Error(`Auth failed: ${data?.message}`);
}