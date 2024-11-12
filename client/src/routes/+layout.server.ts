import { type LayoutServerLoadEvent } from "./$types";

export function load({ cookies }: LayoutServerLoadEvent) {
    const Token = cookies.get('psychic_waffle_authorisation');
    const UserID = cookies.get('psychic_waffle_userid');

    return {
        Token,
        UserID
    };
}
