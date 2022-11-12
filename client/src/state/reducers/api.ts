import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import {
    NewProduct,
    Product,
    ProductUser,
    User,
    Post,
    UserInfo,
} from '../interfaces'
import { RootState } from '../store'
import { ProductPostsResponse } from '../interfaces'
import { number } from 'prop-types'
import Posts from '../../components/Products/Home/Posts/Posts'
interface ProductSearch {
    name?: string
    field?: string
    page: number
}
export interface ProductResponse {
    page: number
    products: Product[]
}
interface OneProductResponse {
    product: Product
}

interface CreatePostArgs {
    post: any
    productName: string
    channel: string
}
const api = createApi({
    reducerPath: 'api',
    baseQuery: fetchBaseQuery({
        baseUrl: 'http://localhost:8000',
        prepareHeaders: (headers) => {
            const token = localStorage.getItem('token')
            if (token) {
                headers.set('authorization', `Bearer ${token}`)
            }
            return headers
        },
    }),
    tagTypes: [
        'Posts',
        'User',
        'Product',
        'Comments',
        'Stats',
        'Board',
        'Post',
    ],
    endpoints: (builder) => ({
        getProducts: builder.query<
            ProductResponse | Product | ProductPostsResponse,
            ProductSearch
        >({
            query: (search) =>
                `products/search/${search.name || -1}/${search.field || -1}/${
                    search.page
                }`,
            providesTags: ['Posts', 'User'],
        }),
        getLatestPosts: builder.query<
            Post[],
            { field: string; lastId?: number }
        >({
            query: ({ field, lastId }) =>
                `products/latest_posts/${field}/${lastId || 0}`,
            providesTags: ['Posts', 'User'],
        }),
        getUserInfo: builder.query<UserInfo, void>({
            query: () => 'users/user/get',
            providesTags: ['User'],
        }),
        getTopProducts: builder.query<Product[] | null, void>({
            query: () => 'products/today_top_products',
            providesTags: ['Product'],
        }),
        signIn: builder.mutation<
            { token: string },
            { username: string; password: string }
        >({
            query: (userInfo) => ({
                url: 'users/login',
                method: 'POST',
                body: userInfo,
            }),
        }),
        createUser: builder.mutation<void, any>({
            query: (user) => ({
                url: 'users/create',
                method: 'POST',
                body: user,
            }),
        }),
        createPost: builder.mutation<void, CreatePostArgs>({
            query: ({ productName, post, channel }) => ({
                url: `products/create/post/${productName}/${channel}`,
                method: 'POST',
                body: post,
            }),
            invalidatesTags: ['Posts'],
        }),
        followProduct: builder.mutation<
            void,
            { name: string; follow: boolean; accessToken?: string }
        >({
            query: ({ name, follow, accessToken }) => ({
                url: `products/${follow ? 'follow' : 'unfollow'}/${name}`,
                method: 'POST',
                body: { accessToken: accessToken || '' },
            }),
            invalidatesTags: ['User'],
        }),
        createProduct: builder.mutation<{ message: string }, NewProduct>({
            query: (product) => ({
                url: 'products/create',
                method: 'POST',
                body: product,
            }),
            invalidatesTags: ['Product'],
        }),
        deleteProduct: builder.mutation<{ message: string }, { id: number }>({
            query: ({ id }) => ({
                url: `products/delete/product/${id}`,
                method: 'DELETE',
            }),
        }),
        updateUserRole: builder.mutation<
            void,
            { productId: number; userId: number; role: string }
        >({
            query: ({ userId, productId, role }) => ({
                url: `products/update/role/${productId}/${userId}`,
                method: 'PATCH',
                body: { role },
            }),
            invalidatesTags: ['User'],
        }),
        checkIfProductFollowed: builder.query<{ followed: boolean }, number>({
            query: (productId) => `products/isFollowed/${productId}`,
        }),
        getPostDetails: builder.query<
            Post,
            { postId: string; channel: string }
        >({
            query: ({ postId, channel }) =>
                `products/post/${channel}/${postId}`,
            providesTags: ['Post'],
        }),
        createComment: builder.mutation<
            Post,
            { field: string; postId: number; comment: string }
        >({
            query: (data) => ({
                url: 'products/create/comment',
                method: 'POST',
                body: data,
            }),
            invalidatesTags: ['Comments'],
        }),
        getPosts: builder.query<
            { posts: Post[]; lastId: number },
            { lastId?: number; productId: number; channel: string }
        >({
            query: ({ lastId, productId, channel }) =>
                `products/posts/${productId}/${channel}/${lastId || 0}`,
            providesTags: ['Posts'],
        }),
        getProductInfo: builder.query<
            { product: Product; users: ProductUser[] },
            string
        >({
            query: (productId) => `products/product/${productId}`,
            providesTags: ['User'],
        }),
        getProductDayPostCount: builder.query<
            {
                bugs: number
                announcements: number
                changelogs: number
                suggestions: number
            },
            number
        >({
            query: (productId) => `products/dayStats/${productId}`,
            providesTags: ['Stats'],
        }),
        likePost: builder.mutation<void, { postID: number; like: boolean }>({
            query: (data) => ({
                url: `products/${data.like ? 'like' : 'dislike'}_post/${
                    data.postID
                }`,
                method: 'PUT',
            }),
            invalidatesTags: ['Post'],
        }),
        getPostsBoard: builder.query<Post[], { productId: number }>({
            query: ({ productId }) => `products/board/${productId}`,
            providesTags: ['Board'],
        }),
        addToPostsBoard: builder.mutation<
            void,
            {
                productId: number
                postId: number
                field: 'working-on' | 'done' | 'under-review'
                refetchBoard: boolean
            }
        >({
            query: ({ productId, postId, field }) => ({
                url: `products/board/${productId}/${field}/add/${postId}`,
                method: 'PUT',
            }),
            invalidatesTags: (result, error, { refetchBoard }) => {
                return refetchBoard ? ['Post', 'Board', 'Stats'] : ['Post']
            },
        }),
        removeFromPostsBoard: builder.mutation<
            void,
            {
                productId: number
                postId: number
                field: 'working-on' | 'done' | 'under-review'
                refetchBoard: boolean
            }
        >({
            query: ({ productId, postId, field }) => ({
                url: `products/board/${productId}/${field}/remove/${postId}`,
                method: 'PUT',
            }),
            invalidatesTags: (result, error, { refetchBoard }) => {
                return refetchBoard ? ['Post', 'Board'] : ['Post']
            },
        }),
        deletePost: builder.mutation<
            void,
            { productId: number; postId: number; field: string }
        >({
            query: ({ productId, postId, field }) => ({
                url: `products/delete/post/${productId}/${field}/${postId}`,
                method: 'DELETE',
            }),
            invalidatesTags: ['Posts'],
        }),
        getPostComments: builder.query<
            { comments: Comment[]; lastId: number; hasMore: boolean },
            { postId: number; lastId?: number }
        >({
            query: ({ postId, lastId }) =>
                `products/comments/${postId}/${lastId || 0}`,
            providesTags: ['Comments'],
        }),
        getChatMessages: builder.query<
            void,
            { productId: number; lastId?: number }
        >({
            query: ({ productId, lastId }) =>
                `products/messages/${productId}/${lastId || 0}`,
        }),
        updateUser: builder.mutation<void, any>({
            query: (data: any) => ({
                url: `users/update`,
                method: 'PATCH',
                body: data,
            }),
            invalidatesTags: ['User'],
        }),
    }),
})

export const {
    useGetProductsQuery,
    useGetLatestPostsQuery,
    useSignInMutation,
    useLazyGetUserInfoQuery,
    useGetTopProductsQuery,
    useFollowProductMutation,
    useCreatePostMutation,
    useGetUserInfoQuery,
    useCreateProductMutation,
    useDeleteProductMutation,
    useLazyGetProductsQuery,
    useUpdateUserRoleMutation,
    useCreateUserMutation,
    useLazyCheckIfProductFollowedQuery,
    useLazyGetPostDetailsQuery,
    useCreateCommentMutation,
    useGetPostDetailsQuery,
    useGetPostsQuery,
    useLazyGetPostsQuery,
    useLazyGetProductInfoQuery,
    useGetProductDayPostCountQuery,
    useGetPostsBoardQuery,
    useAddToPostsBoardMutation,
    useRemoveFromPostsBoardMutation,
    useDeletePostMutation,
    useLikePostMutation,
    useLazyGetPostCommentsQuery,
    useLazyGetLatestPostsQuery,
    useLazyGetChatMessagesQuery,
    useUpdateUserMutation,
} = api
export default api
