import type { Actions } from '@sveltejs/kit';

export const actions: Actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        data.get('username');
        data.get('password');
        data.get('confirmPassword');
    }
};
