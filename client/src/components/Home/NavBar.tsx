import { useState } from 'react'
import BasicInfo from '../Products/CreateProduct/BasicInfo'
import SearchProducts from '../Products/SearchProducts'
import { useAppSelector } from '../../state/hooks'
import { useRouter } from 'next/router'
import { useAppDispatch } from '../../state/hooks'
import { setShowStep1 } from '../../state/reducers/createProductSlice'

export default function NavBar() {
    const [name, setName] = useState<string>('')
    const dispatch = useAppDispatch()
    const router = useRouter()
    const user = useAppSelector((state) => state.auth.user)
    // const [searchFocused, setSearchFocused] = useState<boolean>(false)
    const logoutHandler = () => {
        localStorage.removeItem('token')
        router.reload()
    }
    return (
        <div>
            <div className="w-[100vw] h-[80px] fixed top-0 bg-[#f5f5f5] ">
                <div className="grid grid-cols-3 justify-center align-middle text-center h-[80px] lg:gap-20 sm:gap-2">
                    <div className="flex justify-left ml-10 items-center">
                        <h1
                            className="text-xl font-bold cursor-pointer"
                            onClick={() => router.push('/')}
                        >
                            Produx
                        </h1>
                    </div>
                    <div className="flex justify-center items-center col-span-1">
                        <div className="form-control w-full">
                            <input
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                type="text"
                                // onFocus={() => setSearchFocused(true)}
                                // onBlur={() => setSearchFocused(false)}
                                placeholder="Search Products..."
                                className="input input-bordered"
                            />
                            {name.length > 1 && (
                                <div className={'flex justify-center'}>
                                    <SearchProducts name={name} />
                                </div>
                            )}
                        </div>
                    </div>
                    <div className="fixed right-10  md:right-20 top-1 justify-center  cursor-pointer  dropdown">
                        <div className="m-1 hover:bg-gray-300  rounded-box p-2 items-center  cursor-pointer  ">
                            <div>
                                <div tabIndex={0} className="flex items-center">
                                    <img
                                        src={user ? user.pfp : ''}
                                        className="rounded-box object-cover w-12 aspect-square"
                                    />
                                    <p className="ml-2 hidden md:block">
                                        {user && user.name}
                                    </p>
                                </div>
                            </div>
                        </div>
                        <ul
                            tabIndex={0}
                            className="p-2  shadow z-40 fixed right-1 md:right-auto md:relative dropdown-content bg-base-100 rounded-xl min-w-max text-left"
                        >
                            <li
                                className={`active:bg-blue-500 p-2 m-0 cursor-pointer rounded-xl active:text-white`}
                                onClick={() => router.push('/profile')}
                            >
                                <a>Profile</a>
                            </li>
                            <hr />
                            <li
                                className={`active:bg-red-500 p-2 m-0 cursor-pointer rounded-xl active:text-white`}
                                onClick={logoutHandler}
                            >
                                <a>logout</a>
                            </li>
                            <hr />
                            <li
                                onClick={() => dispatch(setShowStep1(true))}
                                className="p-2 active:bg-green-500 rounded-xl active:text-white"
                            >
                                <a>Create a Product</a>
                            </li>
                        </ul>
                    </div>
                </div>
                <hr className="h-[2px] bg-gray-200" />
            </div>
            <div className={'mb-24'}></div>
            <BasicInfo />
        </div>
    )
}
