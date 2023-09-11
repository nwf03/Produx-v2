import { Product } from '../../../../state/interfaces'
import { useEffect, useState } from 'react'
import Link from 'next/link'
export default function NavBar({ product }: { product: Product }) {
    const [showMenu, setShowMenu] = useState(false)
    const channels = {
        Announcements: { icon: '🎉', color: '#FF9900' },
        Bugs: { icon: '🐞', color: '#DBFF00' },
        Suggestions: { icon: '🙏', color: '#0094FF' },
        Changelogs: { icon: '🔑', color: '#FF4D00' },
        Questions: { icon: '🤨', color: '#FF4D00' },
    }
    const handleChannelChange = (e: any) => {
        e.preventDefault()
    }
    return (
        <div className={'flex items-center p-4 '}>
            {product.images && (
                <img src={product.images[0]} className={'w-14 mr-3'} />
            )}
            <div className={'block'}>
                <h1 className={'w-max font-bold text-xl'}>{product.name}</h1>
                {product.description.length <= 21 && (
                    <h1 className={'w-max'}>{product.description} </h1>
                )}
            </div>
            <div className={'divider divider-vertical'}></div>
            <div
                className={
                    'xsm:mx-auto xsm:relative  xsm:text-center right-0 absolute  mr-4'
                }
            >
                <div
                    className="cursor-pointer"
                    onClick={() => setShowMenu(!showMenu)}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-6 w-6"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M4 6h16M4 12h16M4 18h16"
                        />
                    </svg>
                </div>
                {showMenu && (
                    <div className="w-screen z-50 fixed left-0 top-20 bg-white">
                        <div className="mx-auto mt-20">
                            {Object.keys(channels).map((c) => {
                                return (
                                    <div
                                        key={c}
                                        className=" text-3xl p-4  hover:bg-gray-200 cursor-pointer mx-12 rounded-xl mb-20"
                                        onClick={() => setShowMenu(false)}
                                    >
                                        <Link
                                            href={{
                                                pathname: `/products/[name]/${
                                                    c === 'Questions'
                                                        ? '/questions'
                                                        : c
                                                }`,
                                                query: {
                                                    name: product.name,
                                                    channel: c,
                                                },
                                            }}
                                        >
                                            <a>
                                                <p>{c}</p>
                                            </a>
                                        </Link>
                                    </div>
                                )
                            })}
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}
