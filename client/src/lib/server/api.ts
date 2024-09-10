export const apiVer = `v1`;
export const apiUrl = `api:8080`;
export const rootPath = `${apiUrl}/api/${apiVer}`;

export type tMethod = 'GET' | 'POST' | 'PUT' | 'DELETE';
export async function apiFetch(path: `/${string}`, method: tMethod, body: any): Promise<any> {
    console.log(`${rootPath}${path}`);
    const response = await fetch(`${rootPath}${path}`, {
        method,
        body: JSON.stringify(body)
    });

    const res = await response.json();
    if (response.ok) {
        return res;
    } else {
        throw Error(await response.text());
    }
}
