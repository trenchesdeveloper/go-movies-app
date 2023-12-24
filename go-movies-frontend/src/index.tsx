import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Home from "./components/Home";
import Movies from "./components/Movies";
import Genres from "./components/Genres";
import EditMovie from "./components/EditMovie";
import Graphql from "./components/Graphql";
import Login from "./components/Login";
import ManageCatalogue from "./components/ManageCatalogue";
import Movie from "./components/Movie";

const router = createBrowserRouter([{ path: "/", element: <App />,
  errorElement: <ErrorPage/>,
  children:[
    {index:true, element: <Home />},
    {path: "/movies", element: <Movies />},
    {path: "/movies/:id", element: <Movie />},
    {path: "/genres", element: <Genres />},
    {path: "/admin/movie/0", element: <EditMovie />},
    {path: "/manage-catalogue", element: <ManageCatalogue />},
    {path: "/graphql", element: <Graphql/>},
    {path: "/login", element: <Login/>},
  ]
}]);

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
