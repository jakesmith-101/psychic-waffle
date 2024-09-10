import { apiFetch } from '$lib/server/api.js';

export function load({ cookies }) {
    const token = cookies.get('psychic_waffle_authorisation');

    apiFetch("/user/GetUser", "GET", { token })
        .then(data => {
            console.log(data); // TODO: get nickname
        });
}