import {Product} from '../../state/interfaces/'
export default function ProductCard({product} :{product: Product}) {
    return (
        <div className="flex items-center mb-4">
            {product.images && <img src={product.images[0]} className="h-16 w-16 object-contain"/> }
            <h1 className="text-xl ml-2">{product.name}</h1>
        </div>
    )
}