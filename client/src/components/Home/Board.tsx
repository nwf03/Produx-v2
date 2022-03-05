import PostCard from '../Products/Home/homepage/Post'
export default function Board({
    name,
    bgcolor,
    posts,
    ref,
    icon,
    setPosts,
    dropInfo,
    productId,
}: {
    name: string
    bgcolor: string
    posts: any
    ref: any
    icon: any
    setPosts: any
    dropInfo: any
    productId: number
}) {
    return (
        <div
            ref={ref}
            className={`h-fit relative ${bgcolor} rounded-lg md:col-span-2 lg:col-span-1`}
        >
            <p
                className={
                    'flex items-center mt-5 text-left mx-6 text-white text-xl'
                }
            >
                {icon}
                {name}
                <span className="ml-2 text-green-500 bg-white px-2 text-sm absolute right-0 mr-5 rounded-box">
                    {posts.length}
                </span>
            </p>
            <div className="">
                <div className="h-1" />
                {posts.length > 0
                    ? posts.map((t, idx) => {
                          return (
                              <div key={idx}>
                                  <PostCard
                                      productId={productId}
                                      setList={setPosts}
                                      post={t}
                                  />
                              </div>
                          )
                      })
                    : !dropInfo.isOver && (
                          <p className={' my-4 text-left ml-6 text-white'}>
                              No tasks
                          </p>
                      )}
            </div>
            {dropInfo.canDrop && dropInfo.isOver && (
                <h1 className="bg-green-400 shadow-xl text-white  mx-4 mb-3 text-left p-3 h-[100px] rounded-md">
                    Drop here
                </h1>
            )}
        </div>
    )
}
