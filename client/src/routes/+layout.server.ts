export function load({ cookies }) {
    const Token = cookies.get('psychic_waffle_authorisation');
    const Username = cookies.get('psychic_waffle_username');

    return {
        Token,
        Username,
    };
}