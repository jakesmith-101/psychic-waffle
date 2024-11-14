import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
    default: async ({ cookies }) => {
        cookies.delete('psychic_waffle_authorisation', { path: '/' });
        cookies.delete('psychic_waffle_username', { path: '/' });
        redirect(302, `/`);
    }
};
