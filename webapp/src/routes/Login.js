import React, { useState } from "react";
import axios from "axios";
import "./styles/Login.css";
import userIcon from "../assets/imgs/user.png";
import padlockIcon from "../assets/imgs/padlock.png";
import jscookie from "js-cookie";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";
import { usePlayer } from "../hooks/userHooks";

function LoginPage() {
  let nav = useNavigate();

  const [userInfo, setUserInfo] = useState({ user: "", password: "" });
  const [editableContent, setEditableContent] = useState(true);

  function UpdateInput(node) {
    if (editableContent)
      setUserInfo({ ...userInfo, [node.target.name]: node.target.value });
  }

  function Submit(event) {
    event.preventDefault();
    setEditableContent(false);

    console.log("working hard");

    const id = toast("Connecting to server...", {
      isLoading: true,
      autoClose: 5000,
    });

    axios
      .post("http://localhost:8080/auth/login", {
        username: userInfo.user,
        password: userInfo.password,
      })
      .then((res) => {
        toast.update(id, {
          render: (
            <>
              <strong>Connected!</strong>
              <br />
              Welcome!
            </>
          ),
          type: "success",
          isLoading: false,
          autoClose: 5000,
          pauseOnFocusLoss: false,
          pauseOnHover: false,
        });

        jscookie.set("access_token", res.data.token);
        jscookie.set("username", userInfo.user);

        setTimeout(() => {
          nav("/v1/home");
        }, 500);
      })
      .catch((err) => {
        if (err.code == "ERR_NETWORK") {
          toast.update(id, {
            render: (
              <>
                We are facing internal error problems. <br />
                Please contact the support service.
              </>
            ),
            type: "error",
            isLoading: false,
            autoClose: 5000,
            pauseOnFocusLoss: false,
            pauseOnHover: false,
          });
          return;
        }

        if (err.response.status >= 500)
          toast.update(id, {
            render: (
              <>
                We are facing internal error problems. <br />
                Please contact the support service.
              </>
            ),
            type: "error",
            isLoading: false,
            autoClose: 5000,
            pauseOnFocusLoss: false,
            pauseOnHover: false,
          });
        else if (err.response.status == 401)
          toast.update(id, {
            render: (
              <>
                <strong>{err.response.status}:</strong> Something went wrong.
                <br />
                Try check your credentials.
              </>
            ),
            type: "error",
            isLoading: false,
            autoClose: 5000,
            pauseOnFocusLoss: false,
            pauseOnHover: false,
          });
        else if (err.response.status == 404)
          toast.update(id, {
            render: (
              <>
                <strong>{err.response.status}:</strong> Server not reached.{" "}
                <br />
                Please contact the support.
              </>
            ),
            type: "error",
            isLoading: false,
            autoClose: 5000,
            pauseOnFocusLoss: false,
            pauseOnHover: false,
          });
        else
          toast.update(id, {
            autoClose: 5000,
            render: (
              <>
                <strong>{err.response.status}:</strong> Unhandled error.
                <br />
                Please contact the support.
              </>
            ),
            type: "error",
            isLoading: false,
            autoClose: 5000,
            pauseOnFocusLoss: false,
            pauseOnHover: false,
          });
      })
      .finally((e) => {
        setEditableContent(true);
      });
  }

  function CreateUser() {
    toast(
      <>
        <strong>Play the game!</strong>
        <br />
        Click{" "}
        <a href="https://rgnwld.itch.io/astro" target="_blank">
          Here
        </a>{" "}
        to access it
      </>,
      { type: "info", closeButton: true, closeOnClick: false, autoClose: false }
    );
  }

  return (
    <div className="basePage loginPage">
      <form className="loginContainer" onSubmit={Submit}>
        <div className="title">
          <h1>ASTRO</h1>
        </div>
        <div className="inputContainer">
          <div tabIndex="0" className="customInput">
            <img src={userIcon} alt="userIcon" />
            <input
              onChange={UpdateInput}
              value={userInfo.user}
              type="text"
              placeholder="User"
              name="user"
            ></input>
          </div>
          <div tabIndex="0" className="customInput">
            <img src={padlockIcon} alt="padlockIcon" />
            <input
              onChange={UpdateInput}
              value={userInfo.password}
              type="password"
              placeholder="Password"
              name="password"
            ></input>
          </div>
        </div>
        <div className="buttonContainer">
          <div>
            <button type="submit" className="customButton">
              Start session
            </button>
          </div>
          <div className="forgotPass">
            <a href="#" onClick={CreateUser}>
              New here?
            </a>
          </div>
        </div>
      </form>
    </div>
  );
}

export default LoginPage;
