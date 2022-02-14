import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { NewProduct, Product, ProductUser, User, Post, UserInfo } from "../interfaces";
import { RootState } from "../store";
import { ProductPostsResponse } from "../interfaces";
import { number } from "prop-types";
interface ProductSearch {
  name?: string;
  field?: string;
  page: number;
}
export interface ProductResponse {
  page: number;
  products: Product[];
}
interface OneProductResponse {
  product: Product;
}

interface CreatePostArgs {
  post: Pick<Post, "title" | "description">;
  productName: string;
  channel: string;
}
const api = createApi({
  reducerPath: "api",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:8000",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem("token");
      if (token) {
        headers.set("authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  tagTypes: ["Posts", "User", "Product", "Comments"],
  endpoints: (builder) => ({
    getProducts: builder.query<
      ProductResponse | Product | ProductPostsResponse,
      ProductSearch
    >({
      query: (search) =>
        `products/search/${search.name || -1}/${search.field || -1}/${
          search.page
        }`,
      providesTags: ["Posts", "User"],
    }),
    getLatestPosts: builder.query<Post[], string>({
      query: (channel) => `products/latest_posts/${channel}`,
      providesTags: ["Posts", "User"],
    }),
    getUserInfo: builder.query<UserInfo, void>({
      query: () => "users/user/get",
      providesTags: ["User"],
    }),
    getTopProducts: builder.query<Product[] | null, void>({
      query: () => "products/today_top_products",
      providesTags: ["Product"],
    }),
    signIn: builder.mutation<
      { token: string },
      { username: string; password: string }
    >({
      query: (userInfo) => ({
        url: "users/login",
        method: "POST",
        body: userInfo,
      }),
    }),
    createUser: builder.mutation<void, any>({
      query: (user) => ({
        url: "users/create",
        method: "POST",
        body: user,
      }),
    }),
    createPost: builder.mutation<void, CreatePostArgs>({
      query: ({ productName, post, channel }) => ({
        url: `products/create/post/${productName}/${channel}`,
        method: "POST",
        body: post,
      }),
      invalidatesTags: ["Posts"],
    }),
    followProduct: builder.mutation<
      void,
      { name: string; follow: boolean; accessToken?: string }
    >({
      query: ({ name, follow, accessToken }) => ({
        url: `products/${follow ? "follow" : "unfollow"}/${name}`,
        method: "POST",
        body: { accessToken: accessToken || "" },
      }),
      invalidatesTags: ["User"],
    }),
    createProduct: builder.mutation<{message:string}, NewProduct>({
      query: (product) => ({
        url: "products/create",
        method: "POST",
        body: product,
      }),
      invalidatesTags: ["Product"],
    }),
    deleteProduct: builder.mutation<{ message: string }, { id: number }>({
      query: ({ id }) => ({
        url: `products/delete/product/${id}`,
        method: "DELETE",
      }),
    }),
    updateUserRole: builder.mutation<
      void,
      { productId: number; userId: number; role: string }
    >({
      query: ({ userId, productId, role }) => ({
        url: `products/update/role/${productId}/${userId}`,
        method: "PATCH",
        body: { role },
      }),
      invalidatesTags: ["User"],
    }),
    checkIfProductFollowed: builder.query<{followed:boolean
    }, number>({
      query: (productId) => `products/isFollowed/${productId}`,
    }),
    getPostDetails: builder.query<Post, {postId: string,channel:string}>({
      query: ({postId, channel}) => `products/comments/${channel}/${postId}`,
      providesTags: ["Comments"]
    }),
  createComment: builder.mutation<Post, {field: string, postId: number, comment:string}>({
   query: (data) => ({
     url : "products/create/comment",
     method: "POST",
     body: data
   }),
    invalidatesTags: ["Comments"]
  }),
    getPosts: builder.query<{posts: Post[], lastId: number}, {lastId?: number, productId: number, channel: string}>({
      query: ({lastId, productId, channel}) => `products/posts/${productId}/${channel}/${lastId || 0}`,
      providesTags: ["Posts"],
    }),
    getProductInfo: builder.query<{product: Product, users: ProductUser[]}, string>({
      query: (productId) => `products/product/${productId}`,
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetLatestPostsQuery,
  useSignInMutation,
  useLazyGetUserInfoQuery,
  useGetTopProductsQuery,
  useCreatePostMutation,
  useFollowProductMutation,
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
    useLazyGetProductInfoQuery
} = api;
export default api;