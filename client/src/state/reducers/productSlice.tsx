import {createSlice, PayloadAction} from '@reduxjs/toolkit';

interface initialState {
    isOwner: boolean,
    productId: number | null
}
const initialState: initialState = {
    isOwner: false,
    productId: null
}
const productSlide = createSlice({
    name: "product",
    initialState,
    reducers: {
        setIsOwner: (state, action: PayloadAction<boolean>) => {
            state.isOwner = action.payload;
        },
        setProductId: (state, action: PayloadAction<number>) => {
           state.productId = action.payload
        }
    }
})
export const {setIsOwner, setProductId} = productSlide.actions;
export default productSlide.reducer;