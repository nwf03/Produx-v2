import { useRouter } from "next/router";
import { useGetProductsQuery } from "../../../src/state/reducers/api";
import { Product } from "../../../src/state/interfaces";
import Sidebar from "../../../src/components/Products/Home/Sidebar";
import Posts from "../../../src/components/Products/Home/Posts";
import { useState } from "react";
import SmallNavBar from '../../../src/components/Products/Home/SMScreens/NavBar'
import { useAppSelector } from "../../../src/state/hooks";
import ChannelFilters from "../../../src/components/Products/Home/SMScreens/ChannelFilters";
import NavBar from "../../../src/components/Home/NavBar";
export default function ProductHome() {
  const router = useRouter();
  const channel = useAppSelector((state) => state.channel.channel);
  const { name } = router.query;
  const { data, isLoading, error } = useGetProductsQuery({
    name: name as string,
    page: 0,
  });
  const product = data as Product;
  return (
      <div>
        {/*<NavBar/>*/}
    <div className="grid grid-cols-4 lg:grid-cols-5 items-center h-screen">
      {product && (
          <div>
          <div className="hidden md:block ">
            <Sidebar product={product} channel={channel} />
          </div>
            <div className={"md:hidden"}>
                <SmallNavBar product={product}/>
            </div>
          </div>

      )}
      {error && (
        <div className="flex w-screen h-screen items-center justify-center">
          <h1 className="font-bold text-4xl">
            Product <span className="text-red-400">&apos;{name}&apos;</span> Not
            Found
          </h1>
        </div>
      )}
      {isLoading && "Loading..."}
      <div className="bg-white col-span-full md:col-span-3 lg:col-span-3 h-screen overflow-y-scroll overflow-x-hidden">
        {<Posts name={name as string} channel={channel} />}
      </div>
      <div className="hidden lg:block bg-amber-500 h-screen">
        <h1>No content for now</h1>
      </div>
        <div className={"md:hidden fixed bottom-0 bg-red-500 w-screen "}>
            <ChannelFilters channel={channel}/>
        </div>
    </div>
      </div>
  );
}
