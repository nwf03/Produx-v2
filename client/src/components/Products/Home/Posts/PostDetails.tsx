import { Post } from "../../../../state/interfaces";
import PostView from "./Post";
import CommentCard from "./Comments/CommentCard";
import {Avatar, Button, Link, Tooltip} from "@nextui-org/react";
import moment from "moment";
import AddComment from "./Comments/AddComment";
import { useRouter } from "next/router";
import { useLazyGetPostDetailsQuery, useRemoveFromPostsBoardMutation, useLikePostMutation } from "../../../../state/reducers/api";
import { useEffect, useState } from "react";
import LoadingSpinner from "../../../LoadingSpinner";
import {useAppSelector} from "../../../../state/hooks";
import {PostTags} from "../../../../state/interfaces";
import AddPostToBoard from "./AddPostToBoard";
export default function PostDetails() {
  const router = useRouter();
  const isOwner = useAppSelector(state => state.product.isOwner);
  const { channel, postId } = router.query;
  const [getPostDetails, { error, data, isLoading }] =
    useLazyGetPostDetailsQuery();
  useEffect(() => {
    getPostDetails({
      postId: postId as string,
      channel: channel as string,
    });
  }, [postId, channel]);
  const [removePostFromBoard] = useRemoveFromPostsBoardMutation()
  const [likePost] = useLikePostMutation()
  const handlePostLike = async () => {
    if (data){
      await likePost({
        postID: data.ID,
        like: true
      })
    }
  }
  const handleRemoveTag = async (field: PostTags) => {
    if(data) {
      await removePostFromBoard({
        productId: data.product.ID,
        postId: data.ID,
        field
      })
    }
  }
  const tagRemovable = (tag: PostTags) : boolean =>{
    if (tag === "working-on" || tag === "done" || tag == "under-review"){
      return true
    }
    return false
  }
  const [showTooltip, setShowToolTip] = useState(false)
  return (
    <div>
      {data && (
        <div>
          <div className="bg-white p-5 rounded-3xl  min-w-[20vw] mb-4 items-center mx-12 mt-4 md:mt-0">
            <button
              onClick={() => router.back()}
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
              <div className={"flex items-center my-auto ml-auto"}>
              {data.type.map((type,idx) => (
              <span key={idx}
                style={{ backgroundColor: "orange" }}
                className={`text-black  rounded-[10px] items-center flex mx-1 px-2 py-[4.5px] text-2xs  cursor-pointer`}
                    onClick={() => {
                      isOwner && tagRemovable(type)? handleRemoveTag(type) : null
                    }}
              >
                {isOwner && tagRemovable(type) &&
                <svg xmlns="http://www.w3.org/2000/svg" className="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                </svg>
                }
                {type}
                </span>
              ))}
              {isOwner &&
                  <Tooltip className="ml-auto" color={"invert"} placement={"bottom"} trigger="click" content={<AddPostToBoard postTypes={data.type} postId={data.ID} productId={data.product.ID}/>}>
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z" clipRule="evenodd" />
                  </svg>
                  </Tooltip>
              }
              </div>
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
                <p className={"mx-1"}>|</p>
                <div className={"flex z-50 relative"}>
                <Tooltip  content={"Like post"}>
                  <svg onClick={() => handlePostLike()} xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5" />
                  </svg>
                  <p className={'ml-1'}>{data.likes}</p>
                </Tooltip>
                </div>
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
      {isLoading && <div className={'h-screen'}><LoadingSpinner height={'full'} /></div>}
    </div>
  );

}
