import HomeLayout from "../../../src/components/Products/Home/HomeLayout";
import { useEffect, useState, useRef } from "react";
import {
  useGetProductDayPostCountQuery,
    useGetPostsBoardQuery
} from "../../../src/state/reducers/api";
import Post from "../../../src/components/Products/Home/homepage/Post";
import { useAppDispatch } from "../../../src/state/hooks";
import { setChannel } from "../../../src/state/reducers/channelSlice";
import {useDrop} from "react-dnd";
import {useAddToPostsBoardMutation, useRemoveFromPostsBoardMutation} from "../../../src/state/reducers/api";

export default function ProductHome({ productId }: { productId: number }) {
  const refreshes = useRef(0)
  const [addPostToBoard] = useAddToPostsBoardMutation()
  const [URD, UnderReviewDrop] = useDrop(() => ({
    accept: "post",
    drop: async(item: any) => {
      addItemToBoard(item.post, "under-review")
      await addPostToBoard({
        productId,
        postId: item.post.ID,
        field: 'under-review'
      })

    },
  }));


  const [WOD, WorkingOnDrop] = useDrop(() => ({
    accept: "post",
    drop: async(item: any) => {
      addItemToBoard(item.post, "working-on")
      await addPostToBoard({
        productId,
        postId: item.post.ID,
        field: 'working-on'
      })
    },
  }));
  const [DD, DoneDrop] = useDrop(() => ({
    accept: "post",
    drop: async(item: any) => {
      addItemToBoard(item.post, "done")
      await addPostToBoard({
        productId,
        postId: item.post.ID,
        field: 'done'
      })
    },
  }));
  const { data, isLoading, error } = useGetProductDayPostCountQuery(productId);
  const posts = useGetPostsBoardQuery({
    productId,
  });
  const colors = [
    "from-red-500 to-orange-300",
    "from-orange-500 to-yellow-400",
    "from-yellow-400 to-green-400",
    "from-green-400 to-blue-400",
    "from-green-400 to-blue-400",
    "from-green-400 to-blue-400",
    "from-green-400 to-blue-400",
  ];
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(setChannel("Home"));
  }, []);
  useEffect(() => {
    if (posts.data){
      setUnderReveiw(posts.data.underReview)
      setWorkingOn(posts.data.workingOn)
      setDone(posts.data.done)
    }
  }, [posts.data]);

  const [underReview, setUnderReveiw] = useState<any[]>(posts.data ? posts.data.underReview : []);
  const [workingOn, setWorkingOn] = useState<any[]>(posts.data ? posts.data.workingOn : []);
  const [done, setDone] = useState<any[]>(posts.data ? posts.data.done : []);

  const addItemToBoard = (item: any, field: string) => {
    switch (field) {
      case 'under-review':
          for (const p of workingOn){
            if (p.ID === item.ID){
              break
            }
          }
          setUnderReveiw(items => [...items, item])
        break;
      case 'working-on':
          for (const p of done){
            if (p.ID === item.ID){
              break
            }
          }
          setWorkingOn(items => [...items, item])
        break;
      case 'done':
        for (const p of underReview){
          if (p.ID === item.ID){
            break
          }
        }
        setDone(items => [...items, item])
        break;
    }
    }
  const monthNames = ["January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
  ];
  return (
    <div className={"mt-4 ml-4"}>
      <h1 className={"font-bold text-2xl"}>Today&apos;s Numbers</h1>
      <div className={"grid grid-cols-2 md:grid-cols-3 lg:grid-cols-3 gap-y-3 mt-4"}>
        {data &&
          Object.keys(data).map((c, idx) => (
            <div
              className={`h-14 mx-4  font-bold flex justify-center items-center py-3 px-9 rounded-2xl bg-gradient-to-r ${colors[idx]}  text-white mx-4`}
            >
              <span className="font-bold text-center text-md">
                {data[c]} {data[c] == 1 ? c.substring(0, c.length - 1) : c}
              </span>
            </div>
          ))}
      </div>
      <div className="divider w-11/12 mx-auto" />
      <h1 className={"font-bold text-2xl ml-5 my-4"}>Tasks for {monthNames[new Date().getMonth()]} </h1>
      <div className="mx-4 grid gap-4 text-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 ">
        <div ref={UnderReviewDrop} className="h-fit bg-orange-400 rounded-lg ">
          <p className={'flex mt-5 text-left items-center mx-6 text-white text-xl'}>
            <svg xmlns="http://www.w3.org/2000/svg" className="mr-2 h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 16l2.879-2.879m0 0a3 3 0 104.243-4.242 3 3 0 00-4.243 4.242zM21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            Under Review</p>
          <div className="h-1" />
            {underReview.length > 0 ?
              underReview.map((t, idx) => {
                return (
                  <div key={idx}>
                    <Post productId={productId} setList={setUnderReveiw} post={t} />
                  </div>
                );
              }):
                <p className={' my-4 text-left ml-6 text-white'}>No tasks</p>
}
        </div>
        <div  ref={WorkingOnDrop} className={"h-fit bg-blue-500 rounded-lg"}>
          <p className={"flex items-center mt-5 text-left mx-6 text-white text-xl"}>
            <svg xmlns="http://www.w3.org/2000/svg" className="mr-2 h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
            </svg>
            Working on </p>
            <div className="h-1" />
          {workingOn.length > 0 ?
              workingOn.map((t, idx) => {
                return (
                    <div key={idx}>
                      <Post productId={productId} setList={setWorkingOn} post={t} />
                    </div>
                );
              }) :
              <p className={' my-4 text-left ml-6 text-white'}>No tasks</p>
}
        </div>
        <div ref={DoneDrop} className="h-fit bg-green-500 rounded-lg md:col-span-2 lg:col-span-1">
          <p className={'flex items-center mt-5 text-left mx-6 text-white text-xl'}>
            <svg xmlns="http://www.w3.org/2000/svg" className="mr-2 h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            Done</p>
          <div className="">
            <div className="h-1" />
            {done.length > 0 ?
                done.map((t, idx) => {
                  return (
                      <div key={idx}>
                        <Post productId={productId} setList={setDone} post={t} />
                      </div>
                  );
                }):
              <p className={' my-4 text-left ml-6 text-white'}>No tasks</p>
              }
          </div>
        </div>
      </div>
    </div>
  );
}

ProductHome.Layout = HomeLayout;
