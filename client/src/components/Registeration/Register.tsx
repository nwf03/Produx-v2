import { Button } from '@nextui-org/react'
import { ChangeEvent, MutableRefObject, useRef, useState } from 'react'
import { User } from '../../state/interfaces'
import { useAppDispatch } from '../../state/hooks'
import { changeAuthStatus } from '../../state/reducers/auth'
import { useCreateUserMutation } from '../../state/reducers/api'
import { useRouter } from 'next/router'
export default function Register() {
    const [userInfo, setUserInfo] = useState<
        Pick<User, 'name' | 'email'> & { password: string; pfp: File | null }
    >({
        name: '',
        password: '',
        email: '',
        pfp: null,
    })
    const fileRef = useRef() as MutableRefObject<HTMLInputElement>
    const [img, setImg] = useState('')
    const [createUser, { error, isLoading }] = useCreateUserMutation()
    const handleImg = (e: ChangeEvent<HTMLInputElement>) => {
        if (e.target.files) {
            setUserInfo({ ...userInfo, pfp: e.target.files[0] })
        }
    }
    const dispatch = useAppDispatch()
    const handleSubmit = async (e: ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        if (userInfo.pfp) {
            const formData = new FormData()
            formData.append('name', userInfo.name)
            formData.append('email', userInfo.email)
            formData.append('password', userInfo.password)
            formData.append('pfp', userInfo.pfp)
            await createUser(formData).then((res: any) => {
                localStorage.setItem('token', res.data.token)
                dispatch(changeAuthStatus(true))
                window.location.reload()
            })
        }
    }
    return (
        <div className="form-control mt-20">
            <input
                type="file"
                hidden
                ref={fileRef}
                onChange={handleImg}
                multiple={false}
            />
            <div
                onClick={() => fileRef.current.click()}
                className={`cursor-pointer text-center justify-center text-white items-center flex w-[300px] h-[270px] ${
                    !userInfo.pfp && 'bg-gray-300'
                } w-full rounded-box`}
            >
                {userInfo.pfp ? (
                    <img
                        className={'h-full object-contain w-full'}
                        src={URL.createObjectURL(userInfo.pfp)}
                    />
                ) : (
                    <p>Add Image</p>
                )}
            </div>
            <input
                className="input input-bordered my-2"
                value={userInfo.name}
                placeholder="name"
                onChange={(e) =>
                    setUserInfo({ ...userInfo, name: e.target.value })
                }
            />
            <input
                value={userInfo.email}
                className="input input-bordered my-2"
                placeholder="email"
                type="email"
                onChange={(e) =>
                    setUserInfo({ ...userInfo, email: e.target.value })
                }
            />
            <input
                value={userInfo.password}
                className="input input-bordered my-2"
                placeholder="password"
                type="password"
                onChange={(e) =>
                    setUserInfo({ ...userInfo, password: e.target.value })
                }
            />
            <Button onClick={handleSubmit} className="my-1">
                Register
            </Button>
        </div>
    )
}
