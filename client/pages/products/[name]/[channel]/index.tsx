import {useRouter} from "next/router";
import Head from "next/head";
export default function Channel(){
    const router = useRouter()
    const {channel} = router.query
    return(
        <div>
            <Head>
                <title>{channel}</title>
            </Head>
            {channel}
        </div>
    )
}