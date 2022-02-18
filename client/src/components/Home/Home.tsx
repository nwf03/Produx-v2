import LatestPosts from "../Products/LatestPosts";
import Head from "next/head";
import FollowedProducts from "../Products/FollowedProducts";
import TopProducts from "../Products/TopProducts";
export default function Home() {
  return (
    <div className="grid grid-cols-4 justify-center align-middle h-[80px] lg:gap-2 sm:gap-2 ">
        <Head>
            <title>Home</title>
        </Head>
      <div className="hidden md:block">
        <FollowedProducts />
      </div>
      <div className="col-span-4 md:col-span-3 lg:col-span-2">
        <LatestPosts />
      </div>
      <div
        className={"hidden lg:flex lg:justify-center mt-6 mr-4"}
      >
        <TopProducts />
      </div>
    </div>
  );
}
