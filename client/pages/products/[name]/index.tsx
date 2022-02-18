import HomeLayout from "../../../src/components/Products/Home/HomeLayout";
import { useEffect } from "react";
import {
  useGetProductDayPostCountQuery,
  useGetPostsQuery,
} from "../../../src/state/reducers/api";
import Post from "../../../src/components/Products/Home/homepage/Post";
import { useAppDispatch } from "../../../src/state/hooks";
import { setChannel } from "../../../src/state/reducers/channelSlice";
export default function ProductHome({ productId }: { productId: number }) {
  const { data, isLoading, error } = useGetProductDayPostCountQuery(productId);
  const posts = useGetPostsQuery({
    productId,
    channel: "bugs",
  });
  const colors = [
    "from-red-500 to-orange-300",
    "from-orange-500 to-yellow-400",
    "from-yellow-400 to-green-400",
    "from-green-400 to-blue-400",
  ];
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(setChannel("Home"));
  }, []);

  return (
    <div className={"mt-4 ml-4"}>
      <h1 className={"font-bold text-2xl"}>Today&apos;s Numbers</h1>
      <div className={"grid grid-cols-2 gap-y-3 mt-4"}>
        {data &&
          Object.keys(data).map((c, idx) => (
            <div
              className={`h-14 mx-5 font-bold flex justify-center items-center py-3 px-9 rounded-2xl bg-gradient-to-r ${colors[idx]}  text-white mx-4`}
            >
              <span className="font-bold text-center text-md">
                {data[c]} {data[c] == 1 ? c.substring(0, c.length - 1) : c}
              </span>
            </div>
          ))}
      </div>
      <div className="divider w-11/12 mx-auto" />
      <div className="grid  text-center  grid-cols-1 md:grid-cols-2 lg:grid-cols-2 ">
        <div className="bg-red-400 rounded-xl m-4">
          <p className={'mt-5 text-white text-xl'}>Under Review</p>
            <div className="h-1" />
            {posts.data &&
              posts.data.posts.map((t, idx) => {
                return (
                  <div key={idx}>
                    <Post post={t} />
                  </div>
                );
              })}
        </div>
        <div className={"bg-yellow-400 rounded-xl m-4"}>
          <p className={"mt-5 text-white text-xl"}>Working on </p>
            <div className="h-1" />
            {[...Array(13)].map((t, idx) => {
              return (
                <div key={idx} className="bg-white rounded-box p-4 mx-4 my-4">
                  <h1>Post {idx + 1}</h1>
                </div>
              );
            })}
        </div>
        <div className="bg-blue-400 rounded-xl m-4 md:col-span-2 lg:col-span-1">
          <p className={'mt-5 text-white text-xl'}>Done</p>
          <div className="">
            <div className="h-1" />
            {[...Array(13)].map((t, idx) => {
              return (
                <div key={idx} className="bg-white rounded-box p-4 mx-4 my-4">
                  <h1>Post {idx + 1}</h1>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
}

ProductHome.Layout = HomeLayout;
