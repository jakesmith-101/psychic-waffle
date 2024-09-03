import type { Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ cookies, request }) => {
		const data = await request.formData();
		//db.createTodo(cookies.get('userid'), data.get('description'));
	}
};
