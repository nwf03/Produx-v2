import { Product } from "../../../state/interfaces";
import { Channel } from "../../../state/interfaces";
import { useAppDispatch } from "../../../state/hooks";
import { setChannel } from "../../../state/reducers/channelSlice";

import {useEffect, useState} from "react";

export default function Sidebar({
  product,
  channel,
}: {
  product: Product;
  channel: string;
}) {
  const dispatch = useAppDispatch();
  const channels: Channel = {
    Announcements: { icon: "🎉", color: "#FF9900" },
    Bugs: { icon: "🐞", color: "#DBFF00" },
    Suggestions: { icon: "🙏", color: "#0094FF" },
    Changelogs: { icon: "🔑", color: "#FF4D00" },
  };
  const mdBreakpoint = 1024;
  const channelName = channel.charAt(0).toUpperCase() + channel.slice(1);
  const filters = ["😱 All Posts", "🎉 Latest Posts", "🙌 Most Upvoted Posts"]
  const [width,setWidth] = useState(window.innerWidth)
  useEffect(()=>{

    window.addEventListener('resize', ()=>{
      setWidth(window.innerWidth)
      console.log(`width: ${window.innerWidth}`)
    })
    return ()=>{
      window.removeEventListener('resize', ()=>{
        setWidth(window.innerWidth)
      })
    }
  }, [])
  return (
    //  todo fix logo, name, and stat centering when screen size is md
    <div className="bg-gray-100 h-screen justify-center flex overflow-auto">
      <div className="mt-4">
        <div className="items-center justify-center w-max  text-center">
          {product.images && (
            <img
              src={product.images[0]}
              alt={product.name}
              className="w-[10vw]"
            />
          )}
          <h1 className="font-bold text-3xl capitalize">{product.name}</h1>
        </div>
        <div className="grid grid-row-3 lg:grid-cols-3 gap-4 mt-6 m-2">
          <p>🌎 100k</p>
          <p>🌎 100k</p>
          <p>🌎 100k</p>
        </div>
        <div>
          {Object.keys(channels).map((e) => {
            return (
              <div
                key={e}
                // todo make selected channel opacity light gray
                className="p-4 px-6 hover:bg-gray-200 hover:cursor-pointer rounded-2xl my-2 mx-7 lg:mx-0"
                style={{ backgroundColor: e == channelName ? "white" : "" }}
                onClick={() => dispatch(setChannel(e))}
              >
                <p className="lg:text-[calc(10px+0.5vw)] text-center lg:text-left text-4xl">
                  {channels[e].icon} { width >= mdBreakpoint && e}
                </p>
              </div>
            );
          })}
        </div>
        <h1>Filters</h1>
        <div className="divider"></div>
        <div>
          {filters.map(el=>{
            return <h1 key={el} className="bg-gray-200 p-4 m-4 rounded-xl">{el}</h1>
          }
          )}
        </div>
      </div>
    </div>
  );
}
