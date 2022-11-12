import { useRouter } from "next/router";
import {
  useLazyCheckIfProductFollowedQuery,
  useLazyGetProductsQuery,
} from "../../../../src/state/reducers/api";
import { useEffect, useState } from "react";
import PostDetails from "../../../../src/components/Products/Home/Posts/PostDetails";
import HomeLayout from "../../../../src/components/Products/Home/HomeLayout";
import LoadingSpinner from "../../../../src/components/LoadingSpinner";
export default function PostView() {
  const router = useRouter();
  const { name, postId, channel } = router.query;
  const [checkIfFollowed] = useLazyCheckIfProductFollowedQuery();
  const [getProduct, { data, isLoading, error }] = useLazyGetProductsQuery();
  const [showPost, setShowPost] = useState(false);
  useEffect(() => {
    if (name && postId && channel) {
      getProduct({
        name: name as string,
        page: 0,
      });
      setShowPost(true);
    }
  }, [channel, getProduct, name, postId]);
  const [showPrivate, setShowPrivate] = useState(false);
  useEffect(() => {
    if (data && data.product.private) {
      checkIfFollowed(data.product.ID).then((res) => {
        if (res.data.followed) {
          setShowPrivate(false);
        } else {
          setShowPrivate(true);
        }
      });
    }
  }, [data]);
  return (
    <div>
      {data && (
        <div className="bg-white col-span-4 md:col-span-3 lg:col-span-3 h-screen overflow-y-scroll overflow-x-hidden">
          {showPost && <PostDetails />}
        </div>
      )}
      {isLoading && <LoadingSpinner />}
    </div>
  );
}

PostView.Layout = HomeLayout;
