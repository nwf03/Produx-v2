export interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: null;
  name: string;
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
  private: boolean;
  accessToken: string;
}
export interface Comment {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
//  add bug id or suggestions id
  user: User;
  userID: number;
  comment: string
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
  commentsCount?: number;
  likes: number;
  dislikes: number;
  comments?: Comment[];
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

type ChannelName = "bugs" | "suggestions" | "announcements" | "changelogs";
type ChannelBody = {
  icon: string;
  color: string;
};

type ChannelInfo = Record<ChannelName, ChannelBody>;

const b1: ChannelInfo = {
  bugs: { icon: "", color: "" },
};

export interface ProductPostsResponse {
  posts: Post[];
  lastId: number
}

type images = { images: File[] | null };
export type NewProduct = Pick<
  Product,
  "name" | "description" | "private" | "access_token"
> &
  images;
type ChannelNames = "Announcements" | "Changelogs" | "Bugs" | "Suggestions";
export type Channels = Record<ChannelNames, boolean>;

export type ProductUser = Pick<User, "name" | "pfp"> & {
  role: string;
  id: number;
};

export type PostTags = "done" | "working-on" | "under-review"