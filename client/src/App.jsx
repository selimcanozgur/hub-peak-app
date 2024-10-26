import { createBrowserRouter, RouterProvider } from "react-router-dom";
import AppLayout from "./ui/AppLayout";
import Home from "./pages/Home";
import Book from "./features/book/Book";
import { getAllBook } from "./services/bookAPI";

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
          path: "/books",
          element: <Book />,
          loader: getAllBook,
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
};

export default App;
