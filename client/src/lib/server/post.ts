import { apiFetch } from './api';

export interface tGetPosts {
    posts: {
        postID: string;
        postTitle: string;
        postDescription: string;
        votes: number;
        authorID: string;
        createdAt: string;
        updatedAt: string;
    }[]
}

export async function getPosts(type: "latest" | "popular"): Promise<tGetPosts> {
    const data = await apiFetch<tGetPosts>(`/posts/${type}`, 'GET'); // possible API error response message
    if (data?.posts !== undefined)
        return data as tGetPosts;
    throw new Error(`Get user failed: ${data?.message}`);
}