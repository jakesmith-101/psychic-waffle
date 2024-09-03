import * as api from '$lib/server/auth.js';
import type { Actions } from './$types';

export const actions: Actions = {
	default: async ({ request }) => {
		const formData = await request.formData();
		const username = formData.get('username')?.toString();
		const password = formData.get('password')?.toString();
		if (username !== undefined && password !== undefined) await api.login(username, password);
	}
};
