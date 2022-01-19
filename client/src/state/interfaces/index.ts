export interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: null;
  username: string;
  password: string;
  email: string;
  pfp: string;
  products?: Product[] | null;
  followed_products?: Product[] | null;
  liked_products?: Product[] | null;
}
export interface Product {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: null;
  userID: number;
  name: string;
  description: string;
  images?: string[] | null;
  user_likes?: null;
  likes: number;
  verified: boolean;
}

export interface Post {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: null;
  productID: number;
  product: Product;
  user: Partial<User>;
  userID: number;
  title: string;
  description: string;
}

export interface UserInfo {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: null;
  name: string;
  email: string;
  pfp: string;
  products?: Product[] | null;
  followed_products?: Product[] | null;
  liked_products?: Product[] | null;
}
export interface Channel {
  [key: string]: {
    icon: string;
    color: string;
  };
}

export type PostTypes = "bugs" | "suggestions" | "announcements" | "changelogs";

export interface ProductPostsResponse {
  posts: Post[];
  page: number;
}
