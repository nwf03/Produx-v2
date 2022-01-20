import { useGetProductsQuery } from "../../../state/reducers/api";
import {useEffect, useState} from "react";
import { ProductPostsResponse } from "../../../state/interfaces";
import Post from "../Post";
import { Channel } from "../../../state/interfaces";
export default function Posts({
  name,
  channel,
}: {
  name: string;
  channel: string;
}) {
  const { data, isLoading, error } = useGetProductsQuery({
    name: name,
    field: channel,
    page: 1,
  });
  const channels: Channel = {
    Announcements: { icon: "🎉", color: "#FF9900" },
    Bugs: { icon: "🐞", color: "#DBFF00" },
    Suggestions: { icon: "🙏", color: "#0094FF" },
    Changelogs: { icon: "🔑", color: "#FF4D00" },
  };
  const { posts } = data ? (data as ProductPostsResponse) : { posts: [] };
  useEffect(()=>{
    console.log("channels color: ", channel)
  },[])
  return (
    <div className="items-center justify-center overflow-x-hidden">
      <h1 className="text-3xl font-bold ml-12 mt-4">Latest {channel}</h1>

      <br />
      {isLoading && "Loading..."}
      <div className="flex items-center justify-center">
        {data && posts.length > 0 ? (
          <div className="">
            {posts.map((post, idx) => {
              return (
                <Post
                  key={idx}
                  data={post}
                  channel={channel}
                  color={channels[channel].color}
                  showDivider={true }
                />
              );
            })}
          </div>
        ) : (
          <h1 className={"text-2xl w-screen text-left ml-12"}>Channel is empty :(</h1>
        )}
      </div>
    </div>
  );
}
