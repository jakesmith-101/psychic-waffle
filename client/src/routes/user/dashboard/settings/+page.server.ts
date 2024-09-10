import * as api from '$lib/server/user.js';
import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
    default: async ({ cookies, request }) => {
        const formData = await request.formData();
        const nickname = formData.get('nickname');
        const password = formData.get('password');

        const token = cookies.get("psychic_waffle_authorisation");
        const userID = cookies.get("psychic_waffle_userid");
        if (token === undefined)
            return fail(400, {
                message: "Not logged in"
            })
        if (userID === undefined)
            return fail(400, {
                message: "Missing User ID"
            })

        if (
            typeof nickname !== "string" ||
            nickname.length < 3 ||
            nickname.length > 31 ||
            !/^[a-z0-9_-]+$/.test(nickname)
        ) {
            return fail(400, {
                message: "Invalid nickname"
            });
        }
        if (typeof password !== "string" || password.length < 6 || password.length > 255) {
            return fail(400, {
                message: "Invalid password"
            });
        }

        const body = await api.updateUser(userID, token, nickname, password);
        console.log(body.message);
    }
};
