import { LegacyRef, useState, useEffect, useRef } from 'react'
import useWebSocket from 'react-use-websocket'
import moment from 'moment'
import { Avatar, Button, Input } from '@nextui-org/react'
import { useLazyGetChatMessagesQuery } from '../../../../state/reducers/api'
export default function QuickQuestions({ productId }: { productId: number }) {
    const [getMessages, { data, isLoading }] = useLazyGetChatMessagesQuery()
    const hasMore = data ? data.hasMore : false
    const [lastId, setLastId] = useState(0)
    const socketUrl = `ws://localhost:8000/ws/${productId}`
    const { lastMessage, readyState, sendMessage } = useWebSocket(socketUrl, {
        onMessage: ({ data }) => {
            const msg = JSON.parse(data)
            if (msg.type === 'chat') {
                console.log('new message: ', msg)
                lastMsg.current?.scrollIntoView()
                setMsgs([...msgs, msg])
                return
            } else if (msg.type === 'usersList') {
                console.log('USERS LIST: ', msg.users)
                setUsers([...new Set(msg.users)])
                return
            }
            setSysMsg(msg)
        },
        onOpen: () => {
            const token = localStorage.getItem('token')
            sendMessage(token as string)
        },
        retryOnError: true,
    })
    const [users, setUsers] = useState([])
    const [msgs, setMsgs] = useState<any>([])
    // ? gets messages history from the server
    // useEffect(() => {
    //     getMessages({
    //         productId,
    //         lastId,
    //     })
    // }, [lastId])
    useEffect(() => {
        if (data) {
            setMsgs((m) => [...m, ...data.messages])
        }
    }, [data])
    const [msg, setMsg] = useState<string>('')
    const [sysMsg, setSysMsg] = useState<any>(null)
    const msgListDiv = useRef<HTMLDivElement>(null)
    const lastMsg = useRef<HTMLDivElement>(null)
    const sendMsg = (e: any) => {
        e.preventDefault()
        if (msg.trim() !== '') {
            sendMessage(msg)
            setMsg('')
        } else {
            alert('Please enter a message')
        }
    }

    useEffect(() => {
        setTimeout(() => {
            setSysMsg(null)
        }, 5000)
        return () => {
            clearTimeout()
        }
    }, [sysMsg])

    return (
        <div className={'overflow-auto w-full mt-4 mx-10'}>
            <div className="grid grid-rows-2">
                <h1 className={'text-3xl font-bold'}>Ask a question</h1>
                {users && (
                    <Avatar.Group className="ml-2 z-50" count={users.length}>
                        {users.map((user, index) => (
                            <Avatar
                                key={index}
                                size="md"
                                className="z-0"
                                pointer
                                src={user.pfp}
                                bordered
                                stacked
                            />
                        ))}
                    </Avatar.Group>
                )}
            </div>
            <div
                className={
                    'fixed bg-white w-[86vw] md:w-[64vw] lg:w-[50vw] bottom-0'
                }
            >
                <form className={'flex mb-5 justify-center items-center'}>
                    <Input
                        fullWidth
                        value={msg}
                        onChange={(e) => setMsg(e.target.value)}
                        type={'text'}
                        placeholder={'send a message'}
                    />
                    <Button auto className="ml-4 " onClick={sendMsg}>
                        Send
                    </Button>
                </form>
            </div>
            {sysMsg && (
                <div>
                    <h1 className={'text-red-400'}>{sysMsg.message}</h1>
                </div>
            )}
            <br />
            <div
                ref={msgListDiv}
                className={`fixed bottom-20 w-full overflow-auto max-h-[80vh] lg:w-[54vw]`}
            >
                {msgs &&
                    msgs.map((m: message, idx: number) => {
                        return (
                            <div key={idx}>
                                <MessageCard msg={m} />
                            </div>
                        )
                    })}
                <div ref={lastMsg}></div>
            </div>
        </div>
    )
}
interface message {
    user: messageUser
    message: string
    type: string
    CreatedAt: string
}
interface messageUser {
    userId: number
    username: string
    pfp: string
}
function MessageCard({ msg }: { msg: message }) {
    return (
        <div className="bg-white  rounded-3xl min-w-[20vw] items-center ">
            <div className="flex items-center">
                {msg.user.pfp && (
                    <Avatar squared bordered src={msg.user.pfp} size="md" />
                )}
                <p className={'ml-2'}>{msg.user.username}</p>
                <p className="text-sm text-gray-400 ml-3">
                    {moment(new Date(msg.CreatedAt)).fromNow()}
                </p>
            </div>
            <div>
                <div className="">
                    <span className={'break-all'}>{msg.message}</span>
                </div>
            </div>
            <div className={'flex justify-start  '}>
                <div className={'divider w-full md:w-8/12 lg:w-full h-0'} />
            </div>
        </div>
    )
}
