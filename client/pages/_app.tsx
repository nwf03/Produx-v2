import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { Provider } from 'react-redux'
import store from '../src/state/store'
import Register from '../src/components/Registeration/Main'
import { DndProvider } from 'react-dnd'

import { useGetUserInfoQuery } from '../src/state/reducers/api'
import { HTML5Backend } from 'react-dnd-html5-backend'
function MyApp({ Component, pageProps }: AppProps) {
    const Layout = Component.Layout || EmptyLayout
    return (
        <Provider store={store}>
            <DndProvider backend={HTML5Backend}>
                <DefaultLayout>
                    <Layout>
                        <Component {...pageProps} />
                    </Layout>
                </DefaultLayout>
            </DndProvider>
        </Provider>
    )
}
const EmptyLayout = ({ children }: { children: any }) => <>{children}</>
const DefaultLayout = ({ children }: { children: any }) => {
    const { data, isLoading, error } = useGetUserInfoQuery()
    return (
        <>
            {error && <Register />}
            {data && children}
        </>
    )
}
export default MyApp
