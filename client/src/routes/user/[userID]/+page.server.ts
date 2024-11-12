import { getUser } from '$lib/server/user.js';
import { type PageServerLoadEvent } from './$types';

export async function load({ params }: PageServerLoadEvent) {
    const data = await getUser(params.userID);

    return data;
}
