import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import ShowProducts from '../src/components/ShowProducts'
import NavBar from '../src/components/Home/NavBar'
import Register from '../src/components/Registeration/Main'
import { useAppDispatch, useAppSelector } from '../src/state/hooks'
import HomePage from '../src/components/Home/Home'
import { changeAuthStatus, setUser } from '../src/state/reducers/auth'
import { useEffect } from 'react'
import { useLazyGetUserInfoQuery } from '../src/state/reducers/api'
const Home: NextPage = () => {
    const auth = useAppSelector((state) => state.auth)
    const dispatch = useAppDispatch()
    const user = useAppSelector((state) => state.auth.user)
    const [getUserInfo] = useLazyGetUserInfoQuery()
    const authStatus = useAppSelector((state) => state.auth.isLoggedIn)
    const getUserInfoHandler = async () => {
        const { data, error } = await getUserInfo()
        if (error) {
            dispatch(changeAuthStatus(false))
        } else if (data) {
            dispatch(setUser(data))
        }
    }
    useEffect(() => {
        getUserInfoHandler()
    }, [])
    return (
        <div>
            <Head>
                <link rel="icon" href={'/produx2.png'} />
            </Head>
            <div>
                <NavBar />
                <HomePage />
            </div>
        </div>
    )
}
export default Home
