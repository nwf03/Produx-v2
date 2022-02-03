import { Product } from "../../state/interfaces/";
import Link from "next/link";
export default function ProductCard({
  product,
  showDesc,
}: {
  product: Product;
  showDesc: boolean;
}) {

  return (
    <Link href={`products/${product.name}`}>
      <a>
        <div className="hover:bg-gray-200 rounded-box flex items-center p-2 active:bg-gray-300">
          {product.images && (
            <img src={product.images[0]} className="h-16 w-16 object-contain" />
          )}
          <div className="block ml-2">
            <h1 className="text-xl ">{product.name}</h1>
            {showDesc && <p className="text-sm ">{product.description}</p>}
          </div>
        </div>
      </a>
    </Link>
  );
}
