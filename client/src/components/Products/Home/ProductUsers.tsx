import { ProductUser } from "../../../state/interfaces";
import { Avatar, FormElement, Input } from "@nextui-org/react";
import { ChangeEvent, useEffect, useState } from "react";
import { useUpdateUserRoleMutation } from "../../../state/reducers/api";
import { useAppSelector } from "../../../state/hooks";
export default function ProductUsers({
  users,
  productId,
  productUserID,
}: {
  users?: ProductUser[];
  productId: number;
  productUserID: number;
}) {
  const [filteredUsers, setFilteredUsers] = useState(users);
  const [updateRole, { data, error, isLoading }] = useUpdateUserRoleMutation();
  const onlineUsers = useAppSelector((state) => state.product.onlineUsers);
  const [showUserInfo, setShowUserInfo] = useState({
    show: false,
    userId: 0,
  });
  const [role, setRole] = useState<string>("");
  const searchUsers = (e: ChangeEvent<FormElement>) => {
    if (e.target.value) {
      const filtered = users?.filter((user) =>
        user.name.toLowerCase().includes(e.target.value.toLowerCase())
      );
      setFilteredUsers(filtered);
    } else {
      setFilteredUsers(users);
    }
  };
  useEffect(() => {
    setFilteredUsers(users);
  }, [users]);

  const handleRoleUpdate = async (
    e: MouseEvent<HTMLButtonElement>,
    userId: number,
  ) => {
    e.preventDefault();
    if (!role) return;
    await updateRole({ productId, userId, role });
    setShowUserInfo({ show: false, userId: 0 });
  };
  const [userID, setUserId] = useState<number>(0);
  useEffect(() => {
    const auth = localStorage.getItem("auth");
    if (auth) {
      setUserId(JSON.parse(auth).user.ID);
    }
  }, []);

  return (
    <div className="ml-4 mt-4 ">
      <h1 className="font-bold text-3xl mb-3">Users:</h1>
      <Input
        onChange={searchUsers}
        size={"md"}
        className={"w-[16vw]"}
        bordered
        placeholder="Search Users..."
      />
      <br />
      <br />
      {filteredUsers && filteredUsers.length > 0
        ? (
          filteredUsers.map((user, idx) => {
            return (
              <div key={idx} className={"block"}>
                {userID == productUserID
                  ? (
                    <div>
                      <div
                        onClick={() =>
                          setShowUserInfo({
                            show: !showUserInfo.show,
                            userId: user.id,
                          })}
                        className="flex cursor-pointer hover:bg-gray-200 items-center mb-1 p-3 rounded-box"
                      >
                        {user.pfp && (
                          <Avatar
                            size={"lg"}
                            squared
                            bordered
                            src={user.pfp}
                          />
                        )}
                        <h1 className="ml-2">{user.name}</h1>
                        {onlineUsers.filter(
                              (u) => u.id == user.id,
                            ).length > 0 && (
                          <div className="bg-green-500 rounded-3xl w-2 h-2 ml-2" />
                        )}
                        <p className="ml-auto text-sm mr-6 text-orange-400">
                          {user.role}
                        </p>
                      </div>
                      {showUserInfo.show &&
                        showUserInfo.userId == user.id && (
                        <form
                          className={"p-4 rounded-box w-11/12 bg-white "}
                        >
                          <Input
                            placeholder={"change role"}
                            onChange={(e) => setRole(e.target.value)}
                            className={"w-full"}
                          />
                          <button
                            className="hover:bg-blue-500 hover:text-white p-2 rounded-xl"
                            onClick={(e) =>
                              handleRoleUpdate(
                                e,
                                user.id,
                              )}
                          >
                            Change
                          </button>
                        </form>
                      )}
                    </div>
                  )
                  : (
                    <div className="flex cursor-pointer hover:bg-gray-200 items-center mb-1 p-3 rounded-box">
                      {user.pfp && (
                        <Avatar
                          size={"lg"}
                          squared
                          bordered
                          src={user.pfp}
                        />
                      )}
                      <h1 className="ml-2">{user.name}</h1>
                      <p className="ml-auto text-sm mr-6 text-orange-400">
                        {user.role}
                      </p>
                    </div>
                  )}
              </div>
            );
          })
        )
        : <h1>no users</h1>}
      {error && <p className={"text-red-500"}>{error.data.message}</p>}
    </div>
  );
}
