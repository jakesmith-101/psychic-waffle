import { type HandleServerError } from "@sveltejs/kit";

export function handleError({ status, message }: Parameters<HandleServerError>[0]) {
    console.error(status, message);

    return {
        message,
        status
    };
}
