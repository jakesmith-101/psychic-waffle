import { type tGetPosts, getPosts } from '$lib/server/post';
import { type tGetUser, getUser } from '$lib/server/user';
import { type tGetRole, getRole } from '$lib/server/role';

type tPost = Omit<tGetPosts["posts"][0], "authorID"> & {
    author: Omit<tGetUser, "roleID"> & {
        role: tGetRole
    }
}
type tCache<T> = { [k: string]: T };

export async function load(): Promise<{ posts: tPost[] }> {
    const roles: tCache<tGetRole> = {};
    const users: tCache<tGetUser> = {};
    const data = await getPosts(true);
    const posts = await Promise.all(data.posts.map(async post => {
        // attempting to cache each user to prevent repeated api calls for the same info
        let user: tGetUser | undefined = undefined;
        if (post.authorID in Object.entries(users))
            user = users?.[post.authorID];
        else {
            user = await getUser(post.authorID);
            users[post.authorID] = user;
        }
        const { roleID, ...newUser } = user;

        // attempting to cache each role to prevent repeated api calls for the same info
        let role: tGetRole | undefined = undefined;
        if (roleID in Object.entries(roles))
            role = roles?.[roleID];
        else {
            role = await getRole(roleID);
            roles[roleID] = role;
        }
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
