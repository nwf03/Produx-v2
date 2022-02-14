import { useRouter } from "next/router";
import HomeNavBar from "../../Home/NavBar";
import {
  useLazyGetProductsQuery,
  useLazyCheckIfProductFollowedQuery,
} from "../../../state/reducers/api";
import { Product, ProductUser } from "../../../state/interfaces";
import Sidebar from "./Sidebar";
import Posts from "./Posts";
import { useState, useEffect, cloneElement } from "react";
import SmallNavBar from "./SMScreens/NavBar";
import { useAppSelector } from "../../../state/hooks";
import ChannelFilters from "./SMScreens/ChannelFilters";
import NavBar from "../../Home/NavBar";
import ProductUsers from "./ProductUsers";
import Head from "next/head";
import AccessPrivateProduct from "../AccessPrivateProduct";
import { Loading } from "@nextui-org/react";
interface ProductResponse {
  product: Product;
  users: ProductUser[];
}
import LoadingSpinner from "../../LoadingSpinner";
export default function ProductHome({ children }) {
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
      <Head>
        <title>{name} - produx</title>
        <link
          rel="icon"
          href={
            data && data.product.images
              ? data.product.images[0]
              : "/produx2.png"
          }
        />
      </Head>
      {data && (
        <div className="grid  md:grid-cols-4 lg:grid-cols-5 items-center h-screen">
          <div className="hidden md:block">
            <Sidebar product={data.product} channel={channel} />
          </div>
          <div className={"md:hidden"}>
            <SmallNavBar product={data.product} />
          </div>
          <div className="bg-white col-span-4 md:col-span-3 lg:col-span-3 h-screen overflow-y-scroll overflow-x-hidden">
            {cloneElement(children, {
              productId: data.product.ID,
            })}
          </div>
          <div className="hidden lg:block h-screen">
            <ProductUsers
              users={data ? data.users : []}
              productId={data ? data.product.ID : 0}
              productUserID={data ? data.product.userID : 0}
            />
          </div>
        </div>
      )}
      {data && (
        <AccessPrivateProduct
          show={showPrivate}
          setShow={setShowPrivate}
          productName={data.product.name}
          productIcon={data.product.images ? data.product.images[0] : null}
        />
      )}
      {error && (
        <div className="flex w-screen h-screen items-center justify-center">
          <h1 className="font-bold text-4xl">
            Product <span className="text-red-400">&apos;{name}&apos;</span> Not
            Found
          </h1>
        </div>
      )}
      {isLoading && <LoadingSpinner />}
    </div>
  );
}
