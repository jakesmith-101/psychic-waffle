import { getUser } from '$lib/server/user.js';

export async function load({ params }) {
    const data = await getUser(params.userID);

    return data;
}
