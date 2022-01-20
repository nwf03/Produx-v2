import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { PostTypes } from "../interfaces";

interface ChannelInterface {
  channel: string;
}
const initialState = {
  channel: "Bugs",
};

const channelSlice = createSlice({
  name: "channel",
  initialState,
  reducers: {
    setChannel: (state, action: PayloadAction<string>) => {
      state.channel = action.payload;
    },
  },
});

export const { setChannel } = channelSlice.actions;

export default channelSlice.reducer;
