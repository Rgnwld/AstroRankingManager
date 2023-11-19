import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import LoginPage from "./routes/Login";
import HomePage from "./routes/Home";
import "react-toastify/dist/ReactToastify.css";
import "./assets/styles/main.css";
import { ToastContainer } from "react-toastify";
import { PlayerProvider } from "./hooks/userHooks";
import MapPage from "./routes/Map";

const router = createBrowserRouter([
  {
    path: "/home",
    element: <HomePage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/map/:mapId",
    element: <MapPage />,
    
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <>
    <PlayerProvider>
      <ToastContainer
        position="top-left"
        newestOnTop={false}
        closeOnClick
        rtl={false}
        theme="colored"
        hideProgressBar
      />
      <RouterProvider router={router} />
    </PlayerProvider>
  </>
);
