import Login from "./Login";
import Register from "./Register";
import { useState, useEffect } from "react";
import { changeAuthStatus } from "../../state/reducers/auth";
import { useAppDispatch } from "../../state/hooks";
import { Button, Input } from "@nextui-org/react";
export default function Main() {
  const [login, setLogin] = useState(true);
  const dispatch = useAppDispatch();
  useEffect(() => {
    if (localStorage.getItem("token")) {
      dispatch(changeAuthStatus(true));
    }
  }, [dispatch]);

  return (
    <div className="justify-center mx-auto w-[300px] items-center">
      {login ? <Login /> : <Register />}
      <div className="block">
        <button className="mt-2" onClick={() => setLogin(!login)}>
          {login ? "Register" : "Login"}
        </button>
      </div>
    </div>
  );
}
