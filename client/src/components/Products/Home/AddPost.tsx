import { useEffect, useState } from "react";
import { useAppSelector } from "../../../state/hooks";
import { useCreatePostMutation } from "../../../state/reducers/api";
import {Button, Checkbox, Input, Loading, Modal, Text} from "@nextui-org/react";
import {tag} from "postcss-selector-parser";

export default function AddPost({
    show,
  setShow,
  productName,
}: {
  show: boolean;
  setShow: any;
  productName: string;
}) {
  const currentChannel = useAppSelector((state) =>
    state.channel.channel.slice(0, -1)
  );
  const tags = { Bug: "Report a ", Suggestion: "Create a " };
  const tagNames = Object.keys(tags);
  const [createPost, { data, isLoading, error }] = useCreatePostMutation();
  const [selectedTag, setSelectedTag] = useState(
    tagNames.includes(currentChannel) ? currentChannel : tagNames[0]
  );
  const [post, setPost] = useState({
    title: "",
    description: "",
  });
  const addTag = (tag: string) => {
    setSelectedTag(tag);
  };
  useEffect(() => {
    tagNames.includes(currentChannel) ? currentChannel : tagNames[0];
  }, [currentChannel]);
  const [showErr, setShowErr] = useState(false);
  const submitHandler = async () => {
    if (post.title && post.description) {
      const res = await createPost({
        productName,
        channel: `${selectedTag.toLowerCase()}s`,
        post,
      });
      setShow(false);
    }else{
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
          onClose={() => setShow(false)}
      >
                <Modal.Header>
          <Text id="modal-title" size={18}>
            {tags[selectedTag]}
            <Text b size={18} >
              {selectedTag}
            </Text>
          </Text>
        </Modal.Header>
        <Modal.Body css={{ display: "flex" }}>
          {showErr && <p className={'text-center text-red-500'}>Please fill all the fields</p>}
          <div className={""}>
                <div>
                  <div className={"form-control items-center "}>
                    <div className={"flex mb-2 mr-auto ml-2"}>
                      {tagNames.map((tag) => {
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
                            setPost({ ...post, title: e.target.value })
                        }
                        type={"text"}
                        className="w-11/12 input input-bordered focus:border-none font-bold "
                        placeholder={"Title"}
                    />
                    <br />
                    <textarea
                        value={post.description}
                        onChange={(e) =>
                            setPost({ ...post, description: e.target.value })
                        }
                        className="textarea h-40 w-11/12 textarea-bordered"
                        placeholder="Details"
                    />
                    <p className={"text-[12px] mr-auto ml-4 mt-2 text-amber-500"}>
                      image upload coming soon!
                    </p>
                  </div>
                </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button clickable={!isLoading} auto flat color="error" onClick={() => setShow(false)}>
            Cancel
          </Button>
          <Button clickable={!isLoading} auto onClick={submitHandler}>
            {isLoading ? <Loading color={'white'} size={'sm'} /> : 'Create'}
          </Button>
        </Modal.Footer>
        {error && <p className={'mb-4 text-red-500'}>{error.data.message}</p>}
      </Modal>
    </div>
  );
}
