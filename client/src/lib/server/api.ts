export const apiVer = `v1`;
export const apiUrl = `http://api:8080`;
export const rootPath = `${apiUrl}/api/${apiVer}`;

export type tMethod = 'GET' | 'POST' | 'PUT' | 'DELETE';
export async function apiFetch(path: `/${string}`, method: tMethod, rawBody?: any): Promise<any> {
    console.log(`${rootPath}${path}`);
    const rInit: RequestInit = {
        method
    };
    if (rawBody !== undefined) {
        const body = JSON.stringify(rawBody);
        console.log(body);
        rInit.body = body;
        rInit.headers = {
            'Content-Type': 'application/json;charset=utf-8'
        }
    }
    const response = await fetch(`${rootPath}${path}`, rInit);

    const res = await response.json();
    if (response.ok) {
        return res;
    } else {
        throw Error(JSON.stringify(res));
    }
}
