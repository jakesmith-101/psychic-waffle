export const apiVer = `v1`;
export const apiUrl = `http://api:8080`;
export const rootPath = `${apiUrl}/api/${apiVer}`;

export type tMethod = 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'OPTIONS' | 'TRACE' | 'PATCH';
type tAuth = ["psychic-waffle-token" | "psychic-waffle-username", string][]
export async function apiFetch<T = any>(path: `/${string}`, method: tMethod, rawBody?: any): Promise<[tAuth | undefined, (T & { message?: string }) | undefined]> {
    console.log(`${rootPath}${path}`);
    const rInit: RequestInit = {
        method
    };
    if (rawBody !== undefined) {
        const body = JSON.stringify(rawBody);
        rInit.body = body;
        rInit.headers = {
            'Content-Type': 'application/json;charset=utf-8'
        };
    }
    const response = await fetch(`${rootPath}${path}`, rInit);

    const res = await response.json();
    if (response.ok) {
        const cookies = response.headers.getSetCookie(); // ["name1=value1", "name2=value2"]
        console.log(cookies);
        let authCookies: tAuth | undefined = undefined;
        if (cookies.length === 2)
            authCookies = cookies.map(c => c.split("=") as tAuth[0]);
        return [authCookies, res];
    } else {
        throw Error(JSON.stringify(res));
    }
}