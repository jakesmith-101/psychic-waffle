import * as api from '$lib/server/test.js';
import type { Actions } from './$types';

export const actions: Actions = {
    default: async ({ request }) => {
        const formData = await request.formData();
        const name = formData.get('name')?.toString();
        if (name !== undefined) await api.healthcheck(name);
    }
};
