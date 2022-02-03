import { Product } from "../../../../state/interfaces";
import { useEffect } from "react";
export default function NavBar({ product }: { product: Product }) {
  return (
    <div className={"flex items-center p-4 "}>
      {product.images && (
        <img src={product.images[0]} className={"w-14 mr-3"} />
      )}
      <div className={"block"}>
        <h1 className={"w-max font-bold text-xl"}>{product.name}</h1>
        {product.description.length <= 21 && (
          <h1 className={"w-max"}>{product.description} </h1>
        )}
      </div>
      <div className={"divider divider-vertical"}></div>
      <div
        className={
          "xsm:mx-auto xsm:relative  xsm:text-center right-0 absolute  mr-4"
        }
      >
        <div className={"flex gap-7 text-center"}>
          <p>🌎 100k</p>
          <p>🌎 100k</p>
          <p>🌎 100k</p>
        </div>
      </div>
    </div>
  );
}
