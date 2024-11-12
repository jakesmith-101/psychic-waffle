import { getUser } from '$lib/server/user.js';

export async function load({ params }: import("./$types.js").PageServerLoadEvent) {
    const data = await getUser(params.userID);

    return data;
}
