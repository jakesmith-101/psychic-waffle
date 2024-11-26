import { apiFetch } from './api';

interface tUpdateUser {
    message: string;
}

export async function updateUser(
    token: string,
    nickname?: string,
    password?: string,
    roleID?: string
): Promise<tUpdateUser> {
    if (token === '') throw new Error('Missing Token');

    const [, data] = await apiFetch<tUpdateUser>(`/users/update`, 'POST', { token, nickname, password, roleID });
    if (typeof data?.message === 'string')
        return data as tUpdateUser;
    console.log(data);
    throw new Error(`Update user failed: ${data?.message}`);
}

export interface tGetUser {
    userID: string;
    username: string;
    nickname: string;
    roleID: string;
    updatedAt: string;
    createdAt: string;
}

export async function getUser(userID: string): Promise<tGetUser> {
    if (userID === '') throw new Error('Missing user ID');

    const [, data] = await apiFetch<tGetUser>(`/users/${userID}`, 'GET'); // possible API error response message
    if (
        typeof data?.username === 'string' &&
        typeof data?.nickname === 'string' &&
        typeof data?.roleID === 'string' &&
        typeof data?.userID === 'string' &&
        typeof data?.updatedAt === 'string' &&
        typeof data?.createdAt === 'string'
    )
        return data as tGetUser;
    throw new Error(`Get user failed: ${data?.message}`);
}
