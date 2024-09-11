import { getUser } from '$lib/server/user.js';
import { redirect } from '@sveltejs/kit';

export async function load({ cookies }) {
    const Token = cookies.get('psychic_waffle_authorisation');
    if (Token === undefined) throw redirect(303, '/');

    const UserID = cookies.get('psychic_waffle_userid');
    if (UserID !== undefined && UserID !== '') {
        const data = await getUser(UserID);

        return data;
    }
    throw redirect(303, '/');
}
