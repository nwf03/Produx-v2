import { useState } from "react";
import { Button } from "@nextui-org/react";
import { useSignInMutation } from "../../state/reducers/api";
import { User } from "../../state/interfaces";
import { useAppDispatch } from "../../state/hooks";
import { changeAuthStatus } from "../../state/reducers/auth";
import {useRouter} from "next/router";
export default function Login() {
  const [signIn] = useSignInMutation();
  const dispatch = useAppDispatch();
  const [userInfo, setUserInfo] = useState<{
    username: string;
    password: string;
  }>({
    username: "",
    password: "",
  });

  const submitHandler = async (e: any) => {
    await signIn(userInfo)
      .then((res: any) => {
        localStorage.setItem("token", res.data.token);
        dispatch(changeAuthStatus(true));
        window.location.reload()
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <div className="form-control mt-20">
      <input
        className="input input-bordered my-2"
        value={userInfo.username}
        placeholder="name"
        onChange={(e) => setUserInfo({ ...userInfo, username: e.target.value })}
      />
      <input
        value={userInfo.password}
        className="input input-bordered my-2"
        placeholder="password"
        type="password"
        onChange={(e) => setUserInfo({ ...userInfo, password: e.target.value })}
      />
      <Button className="my-1" onClick={submitHandler}>
        Sign In!
      </Button>
    </div>
  );
}
