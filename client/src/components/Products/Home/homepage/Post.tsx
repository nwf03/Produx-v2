import { Post } from '../../../../state/interfaces'
import { useRouter } from 'next/router'
import { useDrag } from 'react-dnd'
import { useEffect, useState } from 'react'
import { useRemoveFromPostsBoardMutation } from '../../../../state/reducers/api'
export default function PostCard({
    post,
    setList,
    productId,
}: {
    post: Post
    setList: any
    productId: number
}) {
    const router = useRouter()
    const [removePostFromBoard] = useRemoveFromPostsBoardMutation()
    const pathToPost = `/products/[name]/[channel]/[postId]`
    const { name } = router.query
    const [{ isDragging }, drag] = useDrag(() => ({
        type: 'post',
        item: {
            post: post,
        },
        collect: (monitor) => ({
            isDragging: !!monitor.isDragging(),
        }),
    }))
    useEffect(() => {
        if (isDragging) {
            setList((items: Post[]) =>
                items.filter((i: Post) => i.ID !== post.ID)
            )
        }
    }, [isDragging])

    return (
        <div
            ref={drag}
            onClick={() => {
                router.push({
                    pathname: pathToPost,
                    query: {
                        name: post.product.name || (name as string),
                        channel: post.type[0],
                        postId: post.ID,
                    },
                })
            }}
            className={`flex w-12/12 cursor-pointer bg-white rounded-md p-3 my-3 mx-3 `}
        >
            <h1 className={'text-left ml-2'}>{post.title}</h1>
        </div>
    )
}
