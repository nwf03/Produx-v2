import ProductCard from './Product'
import { useGetTopProductsQuery } from '../../state/reducers/api'

export default function TopProducts() {
    const { data, error, isLoading } = useGetTopProductsQuery()
    return (
        <div className="">
            <h1 className="bg-gray-200 p-6 flex justify-center mr-2 rounded-box  ">
                {isLoading && 'Loading....'}
                {data && data.length > 0 ? (
                    <div>
                        <h1 className="font-bold text-xl">Latest Products</h1>
                        {data.map((p, idx) => {
                            return (
                                <div key={idx}>
                                    <ProductCard product={p} showDesc={true} />
                                </div>
                            )
                        })}
                    </div>
                ) : (
                    'No products today :('
                )}
            </h1>
        </div>
    )
}
