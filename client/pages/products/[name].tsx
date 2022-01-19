import { useRouter } from "next/router";
import { useGetProductsQuery } from "../../src/state/reducers/api";
import { Product } from "../../src/state/interfaces";
import Sidebar from "../../src/components/Products/Home/Sidebar";
import Posts from "../../src/components/Products/Home/Posts";
import { useState } from "react";
import { useAppSelector } from "../../src/state/hooks";
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
    <div className="grid grid-cols-4 lg:grid-cols-5 items-center h-screen">
      {product && (
        <div className="hidden md:block ">
          <Sidebar product={product} channel={channel} />
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
      <div className="bg-red-300 col-span-full md:col-span-3 lg:col-span-3 h-screen overflow-y-scroll overflow-x-hidden">
        {<Posts name={name as string} channel={channel} />}
      </div>
      <div className="hidden lg:block bg-amber-500 h-screen">
        <h1>Yessir</h1>
      </div>
    </div>
  );
}
