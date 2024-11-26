import { apiFetch } from './api';

export interface tGetRole {
    roleID: string;
    name: string;
    permissions: number;
}

export async function getRole(roleID: string): Promise<tGetRole> {
    if (roleID === '') throw new Error('Missing role ID');

    const [, data] = await apiFetch<tGetRole>(`/roles/${roleID}`, 'GET'); // possible API error response message
    if (
        typeof data?.roleID === 'string' &&
        typeof data?.name === 'string' &&
        typeof data?.permissions === 'number'
    )
        return data as tGetRole;
    console.log(data);
    throw new Error(`Get role failed: ${data?.message}`);
}