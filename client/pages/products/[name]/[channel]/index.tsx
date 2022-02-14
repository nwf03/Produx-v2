import { useRouter } from "next/router";
import Head from "next/head";
import HomeLayout from "../../../../src/components/Products/Home/HomeLayout";
import { useState, useCallback, useEffect, useRef } from "react";
import { useLazyGetPostsQuery } from "../../../../src/state/reducers/api";
import Post from "../../../../src/components/Products/Post";
import { ProductPostsResponse } from "../../../../src/state/interfaces";
import { Button, Loading } from "@nextui-org/react";
import { setChannel } from "../../../../src/state/reducers/channelSlice";
import AddPost from "../../../../src/components/Products/Home/AddPost";
import { useAppDispatch } from "../../../../src/state/hooks";
import LoadingSpinner from "../../../../src/components/LoadingSpinner";
export default function Channel({ productId }: { productId: number }) {
  const [showAdd, setShowAdd] = useState(false);
  const router = useRouter();
  const dispatch = useAppDispatch();
  const { name, channel } = router.query;
  const [getPosts, { data, error, isLoading }] = useLazyGetPostsQuery();
  if (channel) {
    dispatch(setChannel(channel as string));
  }
  const channels = {
    Announcements: { icon: "🎉", color: "#FF9900" },
    Bugs: { icon: "🐞", color: "#DBFF00" },
    Suggestions: { icon: "🙏", color: "#0094FF" },
    Changelogs: { icon: "🔑", color: "#FF4D00" },
  };
  const [lastId, setLastId] = useState(0);
  useEffect(() => {
    getPosts({
      productId: productId,
      channel: (channel as string).toLowerCase(),
      lastId: lastId,
    });
    console.log("getting posts");
  }, [lastId]);
  useEffect(() => {
    if (data) {
      setPosts([...posts, ...data.posts]);
    }
  }, [data]);
  useEffect(() => {
    setPosts([]);
    getPosts({
      productId: productId,
      channel: (channel as string).toLowerCase(),
    });
  }, [channel]);

  const [posts, setPosts] = useState(data ? data.posts : []);
  const observer = useRef();
  const hasMore = data ? data.hasMore : false;
  const lastPostRef = useCallback(
    (node) => {
      if (isLoading) return;
      if (observer.current) observer.current.disconnect();
      observer.current = new IntersectionObserver((entries) => {
        if (entries[0].isIntersecting) {
          console.log(data.hasMore, data.lastId);
          setLastId(data.lastId);
        }
      });
      if (node) observer.current.observe(node);
    },
    [isLoading, hasMore]
  );
  return (
    <div className="overflow-y-hidden items-center justify-center overflow-x-hidden">
      <div className={"grid grid-rows-2 md:flex md:items-center md:ml-12 mt-4"}>
        <h1 className="text-3xl font-bold ">Latest {channel}</h1>
        <Button
          className={"z-0 bg-black md:ml-auto mr-4"}
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
        {posts && posts.length > 0 ? (
          <div className="w-screen">
            {posts.map((post, idx) => {
              return (
                <div
                  key={idx}
                  ref={idx == posts.length - 1 ? lastPostRef : null}
                >
                  <Post
                    showProductIcon={false}
                    data={post}
                    channel={channel as string}
                    color={channels[channel].color}
                    showDivider={true}
                  />
                </div>
              );
            })}
          </div>
        ) : (
          <h1 className={"text-2xl w-screen text-left ml-12"}>
            Channel is empty :(
          </h1>
        )}
      </div>
      {isLoading && <LoadingSpinner height={"auto"} />}
    </div>
  );
}
Channel.Layout = HomeLayout;
