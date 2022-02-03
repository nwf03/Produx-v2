import { useRouter } from "next/router";
import HomeNavBar from "../../../src/components/Home/NavBar";
import { useLazyGetProductsQuery, useLazyCheckIfProductFollowedQuery } from "../../../src/state/reducers/api";
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
  const [checkIfFollowed] = useLazyCheckIfProductFollowedQuery()
  const [getProduct, { data, isLoading, error }] = useLazyGetProductsQuery();

  useEffect(() => {
    if (name) {
      getProduct({
        name: name as string,
        page: 0,
      });
    }
  }, [getProduct, name]);
  const [showPrivate, setShowPrivate]= useState(false)
  useEffect(() => {
    if (data && data.product.private){
      console.log("product id: ", JSON.stringify(data))
     checkIfFollowed(data.product.ID).then(res => {
       console.log(res)
       if (res.data.followed){
         setShowPrivate(false)
       }else{
         setShowPrivate(true)
       }
    })
    };
  }, [data])
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
          {<Posts name={name as string} channel={channel} />}
        </div>
        <div className="hidden lg:block h-screen">
          <ProductUsers
            users={data ? data.users : []}
            productId={data ? data.product.ID : 0}
            productUserID={data ? data.product.userID : 0}
          />
        </div>
        <div className="md:hidden fixed bottom-0 bg-blue-600 w-screen">
          <ChannelFilters channel={channel} />
          </div>
      </div>
        )}
      {data &&
      <AccessPrivateProduct show={showPrivate} setShow={setShowPrivate} productName={data.product.name} productIcon={data.product.images ? data.product.images[0] : null}/>}
      {error && (
          <div className="flex w-screen h-screen items-center justify-center">
            <h1 className="font-bold text-4xl">
              Product <span className="text-red-400">&apos;{name}&apos;</span>{" "}
              Not Found
            </h1>
          </div>
        )}
        {isLoading && "Loading..."}
    </div>
  );
}
