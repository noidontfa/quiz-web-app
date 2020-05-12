import React from "react";
import "./style.css";
import logo from "./assets/Group 622.png";
import userIcon from "./assets/cat.jpg";

function App() {
  return (
    <>
      <header className="my-navbar">
        <div className="my-container">
          <div className="my-logo">
            <img src={logo} alt="logo" />
          </div>
          <div className="my-icons">
            <a href="#">
              <i className="icon-home"></i>
            </a>
            <a href="#">
              <i className="icon-puzzle"></i>
            </a>
            <a href="#">
              <i className="icon-cup"></i>
            </a>
            <img className="my-user-icon" src={userIcon} alt="user-avartar" />
          </div>
        </div>
      </header>
    </>
  );
}

export default App;
