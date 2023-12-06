import React from "react";

import {
  RouterProvider,
  createBrowserRouter,
  redirect,
} from "react-router-dom";
import LoginPage from "./Login";
import HomePage from "./Home";
import "react-toastify/dist/ReactToastify.css";
import "../assets/styles/main.css";
import { ToastContainer, toast } from "react-toastify";
import { PlayerProvider } from "../hooks/userHooks";
import MapPage from "./Map";
import axios from "axios";
import jscookie from "js-cookie";

const router = createBrowserRouter([
  {
    path: "login",
    loader: CheckAccess,
    element: <LoginPage />,
  },
  {
    path: "v1",
    loader: ValidateAccessToken,
    children: [
      {
        path: "map/:mapId",
        element: <MapPage />,
      },
      {
        path: "home",
        element: <HomePage />,
      },
    ],
  },
]);

async function ValidateAccessToken(e) {
  try {
    const res = await axios.get(
      "http://localhost:8080/auth/token?token=" + jscookie.get("access_token")
    );
  } catch (err) {
    toast("Your token has expired", {
      type: "error",
    });
    jscookie.remove("access_token")
    return redirect("/login");
  }

  return null;
}

async function CheckAccess(e) {
  if(jscookie.get("access_token") != null){
    return redirect("/v1/home")
  }
  else return null;
}

export default function Routes() {
  return (
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
}