import React, { useState } from "react";
import axios from "axios"
import "./Login.css";
import userIcon from "../assets/imgs/user.png";
import padlockIcon from "../assets/imgs/padlock.png";

function LoginPage() {

  const [userInfo, setUserInfo] = useState({user: "", password: ""})

  function UpdateInput(node){
    setUserInfo({...userInfo, [node.target.name]: node.target.value})
  }

  function Submit(event){
    event.preventDefault();
    
  }

  return (
    <div className="loginPage">
      <form className="loginContainer"  onSubmit={Submit}>
        <div className="title">
          <h1>ASTRO</h1>
        </div>
        <div className="inputContainer">
          <div tabIndex="0" className="customInput">
            <img src={userIcon} alt="userIcon" />
            <input onChange={UpdateInput} value={userInfo.user} type="text" placeholder="User" name="user"></input>
          </div>
          <div tabIndex="0" className="customInput">
            <img src={padlockIcon} alt="padlockIcon" />
            <input onChange={UpdateInput} value={userInfo.password} type="password" placeholder="Password" name="password"></input>
          </div>
        </div>
        <div className="buttonContainer">
          <div>
            <button type="submit" className="customButton">Start session</button>
          </div>
          <div className="forgotPass">
            <a href="#">New here?</a>
          </div>
        </div>
      </form>
    </div>
  );
}

export default LoginPage;
