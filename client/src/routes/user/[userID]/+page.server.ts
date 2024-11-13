import { getUser } from '$lib/server/user';

export async function load({ params }: import('./$types.js').PageServerLoadEvent) {
    const data = await getUser(params.userID);

    return data;
}
