import { createBrowserRouter, RouterProvider } from "react-router-dom";
import AppLayout from "./ui/AppLayout";
import Home from "./pages/Home";
import CreateBook from "./features/book/CreateBook";

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
          path: "/createBook",
          element: <CreateBook />,
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
};

export default App;
