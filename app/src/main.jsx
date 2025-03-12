import { createRoot } from 'react-dom/client'
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { StyledEngineProvider } from "@mui/material/styles";
import './index.css'
import Login from './pages/login';
import Movie from './pages/movie';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Login />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/movie",
    element: <Movie />
  }
]);

createRoot(document.getElementById('root')).render(
  <StyledEngineProvider injectFirst>
    <RouterProvider router={router} />
  </StyledEngineProvider>
)
