import { configureStore } from '@reduxjs/toolkit'
import api from './reducers/api'
import { setupListeners } from '@reduxjs/toolkit/query/react'
import authReducer from './reducers/auth'
import channelReducer from './reducers/channelSlice'
import createProductReducer from './reducers/createProductSlice'
import productReducer from './reducers/productSlice'
const store = configureStore({
    reducer: {
        auth: authReducer,
        [api.reducerPath]: api.reducer,
        channel: channelReducer,
        newProduct: createProductReducer,
        product: productReducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(api.middleware),
})
setupListeners(store.dispatch)
export default store
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
