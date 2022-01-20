import Link from "next/link";
import { useState } from "react";
import { ProductResponse } from "../../state/reducers/api";
import { useGetProductsQuery } from "../../state/reducers/api";
export default function SearchProducts({ name }: { name: string }) {
  const { data, isLoading, error } = useGetProductsQuery({
    name: name,
    page: 1,
  });
  const { products } = data ? (data as ProductResponse) : { products: [] };

  return (
    <div className=" w-10/12 md:w-[31vw] absolute top-20 ml-[-11px] min-w-[200px] text-left bg-gray-300 rounded-box">
      <div className="">
        {isLoading && "Loading..."}
        {data ? (
          products.map((product, idx) => {
            return (
              <Link href={`/products/${product.name}`} key={idx}>
                <a>
                  <div className="bg-blue-400 m-2 rounded-box p-5 text-white  flex items-center">
                    {product.images ? (
                      <img
                        src={product.images[0]}
                        className="h-16 w-16 rounded-xl"
                      />
                    ) : (
                      <div className="h-16 w-16 bg-red-500 rounded-xl"></div>
                    )}
                    <div className="ml-2 block">
                      <h3 className="font-bold">{product.name}</h3>
                      <p className="">{product.description}</p>
                    </div>
                  </div>
                </a>
              </Link>
            );
          })
        ) : (
          <div className="bg-red-400 m-2 rounded-box p-5 text-white flex items-center">
            No Products Found
          </div>
        )}
      </div>
    </div>
  );
}
