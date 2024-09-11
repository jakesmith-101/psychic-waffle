/** @type {import('@sveltejs/kit').HandleClientError} */
export async function handleError({ error, event, status, message }) {
    console.error(status, message);

    return {
        message,
        status
    };
}
