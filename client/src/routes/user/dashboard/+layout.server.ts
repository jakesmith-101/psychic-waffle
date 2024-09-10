import { redirect } from '@sveltejs/kit';

export function load({ cookies }) {
    const Token = cookies.get('psychic_waffle_authorisation');
    if (Token === undefined) throw redirect(303, '/');
}