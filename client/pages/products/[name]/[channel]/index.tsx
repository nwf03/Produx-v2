import { useRouter } from "next/router";
import Head from "next/head";
import HomeLayout from "../../../../src/components/Products/Home/HomeLayout";
import { useEffect } from "react";
import { useGetProductsQuery } from "../../../../src/state/reducers/api";
import Post from "../../../../src/components/Products/Post";
import { ProductPostsResponse } from "../../../../src/state/interfaces";
import {Button, Loading} from "@nextui-org/react";
import { useState } from "react";
import { setChannel } from "../../../../src/state/reducers/channelSlice";
import AddPost from "../../../../src/components/Products/Home/AddPost";
import { useAppDispatch } from "../../../../src/state/hooks";
import LoadingSpinner from "../../../../src/components/LoadingSpinner";
export default function Channel() {
  const [showAdd, setShowAdd] = useState(false);
  const router = useRouter();
  const dispatch = useAppDispatch();
  const { name, channel } = router.query;
  const { data, error, isLoading } = useGetProductsQuery({
    name: name as string,
    field: channel as string,
    page: 1,
  });
  if (channel) {
    dispatch(setChannel(channel as string));
  }
  const channels = {
    Announcements: { icon: "🎉", color: "#FF9900" },
    Bugs: { icon: "🐞", color: "#DBFF00" },
    Suggestions: { icon: "🙏", color: "#0094FF" },
    Changelogs: { icon: "🔑", color: "#FF4D00" },
  };
  const { posts } = data ? (data as ProductPostsResponse) : { posts: [] };
  return (
    <div className="items-center justify-center overflow-x-hidden">
      <div className={"grid grid-rows-2 md:flex md:items-center md:ml-12 mt-4"}>
        <h1 className="text-3xl font-bold ">Latest {channel}</h1>
        <Button
          className={"bg-black md:ml-auto mr-4"}
          onClick={() => setShowAdd(!showAdd)}
        >
          Create Post
        </Button>
      </div>
      {showAdd && (
        <AddPost show={showAdd} setShow={setShowAdd} productName={name} />
      )}
      <br />
      {isLoading && <LoadingSpinner />}
      <div className="flex items-center justify-center">
        {data && posts.length > 0 ? (
          <div className="w-screen">
            {posts.map((post, idx) => {
              return (
                <Post
                  showProductIcon={false}
                  key={idx}
                  data={post}
                  channel={channel as string}
                  color={channels[channel].color}
                  showDivider={true}
                />
              );
            })}
          </div>
        ) : (
          <h1 className={"text-2xl w-screen text-left ml-12"}>
            Channel is empty :(
          </h1>
        )}
      </div>
    </div>
  );
}
Channel.Layout = HomeLayout;
