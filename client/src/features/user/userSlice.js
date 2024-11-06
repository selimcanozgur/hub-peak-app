import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  users: [],
  loading: false,
  error: null,
  status: false,
};

export const userSlice = createSlice({
  name: "user",
  initialState,

  reducers: {
    userLoading: (state) => {
      state.loading = true;
      state.error = null;
    },
    userLoaded: (state, action) => {
      state.loading = false;
      state.error = null;
      state.status = true;
      state.users = action.payload;
    },
    userError: (state, action) => {
      state.loading = false;
      state.error = action.payload;
    },
  },
});

export const { userLoading, userLoaded, userError } = userSlice.actions;

export default userSlice.reducer;
