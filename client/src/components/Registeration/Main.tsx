import Login from "./Login";
import Register from "./Register";
import { useState, useEffect } from "react";
import { changeAuthStatus } from "../../state/reducers/auth";
import { useAppDispatch } from "../../state/hooks";
export default function Main() {
  const [login, setLogin] = useState(true);
  const dispatch = useAppDispatch();
  useEffect(() => {
    if (localStorage.getItem("token")) {
      dispatch(changeAuthStatus(true));
    }
  }, [dispatch]);

  return (
    <div className="w-[100vw] h-[20vh] justify-center flex  items-center">
      {login ? <Login /> : <Register />}
      <div className="block">
        <button onClick={() => setLogin(!login)} className="btn btn-primary">
          {login ? "Register" : "Login"}
        </button>
      </div>
    </div>
  );
}
