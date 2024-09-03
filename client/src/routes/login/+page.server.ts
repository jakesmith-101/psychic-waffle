import * as api from '$lib/server/api.js';
import type { Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ cookies, request }) => {
		const data = await request.formData();
		const username = data.get('username')?.toString();
		const password = data.get('password')?.toString();
		if (username !== undefined && password !== undefined) api.login(username, password);
	}
};
