import { useRouter } from "next/router";
import HomeNavBar from "../../../src/components/Home/NavBar";
import {
  useLazyGetProductsQuery,
  useLazyCheckIfProductFollowedQuery,
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
  const [getProduct, { data, isLoading, error }] = useLazyGetProductsQuery();

  useEffect(() => {
    if (name) {
      getProduct({
        name: name as string,
        page: 0,
      });
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
    <div>
      {data && (
        <div className="bg-white col-span-4 md:col-span-3 lg:col-span-3 h-screen overflow-y-scroll overflow-x-hidden">
          {<Posts name={name as string} channel={channel} />}
        </div>
      )}
    </div>
  );
}

ProductHome.Layout = HomeLayout;
