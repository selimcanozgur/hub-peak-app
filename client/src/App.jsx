import { createBrowserRouter, RouterProvider } from "react-router-dom";
import AppLayout from "./ui/AppLayout";
import Home from "./pages/Home";
import Book from "./features/book/Book";
import { getAllBook, getOneBook } from "./services/bookAPI";
import BookDetail from "./features/book/BookDetail";
import Error from "./components/Error";

const App = () => {
  const router = createBrowserRouter([
    {
      element: <AppLayout />,
      errorElement: <Error />,

      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/books/:all/:title?",
          element: <Book />,
          loader: getAllBook,
          errorElement: <Error />,
        },
        {
          path: "book/:id",
          element: <BookDetail />,
          loader: getOneBook,
          errorElement: <Error />,
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
};

export default App;
