import { Post } from "../../state/interfaces";
import PostView from "./Post";
import CommentCard from "./CommentCard";
import { Avatar, Button, Link } from "@nextui-org/react";
import moment from "moment";
import AddComment from "./Home/AddComment";
import { useRouter } from "next/router";
import { useLazyGetPostDetailsQuery } from "../../state/reducers/api";
import { useEffect } from "react";

export default function PostDetails() {
  const router = useRouter();
  const { channel, postId } = router.query;
  const [getPostDetails, { error, data, isLoading }] =
    useLazyGetPostDetailsQuery();
  useEffect(() => {
    getPostDetails({
      postId: postId as string,
      channel: channel as string,
    });
  }, [postId, channel]);

  return (
    <div>
      {data && (
        <div>
          <div className="bg-white p-5 rounded-3xl  min-w-[20vw] mb-4 items-center mx-12 mt-4 md:mt-0">
            <button
              onClick={() => router.replace(`/products/${data.product.name}`)}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M10 19l-7-7m0 0l7-7m-7 7h18"
                />
              </svg>
            </button>
            <div className="grid-cols-3 flex align-bottom mt-4">
              <div className="flex items-center">
                {data.user.pfp && (
                  <Avatar
                    className="z-0"
                    squared
                    bordered
                    src={data.user.pfp}
                    size="xl"
                  />
                )}
                <span className="font-bold text-xl pl-3 lg:mr-1 ">
                  {data.title}
                </span>
                <span className="text-sm text-truncate ">
                  - {data.product.name}
                </span>
              </div>
              <span
                style={{ backgroundColor: "orange" }}
                className={`text-black  rounded-[10px]  px-2 py-2 text-xs text-center my-auto ml-auto`}
              >
                {channel.substring(0, channel.length - 1)}
              </span>
            </div>
            <div>
              <div className="mt-3">
                <span className={"text-xl break-all"}>{data.description}</span>
              </div>
              <div className={"text-sm mt-2 text-gray-400 flex items-center"}>
                <p className="">
                  {moment(new Date(data.CreatedAt)).calendar()}
                </p>
                <p className={"mx-1"}>|</p>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z"
                  />
                </svg>
                <p className={"ml-1"}>{data.commentsCount}</p>
              </div>
            </div>
          </div>
          <AddComment field={channel as string} postId={data.ID} />
          {data.comments &&
            data.comments.map((m, idx) => {
              return (
                <div key={idx}>
                  <CommentCard comment={m} />
                </div>
              );
            })}
        </div>
      )}
    </div>
  );
}
