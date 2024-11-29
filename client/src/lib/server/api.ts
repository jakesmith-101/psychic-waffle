export const apiVer = `v1`;
export const apiUrl = `http://api:8080`;
export const rootPath = `${apiUrl}/api/${apiVer}`;

export type tMethod = 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'OPTIONS' | 'TRACE' | 'PATCH';
export interface tAuth {
    name: string;
    value: string;
}
export async function apiFetch<T = any>(path: `/${string}`, method: tMethod, rawBody?: any): Promise<[tAuth[], (T & { message?: string }) | undefined]> {
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
        const cookies = response.headers
            .getSetCookie()
            .map(c => {
                const matches: tAuth[] = [];
                for (const match of c.matchAll(/{(?<name>\S+)\s+(?<value>\S+)\s+(true|false)\s+[0-9-]+\s+[0-9:]+\s+[+-0-9]+\s+[A-Z]+\s+[0-9]+\s+(true|false)\s+(true|false)\s+[0-9]+\s+(true|false)\s+\[.*?\]}/g)) {
                    const name = match?.groups?.["name"];
                    const value = match?.groups?.["value"];
                    if (name !== undefined && value !== undefined)
                        matches.push({ name, value });
                }
                return matches
            }).flat();
        return [cookies, res];
    } else {
        throw Error(JSON.stringify(res));
    }
}