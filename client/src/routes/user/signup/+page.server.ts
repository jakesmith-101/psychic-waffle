import * as api from '$lib/server/auth.js';
import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';

export const actions: Actions = {
    default: async ({ cookies, request }) => {
        const formData = await request.formData();
        const username = formData.get('username');
        const password = formData.get('password');
        const confirmPassword = formData.get('confirmPassword');

        if (
            typeof username !== "string" ||
            username.length < 3 ||
            username.length > 31 ||
            !/^[A-Za-z0-9_-]+$/.test(username)
        ) {
            return fail(400, {
                message: "Invalid username"
            });
        }
        if (typeof password !== "string" || password.length < 6 || password.length > 255) {
            return fail(400, {
                message: "Invalid password"
            });
        }

        if (username !== undefined && password !== undefined && password === confirmPassword) {
            const loginInfo = await api.signup(username, password);
            cookies.set("psychic_waffle_authorisation", loginInfo.token, { path: '/' });
            cookies.set("psychic_waffle_username", username, { path: '/' });
            throw redirect(303, `/user/dashboard`);
        }
    }
};
