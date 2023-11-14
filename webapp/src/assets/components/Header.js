import React, { useState } from "react";
import "./styles/Header.css";
import { usePlayer } from "../../hooks/userHooks";
import jscookie from "js-cookie";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

export default function Header() {
  const [user, setUser] = useState(jscookie.get("username"));
  const nav = useNavigate();

  function LogOut() {
    jscookie.remove("username");
    jscookie.remove("access_token");
    toast(<>Thanks for coming by.<br/> See you around!</>, {
      type: "info",
      icon: "👋🏼",
      isLoading: false,
      autoClose: 5000,
      pauseOnFocusLoss: false,
      pauseOnHover: false,
    });

    setTimeout((e) => {
      nav("/login");
    }, 1000);
  }

  return (
    <header>
      <div className="leftContainer">{user}</div>
      <div className="midContainer">ASTRO</div>
      <div className="rightContainer">
        <a href="#" onClick={LogOut}>
          Logout
        </a>
      </div>
    </header>
  );
}
