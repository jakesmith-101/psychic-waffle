export const apiVer = `v1`;
export const apiUrl = `http://api:8080`;
export const rootPath = `${apiUrl}/api/${apiVer}`;

export type tMethod = 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'OPTIONS' | 'TRACE' | 'PATCH';
export async function apiFetch<T = any>(path: `/${string}`, method: tMethod, rawBody?: any): Promise<[Headers, (T & { message?: string }) | undefined]> {
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
        console.log(res?.message);
        return response.headers, res;
    } else {
        throw Error(JSON.stringify(res));
    }
}