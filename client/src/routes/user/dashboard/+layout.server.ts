import { getUser } from '$lib/server/user';
import { redirect, fail } from '@sveltejs/kit';

export async function load({ cookies }: import('./$types.js').LayoutServerLoadEvent) {
    const Token = cookies.get('psychic_waffle_authorisation');
    if (Token === undefined) fail(401, {
        message: 'Not logged in'
    });

    const UserID = cookies.get('psychic_waffle_userid');
    if (UserID !== undefined && UserID !== '') {
        const data = await getUser(UserID);

        return data;
    }
    redirect(302, '/');
}
