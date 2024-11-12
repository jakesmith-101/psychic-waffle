import { type HandleClientError } from "@sveltejs/kit";

export async function handleError({ error, event, status, message }: Parameters<HandleClientError>[0]) {
    console.error(status, message);

    return {
        message,
        status
    };
}
