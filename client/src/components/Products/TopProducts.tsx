
import {useGetTopProductsQuery} from "../../state/reducers/api";

export default function TopProducts() {
  const {data, error, isLoading} = useGetTopProductsQuery()
    return (
    <div>
      <h1 className="mt-20 bg-gray-200 p-6 rounded-box w-44 fixed">
          {isLoading && "Loading...."}
          {data && data.length > 0 ? data.map((p, idx)=>{
              return (
                  <div key={idx}>
                      <h1>{JSON.stringify(p)}</h1>
                  </div>
              )
          }) : "Coming Soon..."}
      </h1>
    </div>
  );
}
