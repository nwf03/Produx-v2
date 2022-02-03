import {Comment} from '../../state/interfaces'
import {Avatar, Link} from "@nextui-org/react";
import moment from "moment";
export default function CommentCard({comment}: {comment:Comment}) {
    return (
            <div className="bg-white p-5 rounded-3xl min-w-[20vw] items-center mx-12 mt-4 md:mt-0">
                    <div className="flex items-center">
                        { comment.user.pfp && (
                            <Avatar squared bordered src={comment.user.pfp} size="lg" />
                        )}
                        <p className={'ml-2'}>{comment.user.name}</p>
                    </div>
                <div>
                    <div className="mt-3">
                        <span className={"break-all"}>{comment.comment}</span>
                        <p className="text-gray-400 text-sm">{moment(new Date(comment.UpdatedAt)).fromNow()}</p>
                    </div>
            </div>
                <div className={"flex justify-center "}>
                    <div className={"divider w-11/12"} />
                </div>
        </div>
    )
}