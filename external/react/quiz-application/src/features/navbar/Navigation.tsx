import React from "react";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";
import logo from "../../assets/Group 622.png";
import userIcon from "../../assets/cat.jpg";
import HomePage from "../homepage/HomePage";

const Navigation = () => {
   return (
           <header className="my-navbar">
               <div className="my-container">
                   <div className="my-logo">
                       <img src={logo} alt="logo" />
                   </div>
                   <div className="my-icons">
                       <Link to="/">
                           <i className="icon-home"></i>
                       </Link>
                       <Link to="/dashboard">
                           <i className="icon-puzzle"></i>
                       </Link>
                       <Link to="/acm">
                           <i className="icon-cup"></i>
                       </Link>
                       <Link to="/about">
                           <img className="my-user-icon" src={userIcon} alt="user-avartar" />
                       </Link>
                   </div>
               </div>
           </header>
   )
}



export default Navigation;