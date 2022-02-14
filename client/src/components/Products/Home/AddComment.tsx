import { ChangeEvent, MouseEventHandler, useState } from "react";
import { Loading, Button, Textarea } from "@nextui-org/react";
import { useCreateCommentMutation } from "../../../state/reducers/api";
export default function AddComment({
  field,
  postId,
}: {
  field: string;
  postId: number;
}) {
  const [comment, setComment] = useState("");
  const [createComment, { data, error, isLoading }] =
    useCreateCommentMutation();
  const [shoErr, setShowErr] = useState(false);
  const submitHandler = async (e: any) => {
    e.preventDefault();
    if (!comment) {
      setShowErr(true);
      return;
    }
    setShowErr(false);
    await createComment({ postId, field, comment });
    setComment("");
  };
  return (
    <div className={"mx-14"}>
      {shoErr && (
        <div className={"text-sm text-red-500"}>Comment is required</div>
      )}
      <form>
        <Textarea
          fullWidth
          minRows={3}
          maxRows={10}
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          placeholder="Add a comment..."
        />
        <Button onClick={submitHandler} className={"mt-2 z-0"} auto>
          {isLoading ? <Loading color={"white"} size={"sm"} /> : "Add comment"}
        </Button>
      </form>
      <br />
    </div>
  );
}
