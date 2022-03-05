import { useLazyGetLatestPostsQuery } from '../../state/reducers/api'
import Post from './Home/Posts/Post'
import { useEffect, useState, useRef, useCallback } from 'react'
import LoadingSpinner from '../LoadingSpinner'
import { Channel } from '../../state/interfaces'
import { Loading } from '@nextui-org/react'
import Posts from './Home/Posts/Posts'
export default function LatestPosts() {
    const [channel, setChannel] = useState('Bugs')
    const [getPosts, { data, isLoading, error }] = useLazyGetLatestPostsQuery()
    const [posts, setPosts] = useState([])
    const observer = useRef()
    const hasMore = data ? data.hasMore : false
    const [lastId, setLastId] = useState(0)
    const lastPostRef = useCallback(
        (node) => {
            if (isLoading) return
            if (observer.current) observer.current.disconnect()
            observer.current = new IntersectionObserver((entries) => {
                if (entries[0].isIntersecting) {
                    if (hasMore) {
                        setLastId(data.lastId)
                    }
                }
            })
            if (node) observer.current.observe(node)
        },
        [isLoading, hasMore]
    )
    const fetchPosts = (lastId?: number) => {
        getPosts({
            field: channel as string,
            lastId: lastId || 0,
        })
    }
    useEffect(() => {
        fetchPosts(lastId)
    }, [lastId])

    useEffect(() => {
        setPosts([])
        fetchPosts()
    }, [channel])

    useEffect(() => {
        if (data) {
            setPosts((p) => [...posts, ...data.posts])
        }
    }, [data])
    const channels: Channel = {
        Announcements: { icon: '🎉', color: '#FF9900' },
        Bugs: { icon: '🐞', color: '#DBFF00' },
        Suggestions: { icon: '🙏', color: '#0094FF' },
        Changelogs: { icon: '🔑', color: '#FF4D00' },
    }
    return (
        <div className="">
            <div className="block min-w-max m-5 bg-[#F5F5F5] text-center">
                <div>
                    {Object.keys(channels).map((c, idx) => (
                        <span
                            key={idx}
                            className={`px-5 text-sm overflow-auto mr-[2vw] py-3 hover:bg-gray-200 hover:cursor-pointer rounded-lg ${
                                channel == c && 'bg-gray-300'
                            }`}
                            onClick={() => setChannel(c)}
                        >
                            {channels[c].icon + ' ' + c}
                        </span>
                    ))}
                </div>
            </div>
            {!data && !isLoading && (
                <div className="text-center mt-12">no products followed</div>
            )}
            {posts.length > 0 ? (
                posts.map((p, idx) => {
                    return (
                        <div ref={idx == posts.length - 1 ? lastPostRef : null}>
                            <Post
                                key={idx}
                                showProductIcon={true}
                                data={p}
                                channel={channel}
                                color={channels[channel].color}
                                showDivider={false}
                            />
                        </div>
                    )
                })
            ) : (
                <p className="text-center mt-8">No posts here</p>
            )}
            {isLoading && <LoadingSpinner />}
        </div>
    )
}
