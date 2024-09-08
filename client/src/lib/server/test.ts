import { redirect } from '@sveltejs/kit';
import { apiFetch } from './api';

export async function healthcheck(name: string) {
    const data = await apiFetch(`/healthcheck/${name}`, 'GET', {});
    throw redirect(303, `/dashboard`);
}