import {
  Row,
  Checkbox,
  Loading,
  Modal,
  Button,
  Input,
  Text,
  FormElement,
  CheckboxEvent,
} from "@nextui-org/react";
import { useAppSelector, useAppDispatch } from "../../../state/hooks";
import {
  setShowStep2,
  setShowStep1,
  setProductData,
} from "../../../state/reducers/createProductSlice";
import {
  ChangeEvent,
  MutableRefObject,
  useRef,
  useState,
  useEffect,
} from "react";
import { NewProduct } from "../../../state/interfaces";
import { useCreateProductMutation } from "../../../state/reducers/api";
import confetti from "canvas-confetti";
export default function ProductImages({ product }: { product: NewProduct }) {
  const newProduct = useAppSelector((state) => state.newProduct);
  const dispatch = useAppDispatch();
  const [createProduct, { error, isLoading }] = useCreateProductMutation();
  const closeHandler = () => {
    setImages([]);
    dispatch(setShowStep2(false));
  };
  const fileRef = useRef() as MutableRefObject<HTMLInputElement>;
  const [images, setImages] = useState<File[]>([]);
  const handleNewImage = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const imgs = [];
      for (
        let i = 0;
        i < (e.target.files.length > 4 ? 4 : e.target.files.length);
        i++
      ) {
        imgs[i] = e.target.files[i];
      }
      setImages(imgs);
    }
  };
  const handleCreate = async () => {
    const productIcon = product.images ? product.images : [];
    const productImages = [...productIcon, ...images];
    dispatch(setProductData({ ...product, images: productImages }));
    setSubmit(true);
    console.log(product);
    const formData = new FormData();
    formData.set("name", product.name);
    formData.set("description", product.description);
    formData.set("accessToken", product.accessToken);
    formData.set("private", product.private.toString());
    for (let i = 0; i < productImages.length; i++) {
      formData.append(`image${i + 1}`, productImages[i]);
    }
    await createProduct(formData);
    setSubmit(false);
    dispatch(setShowStep2(false));
    setTimeout(
      () =>
        confetti({
          zIndex: 999,
          particleCount: 500,
          spread: 170,
          origin: { x: 0.5, y: 0.9 },
        }),
      500
    );
    confetti.reset();
  };
  const [submit, setSubmit] = useState(false);
  return (
    <div>
      <Modal
        closeButton
        blur
        width={"50vw"}
        aria-labelledby="modal-title"
        open={newProduct.showStep2}
        onClose={closeHandler}
      >
        <Modal.Header>
          <Text id="modal-title" size={18}>
            Add up to 4 images of your product
            <Text b size={18} css={{ paddingLeft: 7 }}>
              Product!
            </Text>
          </Text>
        </Modal.Header>
        <Modal.Body>
          <div
            onClick={() => fileRef.current.click()}
            className={
              "h-14 p-4 rounded-box bg-red-400 text-white flex justify-center items-center"
            }
          >
            Select images
          </div>
          {images &&
            images.map((i, idx) => (
              <div key={idx} className={"flex justify-center w-full"}>
                {" "}
                <img
                  src={URL.createObjectURL(i)}
                  className={
                    "max-w-[89%] max-h-[80vh] object-contain rounded-box"
                  }
                />
              </div>
            ))}

          <input
            type={"file"}
            multiple={true}
            ref={fileRef}
            hidden
            onChange={(e) => handleNewImage(e)}
          />
        </Modal.Body>
        <Modal.Footer>
          <Button
            auto
            flat
            onClick={() => {
              dispatch(setShowStep2(false));
              dispatch(setShowStep1(true));
            }}
          >
            Go Back
          </Button>

          <Button auto flat color="error" onClick={closeHandler}>
            Cancel
          </Button>
          <Button auto clickable={!isLoading} onClick={handleCreate}>
            {isLoading ? <Loading color="white" size="sm" /> : "Create!"}
          </Button>
        </Modal.Footer>
        {error && <p className={"text-red-400 mb-3"}>{error.data ? error.data.message : "error"}</p>}
      </Modal>
    </div>
  );
}
