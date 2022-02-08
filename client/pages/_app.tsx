import "../styles/globals.css";
import type { AppProps } from "next/app";
import { Provider } from "react-redux";
import store from "../src/state/store";
import Register from "../src/components/Registeration/Main";
import {
  useGetUserInfoQuery,
} from "../src/state/reducers/api";
function MyApp({ Component, pageProps }: AppProps) {
  const Layout = Component.Layout || EmptyLayout;
  return (
    <Provider store={store}>
      <DefaultLayout>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </DefaultLayout>
    </Provider>
  );
}
const EmptyLayout = ({ children }: {children: any}) => <>{children}</>;
const DefaultLayout = ({ children }: {children : any}) => {
  const { data, isLoading, error } = useGetUserInfoQuery();
    return <>
      {error && <Register/>}
      {data && children}
    </>
};
export default MyApp;
