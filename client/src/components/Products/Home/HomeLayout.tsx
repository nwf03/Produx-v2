import { useRouter } from "next/router";
import HomeNavBar from "../../Home/NavBar";
import {
  useLazyGetProductInfoQuery,
  useLazyCheckIfProductFollowedQuery,
} from "../../../state/reducers/api";
import { Product, ProductUser } from "../../../state/interfaces";
import Sidebar from "./Sidebar";
import Posts from "./Posts/Posts";
import { useState, useEffect, cloneElement } from "react";
import SmallNavBar from "./SMScreens/NavBar";
import { useAppSelector, useAppDispatch } from "../../../state/hooks";
import ChannelFilters from "./SMScreens/ChannelFilters";
import NavBar from "../../Home/NavBar";
import ProductUsers from "./ProductUsers";
import Head from "next/head";
import AccessPrivateProduct from "../AccessPrivateProduct";
import { Loading } from "@nextui-org/react";
import {setIsOwner, setProductId} from "../../../state/reducers/productSlice";
interface ProductResponse {
  product: Product;
  users: ProductUser[];
}
import LoadingSpinner from "../../LoadingSpinner";
export default function ProductHome({ children }) {
  const dispatch = useAppDispatch();
  const router = useRouter();
  const channel = useAppSelector((state) => state.channel.channel);
  const { name } = router.query;
  console.log(name);
  const [checkIfFollowed] = useLazyCheckIfProductFollowedQuery();
  const [getProduct, { data, isLoading, error }] = useLazyGetProductInfoQuery();

  useEffect(() => {
    if (name) {
      getProduct(name as string);
    }
  }, [getProduct, name]);
  const [showPrivate, setShowPrivate] = useState(false);
  useEffect(() => {
    if (data && data.product.private) {
      console.log("product id: ", JSON.stringify(data));
      checkIfFollowed(data.product.ID).then((res) => {
        console.log(res);
        if (res.data.followed) {
          setShowPrivate(false);
        } else {
          setShowPrivate(true);
        }
      });
    }else if (data){
      dispatch(setIsOwner(data.owner));
      dispatch(setProductId(data.product.ID))
    }
  }, [data]);
  const showProductUsers = router.pathname != "/products/[name]";
  return (
    <div>
      <Head>
        <title>{name} - produx</title>
        <link
          rel="icon"
          href={
            data && data.product.images
              ? data.product.images[0]
              : "/produx2.png"
          }
        />
      </Head>
      {data && (
        <div className="grid  md:grid-cols-4 lg:grid-cols-5 items-center h-screen">
          <div className="hidden md:block">
            <Sidebar product={data.product} userCount={data.usersCount} postCount={data.postsCount} channel={channel} />
          </div>
          <div className={"md:hidden"}>
            <SmallNavBar product={data.product} />
          </div>
          <div
            className={`bg-white col-span-4 md:col-span-3 ${
              showProductUsers ? "lg:col-span-3" : "lg:col-span-4"
            } h-screen overflow-y-scroll overflow-x-hidden`}
          >
            {cloneElement(children, {
              productId: data.product.ID,
              owner: data.owner,
            })}
          </div>
          {showProductUsers && (
            <div className="hidden lg:block h-screen">
              <ProductUsers
                users={data ? data.users : []}
                productId={data ? data.product.ID : 0}
                productUserID={data ? data.product.userID : 0}
              />
            </div>
          )}
        </div>
      )}
      {data && (
        <AccessPrivateProduct
          show={showPrivate}
          setShow={setShowPrivate}
          productName={data.product.name}
          productIcon={data.product.images ? data.product.images[0] : null}
        />
      )}
      {error && (
        <div className="flex w-screen h-screen items-center justify-center">
          <h1 className="font-bold text-4xl">
            Product <span className="text-red-400">&apos;{name}&apos;</span> Not
            Found
          </h1>
        </div>
      )}
      {isLoading && <LoadingSpinner />}
    </div>
  );
}
