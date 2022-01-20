import { Post } from "../../state/interfaces";
export default function PostView({
  data,
  channel,
  color,
  showDivider
}: {
  data: Post;
  channel: string;
  color: string;
  showDivider: boolean
}) {
  return (
    <div className={!showDivider ?"mb-6" :""}>
      <div className="bg-white p-5 rounded-3xl  mb-4 w-auto items-center mx-12 mt-4 md:mt-0">
      <div className="grid-cols-3 flex align-bottom ">
        <div className="flex items-center">
          {data.user.pfp && (
            <img
              src={data.user.pfp}
              className="w-12  rounded-box align-middle"
            />
          )}
          <span className="text-sm pl-2">{data.product.name}</span>
          <span className="font-bold mr-6"> - {data.title}</span>
        </div>
        <span style={{backgroundColor: color || "white"}} className={`text-black  rounded-[10px]  px-2 py-2 text-xs text-center my-auto ml-auto`}>
          {channel.substring(0, channel.length - 1)}
        </span>
      </div>
      <div>
        <div className="mt-3">
          <span className={"break-words"}>{data.description}</span>
        </div>
      </div>
    </div>
        {showDivider &&
            <div className={"flex justify-center"}>
                <div className={"divider w-11/12"}/>
            </div>
        }
    </div>
  );
}