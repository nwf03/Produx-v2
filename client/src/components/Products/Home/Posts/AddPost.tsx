import { useAppSelector } from "../../../../state/hooks";
import { ChangeEvent, MutableRefObject, useRef, useState } from "react";
import { useCreatePostMutation } from "../../../../state/reducers/api";
import { useRouter } from "next/router";
import { Button, Loading, Modal, Text } from "@nextui-org/react";

export default function AddPost({
  show,
  setShow,
  productName,
  owner,
}: {
  show: boolean;
  setShow: any;
  productName: string;
  owner: boolean;
}) {
  const currentChannel = useAppSelector((state) =>
    state.channel.channel.slice(0, -1)
  );
  const router = useRouter();
  const isOwner = useAppSelector((state) => state.product.isOwner);
  const tags = {
    Bug: "Report a ",
    Suggestion: "Create a ",
    Announcement: "Create a ",
    Changelog: "Create a ",
  };
  const allowedTags = isOwner ? Object.keys(tags) : Object.keys(tags).filter(
    (tag) => tag !== "Announcement" && tag !== "Changelog",
  );
  const [createPost, { data, isLoading, error }] = useCreatePostMutation();
  const [selectedTag, setSelectedTag] = useState(
    allowedTags.includes(currentChannel) ? currentChannel : allowedTags[0],
  );
  const [post, setPost] = useState({
    title: "",
    description: "",
  });
  const addTag = (tag: string) => {
    setSelectedTag(tag);
  };
  const [showErr, setShowErr] = useState(false);
  const allowedFileTypes = ["image/png", "image/jpeg", "video/mp4"];
  const [imgsErr, setImgsErr] = useState("");
  const [images, setImages] = useState<File[]>([]);
  const addImage = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const imgs = [];
      for (
        let i = 0;
        i < (e.target.files.length > 3 ? 3 : e.target.files.length);
        i++
      ) {
        const file = e.target.files[i];
        if (!allowedFileTypes.includes(file.type)) {
          setImgsErr("Invalid file type. Upload png, jpeg, or mp4");
          return;
        }
        imgs.push(file);
      }
      setImages(imgs);
      setImgsErr("");
    }
  };

  const fileRef = useRef() as MutableRefObject<HTMLInputElement>;
  const submitHandler = async () => {
    if (post.title && post.description) {
      const formData = new FormData();
      formData.set("title", post.title);
      formData.set("description", post.description);
      for (let i = 0; i < images.length; i++) {
        formData.set(`image${i + 1}`, images[i]);
      }
      const res = await createPost({
        productName,
        channel: `${selectedTag.toLowerCase()}s`,
        post: formData,
      });
      router.reload();
      setShow(false);
    } else {
      setShowErr(true);
    }
  };
  return (
    <div>
      <Modal
        closeButton
        blur
        aria-labelledby="modal-title"
        open={show}
        width={"600px"}
        onClose={() => setShow(false)}
      >
        <Modal.Header>
          <Text id="modal-title" size={18}>
            {tags[selectedTag]}
            <Text b size={18}>
              {selectedTag}
            </Text>
          </Text>
        </Modal.Header>
        <Modal.Body css={{ display: "flex" }}>
          {showErr && (
            <p className={"text-center text-red-500"}>
              Please fill all the fields
            </p>
          )}
          <div className={""}>
            <div>
              <div className={"form-control items-center "}>
                <div
                  className={"grid grid-cols-4 mb-2 mr-auto ml-2"}
                >
                  {allowedTags.map((tag) => {
                    return (
                      <button
                        key={tag}
                        className={`btn capitalize rounded-xl m-2 ${
                          tag == selectedTag
                            ? "bg-red-500 text-white"
                            : "bg-white text-black"
                        } hover:text-white`}
                        onClick={() => addTag(tag)}
                      >
                        {tag}
                      </button>
                    );
                  })}
                  <br />
                </div>
                <input
                  value={post.title}
                  onChange={(e) =>
                    setPost({
                      ...post,
                      title: e.target.value,
                    })}
                  type={"text"}
                  className="w-11/12 input input-bordered focus:border-none font-bold "
                  placeholder={"Title"}
                />
                <br />
                <textarea
                  value={post.description}
                  onChange={(e) =>
                    setPost({
                      ...post,
                      description: e.target.value,
                    })}
                  className="textarea h-40 w-11/12 textarea-bordered"
                  placeholder="Details"
                />
                <div className="flex items-center mr-auto mx-5 mt-4">
                  <Button
                    color="warning"
                    onClick={() => fileRef.current.click()}
                    auto
                  >
                    Add Images/Videos
                  </Button>
                  <span className="ml-2 text-sm">
                    (3 max)
                  </span>
                </div>
                {imgsErr && (
                  <span className="text-red-400">
                    {imgsErr}
                  </span>
                )}

                {images && (
                  <div
                    className={`mt-4 grid grid-cols-1 md:grid-cols-${images.length} gap-4`}
                  >
                    {images.map((i, idx) => {
                      return (
                        <div
                          className="flex justify-center w-full"
                          key={idx}
                        >
                          <img
                            src={URL.createObjectURL(
                              i,
                            )}
                            className="object-cover w-[400px] h-[150px] rounded-box"
                          />
                        </div>
                      );
                    })}
                  </div>
                )}
              </div>
              <input
                type="file"
                multiple={true}
                ref={fileRef}
                hidden
                onChange={(e) => addImage(e)}
              />
            </div>
          </div>
        </Modal.Body>

        <Modal.Footer>
          <Button
            clickable={!isLoading}
            auto
            flat
            color="error"
            onClick={() => setShow(false)}
          >
            Cancel
          </Button>
          <Button clickable={!isLoading} auto onClick={submitHandler}>
            {isLoading ? <Loading color={"white"} size={"sm"} /> : (
              "Create Post"
            )}
          </Button>
        </Modal.Footer>
        {error && <p className={"mb-4 text-red-500"}>{error.data.message}</p>}
      </Modal>
    </div>
  );
}
