import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { Product, User, Post, UserInfo } from "../interfaces";
import { RootState } from "../store";
import { ProductPostsResponse } from "../interfaces";
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
  endpoints: (builder) => ({
    getProducts: builder.query<
      ProductResponse | Product | ProductPostsResponse,
      ProductSearch
    >({
      query: (search) =>
        `products/search/${search.name || -1}/${search.field || -1}/${
          search.page
        }`,
    }),
    getLatestPosts: builder.query<Post[], string>({
      query: (channel) => `products/latest_posts/${channel}`,
    }),
    getUserInfo: builder.query<UserInfo, void>({
      query: () => "users/user/get",
    }),
    getTopProducts: builder.query<Product[] | null, void>({
      query: () => "products/today_top_products",
    }),
    signIn: builder.mutation<
      { token: string },
      Pick<User, "username" | "password">
    >({
      query: (userInfo) => ({
        url: "users/login",
        method: "POST",
        body: userInfo,
      }),
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetLatestPostsQuery,
  useSignInMutation,
  useLazyGetUserInfoQuery,
  useGetTopProductsQuery,
} = api;
export default api;
