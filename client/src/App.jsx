import { createBrowserRouter, RouterProvider } from "react-router-dom";
import AppLayout from "./ui/AppLayout";
import Home from "./pages/Home";
import Book from "./features/book/Book";
import { getAllBook, getOneBook } from "./services/bookAPI";
import BookDetail from "./features/book/BookDetail";

const App = () => {
  const router = createBrowserRouter([
    {
      element: <AppLayout />,
      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/books/:all/:title?",
          element: <Book />,
          loader: getAllBook,
        },
        {
          path: "book/:id",
          element: <BookDetail />,
          loader: getOneBook,
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
};

export default App;
