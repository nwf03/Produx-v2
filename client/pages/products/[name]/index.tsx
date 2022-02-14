import { useRouter } from "next/router";
import HomeNavBar from "../../../src/components/Home/NavBar";
import {
  useLazyGetProductsQuery,
  useLazyCheckIfProductFollowedQuery,
  useLazyGetProductInfoQuery,
} from "../../../src/state/reducers/api";
import { Product, ProductUser } from "../../../src/state/interfaces";
import Sidebar from "../../../src/components/Products/Home/Sidebar";
import Posts from "../../../src/components/Products/Home/Posts";
import { useState, useEffect } from "react";
import SmallNavBar from "../../../src/components/Products/Home/SMScreens/NavBar";
import { useAppSelector } from "../../../src/state/hooks";
import ChannelFilters from "../../../src/components/Products/Home/SMScreens/ChannelFilters";
import NavBar from "../../../src/components/Home/NavBar";
import ProductUsers from "../../../src/components/Products/Home/ProductUsers";
import Head from "next/head";
import HomeLayout from "../../../src/components/Products/Home/HomeLayout";
import AccessPrivateProduct from "../../../src/components/Products/AccessPrivateProduct";
interface ProductResponse {
  product: Product;
  users: ProductUser[];
}
export default function ProductHome() {
  const router = useRouter();
  const channel = useAppSelector((state) => state.channel.channel);
  const { name } = router.query;
  console.log(name);
  const [checkIfFollowed] = useLazyCheckIfProductFollowedQuery();
  const [getProduct, { data, isLoading, error }] = useLazyGetProductInfoQuery();

  useEffect(() => {
    if (name) {
      getProduct(name as string);
    }
  }, [getProduct, name]);
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
    <div className="bg-green-200 h-screen m-12 rounded-box">
      <h1 className="font-bold text-3xl mx-10 my-7">Homepage</h1>
    </div>
  );
}

ProductHome.Layout = HomeLayout;
