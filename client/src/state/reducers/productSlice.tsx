import {createSlice, PayloadAction} from '@reduxjs/toolkit';

interface initialState {
    isOwner: boolean,
}
const initialState: initialState = {
    isOwner: false,
}
const productSlide = createSlice({
    name: "product",
    initialState,
    reducers: {
        setIsOwner: (state, action: PayloadAction<boolean>) => {
            state.isOwner = action.payload;
        }
    }
})
export const {setIsOwner} = productSlide.actions;
export default productSlide.reducer;