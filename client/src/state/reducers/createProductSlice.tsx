import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Product, Channels, NewProduct } from "../interfaces";

interface initialState {
  product: NewProduct;
  productChannels: Channels;
  showStep1: boolean;
  showStep2: boolean;
}
const initialState: initialState = {
  product: {
    name: "",
    description: "",
    images: null
  },
  showStep1: false,
  showStep2: false,
  productChannels: {
    Announcements: true,
    Changelogs: true,
    Bugs: true,
    Suggestions: true,
  },
};

const createProductSlice = createSlice({
  name: "createProduct",
  initialState,
  reducers: {
    setShowStep1(state, action: PayloadAction<boolean>) {
      state.showStep1 = action.payload;
    },
    setShowStep2(state, action: PayloadAction<boolean>) {
      state.showStep2 = action.payload;
    },
    setProductData(state, action: PayloadAction<NewProduct>) {
      state.product = action.payload;
    },
    setProductChannels(state, action: PayloadAction<Channels>) {
      state.productChannels = action.payload;
    },
    resetProductData(state) {
      state.product = {
        name: "",
        description: "",
        images: null,
      };
    },
  },
});
export const {
  setShowStep1,
  setShowStep2,
  setProductData,
  setProductChannels,
  resetProductData,
} = createProductSlice.actions;
export default createProductSlice.reducer;
