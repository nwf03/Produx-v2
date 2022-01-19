import { useState } from "react";
import SearchProducts from "../Products/SearchProducts";
import { useAppSelector } from "../../state/hooks";
export default function NavBar() {
  const [name, setName] = useState<string>("");
  const user = useAppSelector((state) => state.auth.user);
  const [searchFocused, setSearchFocused] = useState<boolean>(false);
  return (
    <div>
      <div className="w-[100vw] h-[80px] fixed top-0 bg-[#f5f5f5] ">
        <div className="grid grid-cols-3 justify-center align-middle text-center h-[80px] lg:gap-20 sm:gap-2">
          <div className="flex justify-left ml-10 items-center ">
            <h1 className="text-xl font-bold">Produx</h1>
          </div>
          <div className="flex justify-center items-center col-span-1">
            <div className="form-control w-full">
              <input
                value={name}
                onChange={(e) => setName(e.target.value)}
                type="text"
                onFocus={() => setSearchFocused(true)}
                onBlur={() => setSearchFocused(false)}
                placeholder="Search Products..."
                className="input input-bordered"
              />
              {name.length > 1 && <SearchProducts name={name} />}
            </div>
          </div>
          <div className="flex justify-center  cursor-pointer ">
            <div className="flex justify-center items-center hover:bg-gray-200 p-2 rounded-box">
              <img
                src={user ? user.pfp : ""}
                className="rounded-box object-cover w-14"
              />
              <p className="ml-2 hidden md:block">{user && user.name}</p>
            </div>
          </div>
        </div>
        <hr className="h-[2px] bg-gray-200" />
      </div>
    </div>
  );
}
