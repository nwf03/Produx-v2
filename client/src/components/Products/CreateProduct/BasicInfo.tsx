import {
  Button,
  Checkbox,
  CheckboxEvent,
  FormElement,
  Input,
  Modal,
  Text,
} from "@nextui-org/react";
import {
  ChangeEvent,
  MutableRefObject,
  useEffect,
  useRef,
  useState,
} from "react";
import { NewProduct } from "../../../state/interfaces";
import ProductImages from "./ProductImages";
import { useAppDispatch, useAppSelector } from "../../../state/hooks";
import {
  resetProductData,
  setProductData,
  setShowStep1,
  setShowStep2,
} from "../../../state/reducers/createProductSlice";

export default function BasicInfo() {
  const newProduct = useAppSelector((state) => state.newProduct);
  const fileRef = useRef() as MutableRefObject<HTMLInputElement>;
  const dispatch = useAppDispatch();
  const [product, setProduct] = useState(newProduct.product);
  const [showErrMsg, setShowErrMsg] = useState(false);
  const closeHandler = (cancel: boolean) => {
    if (!cancel) {
      if (product.name && product.description && product.images) {
        dispatch(setProductData(product));
        dispatch(setShowStep1(false));
        dispatch(setShowStep2(true));
      } else {
        setShowErrMsg(true);
      }
    } else {
      dispatch(setShowStep1(false));
      dispatch(resetProductData());
    }
  };
  const changeHandler = (e: ChangeEvent<FormElement>) => {
    const newProductData = { ...product };
    newProductData[e.target.id] = e.target.value;
    setProduct(newProductData);
  };
  const handleProductImage = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      let preview = URL.createObjectURL(e.target.files[0]);
      const newProductData = {
        ...product,
        images: [e.target.files[0]],
      } as NewProduct;
      const file1 = e.target.files[0];
      setProduct(newProductData);
    }
  };
  const handlePrivateChange = (e: CheckboxEvent) => {
    const newProductData = { ...product };
    newProductData.private = e.target.checked;
    setProduct(newProductData);
  };
  return (
    <div>
      <ProductImages product={product} />
      <Modal
        closeButton
        blur
        aria-labelledby="modal-title"
        open={newProduct.showStep1}
        onClose={() => closeHandler(true)}
      >
        <Modal.Header>
          <Text id="modal-title" size={18}>
            Create your new
            <Text b size={18} css={{ paddingLeft: 7 }}>
              Product
            </Text>
          </Text>
        </Modal.Header>
        <Modal.Body css={{ display: "flex" }}>
          <div
            onClick={() => fileRef.current.click()}
            className={`text-center justify-center text-white items-center flex h-32 ${
              !product.images && "bg-gray-300"
            } w-full rounded-box`}
          >
            {product.images
              ? (
                <img
                  className={"h-32 object-contain w-full"}
                  src={URL.createObjectURL(product.images[0])}
                />
              )
              : <p>Add Image</p>}
          </div>
          <input
            id={"icon"}
            hidden
            multiple={false}
            onChange={handleProductImage}
            ref={fileRef}
            type={"file"}
          />
          <Input
            value={product.name}
            onChange={(e) => changeHandler(e)}
            id={"name"}
            clearable
            bordered
            fullWidth
            color="primary"
            size="lg"
            placeholder="Product Name"
          />
          <Input
            value={product.description}
            onChange={(e) => changeHandler(e)}
            id={"description"}
            clearable
            bordered
            fullWidth
            color="primary"
            size="lg"
            placeholder="Product Description"
          />

          <div
            className={"flex items-center justify-center text-center"}
          >
            <Checkbox
              onChange={(e) => handlePrivateChange(e)}
              id="private"
              checked={product.private}
            >
              <p className={"text-2sm"}>private</p>
            </Checkbox>
            {product.private && (
              <Input
                value={product.accessToken}
                id="accessToken"
                onChange={changeHandler}
                className="ml-4 transition ease-in-out delay-150 bg-blue-500 hover:-translate-y-1 hover:scale-110 hover:bg-indigo-500 duration-300"
                placeholder="Access Code..."
              />
            )}
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button
            auto
            flat
            color="error"
            onClick={() => closeHandler(true)}
          >
            Cancel
          </Button>
          <Button auto onClick={() => closeHandler(false)}>
            Next
          </Button>
        </Modal.Footer>
        {showErrMsg && (
          <p className={"mb-4 text-red-500"}>
            Please fill all the fields
          </p>
        )}
      </Modal>
    </div>
  );
}
