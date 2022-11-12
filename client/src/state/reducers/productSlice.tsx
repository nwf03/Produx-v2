import { createSlice, PayloadAction } from '@reduxjs/toolkit'
interface User {
    id: number
    name: string
    pfp: string
}
interface initialState {
    isOwner: boolean
    productId: number | null
    onlineUsers: User[]
}
const initialState: initialState = {
    isOwner: false,
    productId: null,
    onlineUsers: [],
}
const productSlide = createSlice({
    name: 'product',
    initialState,
    reducers: {
        setIsOwner: (state, action: PayloadAction<boolean>) => {
            state.isOwner = action.payload
        },
        setProductId: (state, action: PayloadAction<number>) => {
            state.productId = action.payload
        },
        setOnlineUsers: (state, action: PayloadAction<User[]>) => {
            state.onlineUsers = action.payload
        },
    },
})
export const { setIsOwner, setProductId, setOnlineUsers } = productSlide.actions
export default productSlide.reducer
