import { useState } from "react";
import { User } from "../../state/interfaces";
export default function Register() {
  const [userInfo, setUserInfo] = useState<Pick<User, "username" | "password">>(
    {
      username: "",
      password: "",
    }
  );
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
      <button className="btn my-1">Register</button>
    </div>
  );
}
