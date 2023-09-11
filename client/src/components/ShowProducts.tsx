import { useGetProductsQuery } from '../state/reducers/api'
import Image from 'next/image'
import { ProductResponse } from '../state/reducers/api'
export default function ShowProducts() {
    const { data, error, isLoading } = useGetProductsQuery({ page: 1 })
    const { products } = data ? (data as ProductResponse) : { products: [] }
    return (
        <div>
            {isLoading && 'Loading....'}
            {data &&
                products.map((obj, idx) => {
                    return (
                        <div className="mb-10" key={idx}>
                            {obj.images ? (
                                <>
                                    <div className="card w-72 card-bordered card-compact lg:card-normal border-solid border-black bg-white">
                                        <figure>
                                            <img src={obj.images[0]} />
                                        </figure>
                                        <div className="card-body">
                                            <h2 className="card-title">
                                                {obj.name}
                                            </h2>
                                            <p>{obj.description}</p>
                                            <p>
                                                built by -{' '}
                                                <span className="text-blue-700 cursor-pointer">
                                                    @nwf
                                                </span>
                                            </p>
                                        </div>
                                    </div>
                                </>
                            ) : (
                                <div className="h-20 w-20 bg-red-500 rounded-xl text-center">
                                    <span className="text-white">
                                        {obj.name}
                                    </span>
                                </div>
                            )}
                        </div>
                    )
                })}
        </div>
    )
}
