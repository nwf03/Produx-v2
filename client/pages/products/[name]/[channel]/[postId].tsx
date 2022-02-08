import { useRouter } from "next/router";
import {
  useLazyGetProductsQuery,
  useLazyCheckIfProductFollowedQuery,
  useLazyGetPostDetailsQuery,
} from "../../../../src/state/reducers/api";
import { Post, Product, ProductUser } from "../../../../src/state/interfaces";
import Sidebar from "../../../../src/components/Products/Home/Sidebar";
import { useState, useEffect } from "react";
import SmallNavBar from "../../../../src/components/Products/Home/SMScreens/NavBar";
import { useAppSelector } from "../../../../src/state/hooks";
import ChannelFilters from "../../../../src/components/Products/Home/SMScreens/ChannelFilters";
import ProductUsers from "../../../../src/components/Products/Home/ProductUsers";
import Head from "next/head";
import AccessPrivateProduct from "../../../../src/components/Products/AccessPrivateProduct";
import PostDetails from "../../../../src/components/Products/PostDetails";
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
      console.log("product id: ", JSON.stringify(data));
      checkIfFollowed(data.product.ID).then((res) => {
        console.log(res);
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
      {isLoading && <LoadingSpinner/>}
    </div>
  );
}

PostView.Layout = HomeLayout;
