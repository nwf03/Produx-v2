import {
  useCallback,
  useEffect,
  useLayoutEffect,
  useRef,
  useState,
} from "react";
import useWebSocket from "react-use-websocket";
import moment from "moment";
import { Avatar, Button, Input, Loading } from "@nextui-org/react";
import { useLazyGetChatMessagesQuery } from "../../../../state/reducers/api";
import { setOnlineUsers } from "../../../../state/reducers/productSlice";
import { useAppDispatch } from "../../../../state/hooks";
// TODO fix re-rendering issue when sending message
export default function QuickQuestions({ productId }: { productId: number }) {
  const [getMessages, { data, isLoading }] = useLazyGetChatMessagesQuery();
  const hasMore = data ? data.hasMore : false;
  const lastMsg = useRef<HTMLDivElement>(null);
  const observer = useRef<HTMLDivElement>(null);
  const dispath = useAppDispatch();
  const [lastId, setLastId] = useState(0);
  const firstMsgRef = useCallback(
    (node) => {
      if (isLoading) return;
      if (observer.current) observer.current.disconnect();
      observer.current = new IntersectionObserver((entries) => {
        if (entries[0].isIntersecting) {
          if (hasMore && data && data.lastId) {
            setLastId(data.lastId);
          }
        }
      });
      if (node) observer.current.observe(node);
    },
    [isLoading, hasMore],
  );
  const socketUrl = `ws://localhost:8000/ws/${productId}`;
  const { sendMessage } = useWebSocket(socketUrl, {
    onMessage: ({ data }) => {
      const msg = JSON.parse(data);
      if (msg.type === "chat") {
        setMsgs([...msgs, msg]);
        return;
      } else if (msg.type === "usersList") {
        dispath(setOnlineUsers(msg.users));
        return;
      }
      setSysMsg(msg);
    },
    onOpen: () => {
      const token = localStorage.getItem("token");
      sendMessage(token as string);
    },
    retryOnError: true,
  });
  const [msgs, setMsgs] = useState<any>([]);
  useEffect(() => {
    getMessages({
      productId,
      lastId,
    });
  }, [lastId]);
  useEffect(() => {
    if (data) {
      const ms = [...data.messages];
      ms.reverse();
      setMsgs((m) => [...ms, ...m]);
    }
  }, [data]);
  const [msg, setMsg] = useState<string>("");
  const [sysMsg, setSysMsg] = useState<any>(null);
  const msgListDiv = useRef<HTMLDivElement>(null);
  const sendMsg = (e: any) => {
    e.preventDefault();
    if (msg.trim() !== "") {
      sendMessage(msg);
      setMsg("");
    } else {
      alert("Please enter a message");
    }
  };

  useEffect(() => {
    if (sysMsg) {
      setTimeout(() => {
        setSysMsg(null);
      }, 5000);
    } else {
      clearTimeout();
    }
    return () => {
      clearTimeout();
    };
  }, [sysMsg]);

  useLayoutEffect(() => {
    lastMsg.current?.scrollIntoView();
  }, [msgs]);

  useEffect(() => {
    return () => {
      dispath(setOnlineUsers([]));
    };
  }, [dispath]);

  return (
    <div className={"w-full"}>
      <div className="sticky top-0  bg-white z-50 pt-4">
        <h1 className={"px-4 text-3xl font-bold"}>Ask a question</h1>
        <div className="divider h-0 bg-red-400" />
      </div>
      <div className={`fixed z-50 md:w-[64vw] bottom-0`}>
        <form
          className={"grid grid-cols-10 pb-5 w-full justify-start items-center mr-12"}
        >
          <div className="ml-4 lg:col-span-8 md:col-span-9 col-span-8">
            <Input
              fullWidth
              value={msg}
              onChange={(e) => setMsg(e.target.value)}
              type={"text"}
              placeholder={"send a message"}
            />
          </div>
          <Button auto className="ml-4 " onClick={sendMsg}>
            Send
          </Button>
        </form>
      </div>
      {sysMsg && (
        <div>
          <h1 className={"text-red-400"}>{sysMsg.message}</h1>
        </div>
      )}
      <div
        ref={msgListDiv}
        className={`relative bottom-5 mb-16 mx-4 mt-10`}
      >
        {isLoading && <Loading className="mx-auto w-full" />}
        {!hasMore && msgs.length >= 15 && (
          <p className="text-orange-400 text-sm text-center">
            You reached the end of the chat history
          </p>
        )}
        {msgs.length > 0
          ? (
            msgs.map((m: message, idx: number) => {
              return (
                <div
                  key={idx}
                  ref={idx == msgs.length - 1
                    ? lastMsg
                    : idx == 0
                    ? firstMsgRef
                    : null}
                >
                  <MessageCard msg={m} />
                </div>
              );
            })
          )
          : (
            <p className="ml-2 text-gray-400 mb-2">
              Start a conversation
            </p>
          )}
        <div ref={lastMsg}></div>
      </div>
    </div>
  );
}
interface message {
  user: messageUser;
  message: string;
  type: string;
  CreatedAt: string;
}
interface messageUser {
  userId: number;
  name: string;
  pfp: string;
}
function MessageCard({ msg }: { msg: message }) {
  return (
    <div className="bg-white rounded-3xl items-center mb-5">
      <div className="flex ">
        {msg.user.pfp && (
          <Avatar
            className="z-0"
            squared
            bordered
            src={msg.user.pfp}
            size="lg"
          />
        )}
        <div className="ml-2">
          <div className="flex text-sm items-center">
            <p className={"font-bold"}>{msg.user.name}</p>
            <p className="text-xs text-gray-400 ml-3">
              {moment(new Date(msg.CreatedAt)).fromNow()}
            </p>
          </div>{" "}
          <span className={"break-all"}>{msg.message}</span>
        </div>
      </div>
    </div>
  );
}
