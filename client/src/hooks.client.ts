export async function handleError({ error, event, status, message }: Parameters<import("@sveltejs/kit").HandleClientError>[0]) {
    console.error(status, message);

    return {
        message,
        status
    };
}
