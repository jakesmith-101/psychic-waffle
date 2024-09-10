import { redirect } from "@sveltejs/kit";

export function load() {
    throw redirect(303, `/unexpected`); // currently no expected errors
}