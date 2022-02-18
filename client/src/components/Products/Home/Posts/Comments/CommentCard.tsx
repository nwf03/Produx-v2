import { Comment } from "../../../../../state/interfaces";
import { Avatar, Link } from "@nextui-org/react";
import moment from "moment";
import { userInfo } from "os";
export default function CommentCard({ comment }: { comment: Comment }) {
  return (
    <div className="bg-white p-5 rounded-3xl min-w-[20vw] items-center mx-12">
      <div className="flex items-center">
        {comment.user.pfp && (
          <Avatar
            className="z-0"
            squared
            bordered
            src={comment.user.pfp}
            size="lg"
          />
        )}
        <p className={"ml-2"}>{comment.user.name}</p>
        <p className="ml-2 text-gray-400 text-sm">
          {moment(new Date(comment.UpdatedAt)).fromNow()}
        </p>
      </div>
      <div>
        <div className="mt-3">
          <span className={"break-all"}>{comment.comment}</span>
        </div>
      </div>
      <div className={"flex justify-center "}>
        <div className={"divider w-full h-0"} />
      </div>
    </div>
  );
}
