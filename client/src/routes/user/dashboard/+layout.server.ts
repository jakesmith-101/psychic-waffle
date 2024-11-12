import { getUser } from '$lib/server/user.js';
import { redirect } from '@sveltejs/kit';
import { type LayoutServerLoadEvent } from './$types';

export async function load({ cookies }: LayoutServerLoadEvent) {
    const Token = cookies.get('psychic_waffle_authorisation');
    if (Token === undefined) throw redirect(303, '/');

    const UserID = cookies.get('psychic_waffle_userid');
    if (UserID !== undefined && UserID !== '') {
        const data = await getUser(UserID);

        return data;
    }
    throw redirect(303, '/');
}
