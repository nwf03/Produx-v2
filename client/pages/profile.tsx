import { useAppSelector } from "../src/state/hooks";
import NavBar from '../src/components/Home/NavBar'
import {MutableRefObject} from "react";
import { ChangeEvent } from "react";
import {useRef, useEffect, useState } from "react";
import { User } from "../src/state/interfaces";
import {useDeleteProductMutation, useLazyGetUserInfoQuery} from "../src/state/reducers/api";
import {setUser} from "../src/state/reducers/auth";
import Head from "next/head";
import ProductCard from "../src/components/Products/Product";
import {Button, Modal, Text} from "@nextui-org/react";
export default function Profile() {
  const user = useAppSelector((state) => state.auth.user);
  const [userInfo, setUserInfo] = useState(user);
  const [save, setSave] = useState(false);
  const [trigger, result, lastPromiseInfo] = useLazyGetUserInfoQuery();
  useEffect(() => {
    if (!user) {
      trigger().then((res) => setUserInfo(res.data as User));
      console.log("fetched user data");
    } else {
      console.log("did not fetch user data");
    }
  }, []);
  const changeHandler = (e: ChangeEvent<HTMLInputElement>) => {
    setSave(true);
    if (userInfo) {
      const newUserInfo = { ...userInfo };
      newUserInfo[e.target.id] = e.target.value;
      setUserInfo(newUserInfo);
    }
  };
  const handleReset = () => {
    trigger().then(res=>setUserInfo(res.data as User))
    setSave(false)
  }
  const handleNewPfp = (e:ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.id)
    console.log("uploading pfp")
    if (fileRef.current.files){
      let preview = URL.createObjectURL(fileRef.current.files[0])
      const newUserData = {...userInfo, pfp: preview} as User
      setUserInfo(newUserData)

      setSave(true)
    }
  }
  const [deleteProduct, {isLoading, error, data}] = useDeleteProductMutation()
  const [DelProduct, setDelProduct] = useState({
    name: "",
    id: 0
  })
  const handleDelete =  (name: string, id: number) => {
   setDelProduct({name, id})
   setShowDelModal(true)
  }
  const fileRef = useRef() as MutableRefObject<HTMLInputElement>
  const [showDelModal, setShowDelModal] = useState(false)
  return (
      <div>
        <Head>
          <title>Home</title>
            <link rel="icon" href="/produx2.png" />
        </Head>
        <NavBar/>
    <div className=" ml-10">
      <h1 className="font-bold text-[70px]">Your profile</h1>
      {save && (
          <div className={"ml-2"}>

            <button className="btn btn-primary capitalize mr-5">Save</button>
          <button className={"btn bg-white text-black border-1 border-black capitalize hover:bg-gray-200"} onClick={handleReset}>Reset</button>
          </div>
      )}
      <br />
      {userInfo ? (
        <div className="block lg:flex">
          {userInfo.pfp && (
            <div className="relative hover:cursor-pointer">
              <img
                src={userInfo.pfp}
                className=" bg-gray-200 object-cover w-[80vw] z-0 md:w-auto md:h-[50vh] rounded-3xl"
              />
              <div onClick={()=>fileRef.current.click()} className="opacity-0 hover:opacity-100 w-[80vw] h-full md:w-full rounded-3xl flex absolute top-0  z-10">
                <div className="  bg-black w-[80vw] h-full md:w-full rounded-3xl  opacity-50 absolute top-0 " />
                <div className={"form-control m-auto z-20 "}>
                  <label className={"label"} >
                <button className="btn opacity-90 cursor-pointer capitalize rounded-xl bg-red-400 hover:bg-red-500 text-white ">
                  change photo
                </button>
              </label>
                  <input id={"pfp"} hidden multiple={false} onChange={handleNewPfp} ref={fileRef} type={"file"}/>
              </div>
              </div>
            </div>
          )}
          <div className="grid h-48 ml-10 mr-auto mt-6 grid-rows-2">
            <div className="">
              <label className="label">
                <span>Username</span>
              </label>
              <input
                value={userInfo.name}
                onChange={changeHandler}
                id="name"
                className="input input-bordered text-xl "
              />
            </div>
            <br />
            <div className="">
              <label className="label">
                <span>Email</span>
              </label>
              <input
                value={userInfo.email}
                onChange={changeHandler}
                id="email"
                className="input input-bordered text-xl "
              />
            </div>
          </div>
        </div>
      ) : (
        <p>Error fetching user data</p>
      )}
    </div>
        <br/>
        <div className={'mt-4 ml-10'}>
        <h1 className={'text-3xl font-bold'}>Your Products</h1>
        <div className={'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3'}>
          {userInfo && userInfo.products && userInfo.products.map(p => {
            return <div key={p.ID} className={'flex items-center'}>
              <ProductCard product={p} showDesc={true}/>
              <Button size={'sm'} style={{backgroundColor:'red', marginLeft: 'auto', marginRight:'9px'}}  onClick={() => handleDelete(p.name, p.ID)}>Delete</Button>
              {data &&  <p className={'text-green-500'}>{data.message}</p>}
              {error && <p className={"text-red-500"}>{JSON.stringify(error)}</p>}
            </div>
          })}</div>
        </div>
        <br/>
        <ConfirmDelete productName={DelProduct.name} productID={DelProduct.id} isShown={showDelModal} setIsShown={setShowDelModal}/>
      </div>
  );
}

const ConfirmDelete = ({productName, productID, isShown, setIsShown}: {productName:string, productID: number, isShown: boolean, setIsShown: any}) => {
  const closeHandler = () => {
    setIsShown(false)
  }
  const [deleteProduct, {error, data}] = useDeleteProductMutation()
  const handleDelete =  async () => {
   await deleteProduct({id: productID})
    setIsShown(false)
    window.location.reload()
  }
  return (
      <div>
        <Modal
          closeButton
         blur
          aria-labelledby='modal-title'
          open={isShown}
          onClose={closeHandler}
        >
          <Modal.Header>
            <Text b style={{}} size={18} id={'modal-title'}>
              Delete {productName}
            </Text>
          </Modal.Header>
          <Modal.Body css={{display:'flex'}}>
            <Button auto flat color={'error'} onClick={() => setIsShown(false)}>Cancel</Button>
            <Button auto style={{backgroundColor: 'red'}} onClick={() => handleDelete()}>Delete</Button>
          </Modal.Body>
        </Modal>
      </div>
  )
}