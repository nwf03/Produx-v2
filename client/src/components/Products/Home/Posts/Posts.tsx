import { useGetPostsQuery } from '../../../../state/reducers/api'
import { useCallback, useRef, useState } from 'react'
import Post from './Post'
import { Button } from '@nextui-org/react'
import { Channel } from '../../../../state/interfaces'
import AddPost from './AddPost'
export default function Posts({
    productId,
    productName,
    channel,
}: {
    productId: number
    productName: string
    channel: string
}) {
    const [showAdd, setShowAdd] = useState(false)
    const { data, isLoading, error } = useGetPostsQuery({
        productId,
        channel: channel.toLowerCase(),
    })
    const channels: Channel = {
        Announcements: { icon: '🎉', color: '#FF9900' },
        Bugs: { icon: '🐞', color: '#DBFF00' },
        Suggestions: { icon: '🙏', color: '#0094FF' },
        Changelogs: { icon: '🔑', color: '#FF4D00' },
    }
    const [posts, setPosts] = useState(data ? data.posts : [])
    const observer = useRef()
    const lastPostRef = useCallback(
        (node) => {
            if (isLoading) return
            if (observer.current) observer.current.disconnect()
            observer.current = new IntersectionObserver((entries) => {
                if (entries[0].isIntersecting) {
                    // TODO: load more posts
                }
            })
            if (node) observer.current.observe(node)
        },
        [isLoading]
    )

    return (
        <div className="items-center justify-center overflow-x-hidden">
            <div
                className={
                    'grid grid-rows-2 md:flex md:items-center md:ml-12 mt-4'
                }
            >
                <h1 className="text-3xl font-bold ">Latest {channel}</h1>
                <Button
                    className={'z-0 bg-black md:ml-auto mr-4'}
                    onClick={() => setShowAdd(!showAdd)}
                >
                    Create Post
                </Button>
            </div>
            {showAdd && (
                <AddPost
                    show={showAdd}
                    setShow={setShowAdd}
                    productName={productName}
                />
            )}
            <br />
            {isLoading && 'Loading...'}
            <div className="flex items-center justify-center">
                {data && posts.length > 0 ? (
                    <div className="w-screen">
                        {posts.map((post, idx) => {
                            return (
                                <div
                                    key={idx}
                                    ref={
                                        idx == posts.length - 1
                                            ? lastPostRef
                                            : null
                                    }
                                >
                                    {idx == posts.length - 1 && 'LAST ONE'}
                                    <Post
                                        showProductIcon={false}
                                        data={post}
                                        channel={channel}
                                        color={channels[channel].color}
                                        showDivider={true}
                                    />
                                </div>
                            )
                        })}
                    </div>
                ) : (
                    <h1 className={'text-2xl w-screen text-left ml-12'}>
                        Channel is empty :(
                    </h1>
                )}
            </div>
        </div>
    )
}
