import { type tGetPosts, getPosts } from '$lib/server/post';
import { type tGetUser, getUser } from '$lib/server/user';
import { type tGetRole, getRole } from '$lib/server/role';

type tPost = Omit<tGetPosts["posts"][0], "authorID"> & {
    author: Omit<tGetUser, "roleID"> & {
        role: tGetRole
    }
}

export async function load(): Promise<{ posts: tPost[] }> {
    const data = await getPosts(true);
    const posts = await Promise.all(data.posts.map(async post => {
        const user = await getUser(post.authorID);
        const { roleID, ...newUser } = user;
        const role = await getRole(roleID);
        const { authorID, ...newPost } = post;

        return {
            ...newPost,
            author: {
                ...newUser,
                role
            }
        };
    }));

    return {
        posts
    };
}
