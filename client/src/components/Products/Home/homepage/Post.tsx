import {Post} from '../../../../state/interfaces'
import {Avatar} from "@nextui-org/react";
export default function PostCard( {post}: {post:Post}){
    return(
        <div className={'flex items-center bg-white p-4 m-4 rounded-box'}>
            <div className={'hidden lg:block'}>
            <Avatar src={post.user.pfp} />
            </div>
            <h1 className={'ml-3'}>{post.title}</h1>
        </div>
    )
}