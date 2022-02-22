import {Post} from '../../../../state/interfaces'
import {useRouter} from "next/router";
import {useDrag} from "react-dnd";
import {useEffect} from "react";
import {useRemoveFromPostsBoardMutation} from "../../../../state/reducers/api";
export default function PostCard( {post, setList, productId}: {post:Post, setList:any, productId: number}){
    const router = useRouter()
    const [removePostFromBoard] = useRemoveFromPostsBoardMutation()
    const pathToPost = `/products/[name]/[channel]/[postId]`;
    const {name} = router.query;
    const [{isDragging}, drag] = useDrag(() => ({
        type: "post",
        item: {
            post: post,
        },
        collect: (monitor) => ({
            isDragging: !!monitor.isDragging(),
        }),
    }));
    useEffect(() => {
        if (isDragging){
          setList(items => items.filter(i => i.ID !== post.ID) )
          console.log("post removed: ", post)
          removePostFromBoard({
              productId,
              postId: post.ID,
              field: post.type[post.type.length - 1]
            })
        }
    }, [isDragging]);

    return(
        <div ref={drag} onClick={() => {
            router.push({
                pathname: pathToPost,
                query: {
                    name: post.product.name || name as string,
                    channel: post.type[0],
                    postId: post.ID,
                },
            })
        }} className={'flex w-12/12 cursor-pointer bg-white rounded-md p-3 my-4 mx-4 '}>
            <h1 className={'ml-2'}>{post.title}</h1>
        </div>
    )
}
