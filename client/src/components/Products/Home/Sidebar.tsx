import { Product } from "../../../state/interfaces";
import { Channel } from "../../../state/interfaces";
import { useRouter } from "next/router";
import { useAppDispatch, useAppSelector } from "../../../state/hooks";
import { setChannel } from "../../../state/reducers/channelSlice";
import Link from "next/link";
import {
  useFollowProductMutation,
  useGetUserInfoQuery,
} from "../../../state/reducers/api";
import { useEffect, useState } from "react";
import { setUser } from "../../../state/reducers/auth";
export default function Sidebar({
  product,
  channel,
    userCount,
   postCount
}: {
  product: Product;
  channel: string;
  userCount: number;
  postCount: number;
}) {
  const router = useRouter();
  const dispatch = useAppDispatch();
  const channels: Channel = {
    Announcements: { icon: "🎉", color: "#FF9900" },
    Bugs: { icon: "🐞", color: "#DBFF00" },
    Suggestions: { icon: "🙏", color: "#0094FF" },
    Changelogs: { icon: "🔑", color: "#FF4D00" },
  };
  const [followProduct] = useFollowProductMutation();
  const [follow, setFollow] = useState(false);
  const { data, isLoading, error } = useGetUserInfoQuery();
  const mdBreakpoint = 1024;
  const channelName = channel.charAt(0).toUpperCase() + channel.slice(1);
  const filters = ["😱 All Posts", "🎉 Latest Posts", "🙌 Most Upvoted Posts"];
  const [width, setWidth] = useState(window.innerWidth);
  useEffect(() => {
    if (window !== undefined) {
      window.addEventListener("resize", () => {
        setWidth(window.innerWidth);
      });
    }
    return () => {
      window.removeEventListener("resize", () => {
        setWidth(window.innerWidth);
      });
    };
  }, []);

  const followHandler = async () => {
    console.log("name: ", product.name);

    await followProduct({ name: product.name, follow: !follow });
    setFollow(!follow);
  };
  useEffect(() => {
    if (data) {
      for (const p of data.followed_products!) {
        if (p.ID == product.ID) {
          setFollow(true);
          return;
        }
      }
      setFollow(false);
    }
  }, [data, product.ID]);
  // todo have a gear icon next to the product logo with a drop down menu where you can unfollow
  return (
    //  todo fix logo, name, and stat centering when screen size is md
    <div className="bg-gray-100 h-[96vh]  justify-center flex overflow-auto">
      <div className="mt-4">
        <div className="left-0 top-0 md:left-5 md:top-4 fixed dropdown">
          <div className="m-1 cursor-pointer">
            <div>
              <div
                tabIndex={0}
                className="m-1 hover:bg-gray-300 p-0 md:p-2 rounded-box"
              >
                <img src="/gearIcon.svg" className="h-5" />
              </div>
            </div>
          </div>
          <ul
            tabIndex={0}
            className="p-2 shadow mt-[-8px] dropdown-content bg-base-100 rounded-xl w-50"
          >
            <li
              className={`${
                follow ? "active:bg-red-500" : "active:bg-blue-500 "
              } p-2 m-0 cursor-pointer rounded-xl active:text-white`}
              onClick={followHandler}
            >
              <a>{follow ? "unfollow" : "follow"}</a>
            </li>
          </ul>
        </div>
        <div className="w-max mx-auto">
          {product.images && (
            <img
              src={product.images[0]}
              alt={product.name}
              className="w-[9vw] min-w-[80px] md:min-w-[120px]"
            />
          )}
        </div>
        <div className={"block text-center"}>
          <h1 className="font-bold text-3xl capitalize">{product.name}</h1>
          <p className={"max-w-[78%] break-all mx-auto"}>
            {product.description}
          </p>
        </div>
        <div className="hidden md:grid grid-cols-2 gap-0 lg:gap-4 mt-6 m-2 text-center place-items-center">
          <div className={"block lg:inline-flex gap-2"}>
            <p>🌎</p>
            <p>{userCount}</p>
          </div>
          <div className={"block lg:inline-flex gap-2"}>
            <p>💬</p>
            <p>{postCount}</p>
          </div>
        </div>
        <div className="mt-5">
          <Link
            href={{
              pathname: "/products/[name]/",
              query: { name: product.name },
            }}
          >
            <a>
              <div
                className="p-4 px-6 hover:bg-gray-200 hover:cursor-pointer mx-auto rounded-2xl my-2 w-[15vw] "
                style={{
                  backgroundColor: "Home" == channelName ? "white" : "",
                }}
              >
                <p className="text-[30px] lg:text-[calc(10px+0.5vw)] text-center lg:text-left md:text-4xl">
                  {"🏡"} {width >= mdBreakpoint && "Home"}
                </p>
              </div>
            </a>
          </Link>
          {Object.keys(channels).map((e) => {
            return (
              <Link
                key={e}
                href={{
                  pathname: "/products/[name]/[channel]",
                  query: { name: product.name, channel: e },
                }}
              >
                <a>
                  <div
                    className="p-4 px-6 hover:bg-gray-200 hover:cursor-pointer mx-auto rounded-2xl my-2 w-[15vw] "
                    style={{ backgroundColor: e == channelName ? "white" : "" }}
                  >
                    <p className="text-[30px] lg:text-[calc(10px+0.5vw)] text-center lg:text-left md:text-4xl">
                      {channels[e].icon} {width >= mdBreakpoint && e}
                    </p>
                  </div>
                </a>
              </Link>
            );
          })}
          <Link
            href={{
              pathname: "/products/[name]/questions",
              query: { name: product.name },
            }}
          >
            <a>
              <div
                className="p-4 px-6 hover:bg-gray-200 hover:cursor-pointer mx-auto rounded-2xl my-2 w-[15vw]"
                style={{
                  backgroundColor: "Questions" == channelName ? "white" : "",
                }}
              >
                <p className="text-[30px] lg:text-[calc(10px+0.5vw)] text-center lg:text-left md:text-4xl">
                  {"🤨"} {width >= mdBreakpoint && "Questions"}
                </p>
              </div>
            </a>
          </Link>
        </div>
        {/* <div className=""> */}
        {/* <h1 className={"md:ml-2 text-center"}>Filters</h1> */}
        {/* <div className="divider"></div> */}
        {/* <div>
            {filters.map((el) => {
              return (
                <h1
                  key={el}
                  className="bg-gray-200 p-4 mx-12 my-3 md:m-4 rounded-xl"
                >
                  {el}
                </h1>
              );
            })}
          </div> */}
        {/* </div> */}
      </div>
    </div>
  );
}
