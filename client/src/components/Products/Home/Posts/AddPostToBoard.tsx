import { Button, Tooltip } from '@nextui-org/react'
import { useState } from 'react'
import {
    useAddToPostsBoardMutation,
    useRemoveFromPostsBoardMutation,
} from '../../../../state/reducers/api'
import { useDeletePostMutation } from '../../../../state/reducers/api'
import { router } from 'next/client'
import { useRouter } from 'next/router'
export interface fieldValues {
    name: 'working-on' | 'done' | 'under-review'
    color: string
    icon: any
}
type fields = Record<string, fieldValues>
export default function AddPostToBoard({
    postId,
    productId,
    postTypes,
}: {
    postId: number
    productId: number
    postTypes: string[]
}) {
    const [toolTipVisible, setToolTipVisible] = useState(false)
    const [addToPostsBoard, { isLoading }] = useAddToPostsBoardMutation()
    const [removePostFromBoard] = useRemoveFromPostsBoardMutation()
    const [deletePost, { data }] = useDeletePostMutation()
    const fields: fields = {
        'Under Review': {
            name: 'under-review',
            color: 'warning',
            icon: (
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
                        d="M8 16l2.879-2.879m0 0a3 3 0 104.243-4.242 3 3 0 00-4.243 4.242zM21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                </svg>
            ),
        },
        'Working on': {
            name: 'working-on',
            color: 'primary',
            icon: (
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
                        d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                    />
                </svg>
            ),
        },
        Done: {
            name: 'done',
            color: 'success',
            icon: (
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
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                </svg>
            ),
        },
    }
    const router = useRouter()
    const handleAdd = async (field: 'working-on' | 'done' | 'under-review') => {
        if (postTypes.includes(field)) {
            await removePostFromBoard({
                productId,
                postId,
                field,
                refetchBoard: true,
            })
        } else {
            await addToPostsBoard({
                productId,
                postId,
                field,
                refetchBoard: true,
            })
            setToolTipVisible(false)
        }
    }
    const handleDelete = async () => {
        await deletePost({
            productId,
            postId,
            field: postTypes[0],
        })
        setTimeout(() => {
            router.back()
        }, 500)
    }
    return (
        <div className={'py-2 text-white'}>
            <h1 className="ml-2"> Add tags </h1>
            <hr className={'w-11/12 mb-3 mt-1 mx-auto'} />
            {Object.keys(fields).map((field, idx) => {
                return (
                    <Button
                        key={idx}
                        onClick={() => handleAdd(fields[field].name)}
                        icon={fields[field].icon}
                        className={'mb-4'}
                        color={fields[field].color}
                    >
                        {field}
                    </Button>
                )
            })}
            {/*<Button icon={*/}
            {/*    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">*/}
            {/*        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 16l2.879-2.879m0 0a3 3 0 104.243-4.242 3 3 0 00-4.243 4.242zM21 12a9 9 0 11-18 0 9 9 0 0118 0z" />*/}
            {/*    </svg>*/}
            {/*}*/}
            {/*        className={"mb-4"}  color={"warning"} >*/}
            {/*   Under Review*/}
            {/*</Button>*/}
            {/* <Button icon={*/}
            {/*     <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">*/}
            {/*         <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />*/}
            {/*     </svg>*/}
            {/* }*/}
            {/*         className={"mb-4"}  >*/}
            {/*     Working on*/}
            {/* </Button>*/}
            {/* <Button icon={*/}
            {/*     <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">*/}
            {/*         <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />*/}
            {/*     </svg>*/}
            {/* }  color={"success"} >*/}
            {/*    Done*/}
            {/* </Button>*/}

            <hr className={'w-11/12 mb-3 mt-3 mx-auto'} />
            <Tooltip
                visible={toolTipVisible}
                hideArrow
                trigger="click"
                onClick={() => setToolTipVisible(true)}
                content={
                    <div className={'p-2 w-[220px]'}>
                        <h1 className={'mb-2 font-bold text-md'}>
                            Confirm Delete
                        </h1>
                        <div
                            className={'grid grid-cols-1 md:grid-cols-2 gap-2'}
                        >
                            <Button
                                flat
                                onClick={() => setToolTipVisible(false)}
                                auto
                                color={'primary'}
                            >
                                Cancel
                            </Button>
                            <Button
                                onClick={() => handleDelete()}
                                flat
                                auto
                                color={'error'}
                            >
                                {data ? 'Deleted!' : 'Delete'}
                            </Button>
                        </div>
                    </div>
                }
            >
                <Button
                    icon={
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
                                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
                            />
                        </svg>
                    }
                    color={'error'}
                >
                    Delete
                </Button>
            </Tooltip>
        </div>
    )
}
