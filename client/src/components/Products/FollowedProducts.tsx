import { useEffect } from "react";
import { useAppSelector } from "../../state/hooks";
import Product from "./Product";
export default function FollowedProducts() {
  const followed_products = useAppSelector(
    (state) => state.auth.user?.followed_products
  );

  return (
    <div className="fixed ml-10 mt-4">
      <h1 className="font-bold text-2xl mb-6">Followed Products</h1>
      <div className="mt-4">
        {followed_products &&
          followed_products.map((product, idx) => {
            return <Product product={product} key={idx} />;
          })}
      </div>
    </div>
  );
}
