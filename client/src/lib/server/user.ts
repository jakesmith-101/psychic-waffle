import { apiFetch } from "./api";

interface tUpdateReturn {
    message: string,
}

export async function updateUser(token: string, nickname?: string, password?: string, roleID?: string): Promise<tUpdateReturn> {
    if (token === '') throw new Error('Missing Token');

    const data = await apiFetch(`/user/update`, 'POST', { token, nickname, password, roleID });
    if (typeof data?.message === "string") return data as tUpdateReturn;
    throw new Error(`Update user failed: ${data?.message}`);
}

interface tGetReturn {
    userID: string
    username: string
    nickname: string
    roleID: string
    updatedAt: Date
    createdAt: Date
}

export async function getUser(userID: string): Promise<tGetReturn> {
    if (userID === '') throw new Error('Missing user ID');

    const data = await apiFetch(`/user/get`, 'GET', { userID });
    if (
        typeof data?.username === "string" &&
        typeof data?.nickname === "string" &&
        typeof data?.roleID === "string" &&
        typeof data?.userID === "string" &&
        data?.updatedAt instanceof Date &&
        data?.createdAt instanceof Date
    )
        return data as tGetReturn;
    throw new Error(`Get user failed: ${data?.message}`);
}