/** @type {import('@sveltejs/kit').HandleServerError} */
export function handleError({ status, message }) {
    console.error(status, message);

    return {
        message,
        status
    };
}
