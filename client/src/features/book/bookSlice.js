import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

const initialState = {
  bookData: [],
  loading: false,
  error: null,
};

export const addBook = createAsyncThunk("book/addBook", async (bookData) => {
  const res = await fetch("http://localhost:8080/books", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(bookData),
  });
  if (!res.ok) {
    const errorData = await res.json();
    console.error("Hata:", errorData);
    throw new Error("Kitap eklenemedi.");
  }
  return await res.json(); // Assuming the server returns the created book
});

export const bookSlice = createSlice({
  name: "book",
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(addBook.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(addBook.fulfilled, (state, action) => {
        state.loading = false;
        state.bookData = action.payload;
      })
      .addCase(addBook.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
      });
  },
});

export const { setBookData, clearBookData } = bookSlice.actions;
export default bookSlice.reducer;
