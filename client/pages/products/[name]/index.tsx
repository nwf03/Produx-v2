import HomeLayout from '../../../src/components/Products/Home/HomeLayout'
import { useEffect, useState, useRef } from 'react'
import {
    useGetProductDayPostCountQuery,
    useGetPostsBoardQuery,
} from '../../../src/state/reducers/api'
import PostCard from '../../../src/components/Products/Home/homepage/Post'
import { Post } from '../../../src/state/interfaces'
import { useAppDispatch } from '../../../src/state/hooks'
import { setChannel } from '../../../src/state/reducers/channelSlice'
import { useDrop } from 'react-dnd'
import { useAddToPostsBoardMutation } from '../../../src/state/reducers/api'
import Board from '../../../src/components/Home/Board'

export default function ProductHome({ productId }: { productId: number }) {
    const posts = useGetPostsBoardQuery({
        productId,
    })
    const [underReview, setUnderReveiw] = useState<any[]>(
        posts.data ? posts.data.underReview : []
    )
    const [workingOn, setWorkingOn] = useState<any[]>(
        posts.data ? posts.data.workingOn : []
    )
    const [done, setDone] = useState<any[]>(posts.data ? posts.data.done : [])

    const [addPostToBoard] = useAddToPostsBoardMutation()
    const [URD, UnderReviewDrop] = useDrop(() => ({
        accept: 'post',
        drop: async (item: any) => {
            addItemToBoard(item.post, setUnderReveiw)
            await addPostToBoard({
                productId,
                postId: item.post.ID,
                field: 'under-review',
                refetchBoard: false,
            })
        },
        collect: (monitor) => ({
            isOver: !!monitor.isOver(),
            canDrop: !!monitor.canDrop(),
        }),
    }))

    const [WOD, WorkingOnDrop] = useDrop(() => ({
        accept: 'post',
        drop: async (item: any) => {
            addItemToBoard(item.post, setWorkingOn)
            await addPostToBoard({
                productId,
                postId: item.post.ID,
                field: 'working-on',
                refetchBoard: false,
            })
        },
        collect: (monitor) => ({
            isOver: !!monitor.isOver(),
            canDrop: !!monitor.canDrop(),
        }),
    }))
    const [DD, DoneDrop] = useDrop(() => ({
        accept: 'post',
        drop: async (item: any) => {
            addItemToBoard(item.post, setDone)
            await addPostToBoard({
                productId,
                postId: item.post.ID as number,
                field: 'done',
                refetchBoard: false,
            })
        },
        collect: (monitor) => ({
            isOver: !!monitor.isOver(),
            canDrop: !!monitor.canDrop(),
        }),
    }))
    const { data, isLoading, error } = useGetProductDayPostCountQuery(productId)
    const colors = [
        'from-red-500 to-orange-300',
        'from-orange-500 to-yellow-400',
        'from-yellow-400 to-green-400',
        'from-green-400 to-blue-400',
        'from-green-400 to-blue-400',
        'from-green-400 to-blue-400',
        'from-green-400 to-blue-400',
    ]

    const dispatch = useAppDispatch()
    const [draggingPost, setDraggingPost] = useState<Post | null>(null)
    useEffect(() => {
        dispatch(setChannel('Home'))
    }, [])
    useEffect(() => {
        if (posts.data) {
            setUnderReveiw(posts.data.underReview)
            setWorkingOn(posts.data.workingOn)
            setDone(posts.data.done)
        }
    }, [posts.data])

    // todo fix item not being added to board
    const addItemToBoard = (item: any, setList: any) => {
        setList((prevState) => {
            const newItems = prevState
                .filter((i) => i.ID !== item.ID)
                .concat({ ...item })
            return [...newItems]
        })
    }
    const monthNames = [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December',
    ]
    const boards = [
        {
            name: 'Under Review',
            posts: underReview,
            ref: UnderReviewDrop,
            setPosts: setUnderReveiw,
            bgcolor: 'bg-orange-400',
            dropInfo: URD,
            icon: (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="mr-2 h-6 w-6"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2h-1.528A6 6 0 004 9.528V4z" />
                    <path
                        fillRule="evenodd"
                        d="M8 10a4 4 0 00-3.446 6.032l-1.261 1.26a1 1 0 101.414 1.415l1.261-1.261A4 4 0 108 10zm-2 4a2 2 0 114 0 2 2 0 01-4 0z"
                        clipRule="evenodd"
                    />
                </svg>
            ),
        },
        {
            name: 'Working On',
            posts: workingOn,
            ref: WorkingOnDrop,
            setPosts: setWorkingOn,
            bgcolor: 'bg-blue-500',
            dropInfo: WOD,
            icon: (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="mr-2 h-6 w-6"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
                </svg>
            ),
        },
        {
            name: 'Done',
            posts: done,
            setPosts: setDone,
            ref: DoneDrop,
            bgcolor: 'bg-green-500',
            dropInfo: DD,
            icon: (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="mr-2 h-6 w-6"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path
                        fillRule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clipRule="evenodd"
                    />
                </svg>
            ),
        },
    ]
    return (
        <div className={'mt-4 ml-4'}>
            <h1 className={'font-bold text-2xl'}>Today&apos;s Numbers</h1>
            <div
                className={
                    'grid grid-cols-2 md:grid-cols-3 lg:grid-cols-3 gap-y-3 mt-4'
                }
            >
                {data &&
                    Object.keys(data).map((c, idx) => (
                        <div
                            className={`h-14 mx-4  font-bold flex justify-center items-center py-3 px-9 rounded-2xl bg-gradient-to-r ${colors[idx]}  text-white mx-4`}
                            key={idx}
                        >
                            <span className="font-bold text-center text-md">
                                {data[c]}{' '}
                                {data[c] == 1
                                    ? c.substring(0, c.length - 1)
                                    : c}
                            </span>
                        </div>
                    ))}
            </div>
            <div className="divider w-11/12 mx-auto" />
            <h1 className={'font-bold text-2xl ml-5 my-4'}>
                Tasks for {monthNames[new Date().getMonth()]}{' '}
            </h1>
            <div className="mx-4 grid gap-4 text-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 ">
                {/* {boards.map((board, idx) => (
                    <Board {...board} />
                ))} */}
                <div
                    ref={UnderReviewDrop}
                    className="h-fit bg-orange-400 rounded-lg relative"
                >
                    <p
                        className={
                            'flex mt-5 text-left items-center mx-6 text-white text-xl'
                        }
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            className="mr-2 h-6 w-6"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2h-1.528A6 6 0 004 9.528V4z" />
                            <path
                                fillRule="evenodd"
                                d="M8 10a4 4 0 00-3.446 6.032l-1.261 1.26a1 1 0 101.414 1.415l1.261-1.261A4 4 0 108 10zm-2 4a2 2 0 114 0 2 2 0 01-4 0z"
                                clipRule="evenodd"
                            />
                        </svg>{' '}
                        Under Review{' '}
                        <span className="ml-2 text-orange-400 bg-white px-2 text-sm absolute right-0 mr-5 rounded-box">
                            {underReview.length}
                        </span>
                    </p>
                    <div className="h-1" />
                    {underReview.length > 0
                        ? underReview.map((t, idx) => {
                              return (
                                  <div key={idx}>
                                      <PostCard
                                          productId={productId}
                                          setList={setUnderReveiw}
                                          post={t}
                                      />
                                  </div>
                              )
                          })
                        : !URD.isOver && (
                              <p className={'my-4 text-left ml-6 text-white'}>
                                  No tasks
                              </p>
                          )}
                    {URD.canDrop && URD.isOver && (
                        <h1
                            className={`bg-orange-300 shadow-xl text-white mx-4 mb-3 text-left p-3 h-[100px] rounded-md`}
                        >
                            Drop here
                        </h1>
                    )}
                </div>
                <div
                    ref={WorkingOnDrop}
                    className={'h-fit bg-blue-500 relative rounded-lg'}
                >
                    <p
                        className={
                            'flex items-center mt-5 text-left mx-6 text-white text-xl'
                        }
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            className="mr-2 h-6 w-6"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
                        </svg>
                        Working on{' '}
                        <span className="ml-2 text-blue-500 bg-white px-2 text-sm absolute right-0 mr-5 rounded-box">
                            {workingOn.length}
                        </span>
                    </p>
                    <div className="h-1" />
                    {workingOn.length > 0
                        ? workingOn.map((t, idx) => {
                              return (
                                  <div key={idx}>
                                      <PostCard
                                          productId={productId}
                                          setList={setWorkingOn}
                                          post={t}
                                      />
                                  </div>
                              )
                          })
                        : !WOD.isOver && (
                              <p className={' my-4 text-left ml-6 text-white'}>
                                  No tasks
                              </p>
                          )}

                    {WOD.canDrop && WOD.isOver && (
                        <h1 className="bg-blue-400 shadow-xl text-white  mx-4 mb-3 text-left p-3 h-[100px] rounded-md">
                            Drop here
                        </h1>
                    )}
                </div>
                <div
                    ref={DoneDrop}
                    className="h-fit relative bg-green-500 rounded-lg md:col-span-2 lg:col-span-1"
                >
                    <p
                        className={
                            'flex items-center mt-5 text-left mx-6 text-white text-xl'
                        }
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            className="mr-2 h-6 w-6"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path
                                fillRule="evenodd"
                                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                                clipRule="evenodd"
                            />
                        </svg>{' '}
                        Done{' '}
                        <span className="ml-2 text-green-500 bg-white px-2 text-sm absolute right-0 mr-5 rounded-box">
                            {done.length}
                        </span>
                    </p>
                    <div className="">
                        <div className="h-1" />
                        {done.length > 0
                            ? done.map((t, idx) => {
                                  return (
                                      <div key={idx}>
                                          <PostCard
                                              productId={productId}
                                              setList={setDone}
                                              post={t}
                                          />
                                      </div>
                                  )
                              })
                            : !DD.isOver && (
                                  <p
                                      className={
                                          ' my-4 text-left ml-6 text-white'
                                      }
                                  >
                                      No tasks
                                  </p>
                              )}
                    </div>
                    {DD.canDrop && DD.isOver && (
                        <h1 className="bg-green-400 shadow-xl text-white  mx-4 mb-3 text-left p-3 h-[100px] rounded-md">
                            Drop here
                        </h1>
                    )}
                </div>
            </div>
        </div>
    )
}

ProductHome.Layout = HomeLayout
