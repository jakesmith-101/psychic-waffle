export function handleError({ status, message }: Parameters<import("@sveltejs/kit").HandleServerError>[0]) {
    console.error(status, message);

    return {
        message,
        status
    };
}
