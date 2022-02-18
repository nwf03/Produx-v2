import QuickQuestions from "../../../src/components/Products/Home/Posts/QuickQuestions";
import { useRouter } from "next/router";
import HomeLayout from "../../../src/components/Products/Home/HomeLayout";
import { useAppDispatch } from "../../../src/state/hooks";
import { setChannel } from "../../../src/state/reducers/channelSlice";
import { useEffect } from "react";
export default function Questions({ productId }: { productId: number }) {
  const router = useRouter();
  const dispatch = useAppDispatch();
  const { name } = router.query;
  useEffect(() => {
    dispatch(setChannel("questions"));
  }, []);
  return (
    <>
      <QuickQuestions productId={productId} />
    </>
  );
}

Questions.Layout = HomeLayout;
