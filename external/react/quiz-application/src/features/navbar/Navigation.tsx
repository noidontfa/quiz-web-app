import React from "react";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";
import logo from "../../assets/Group 622.png";
import userIcon from "../../assets/cat.jpg";
import HomePage from "../homepage/HomePage";

interface P {
    user? : UserInterface;
}

const Navigation : React.FC<P> = ({user}) => {
   return (
           <header className="my-navbar">
               <div className="my-container">
                   <Link to="/">
                       <div className="my-logo">
                           <img src={logo} alt="logo" />
                       </div>
                   </Link>
                   <div className="my-icons">
                       <Link to="/">
                           <i className="icon-home"></i>
                       </Link>
                       {
                           user?.username &&
                           <Link to="/quiz/my">
                               <i className="icon-puzzle"></i>
                           </Link>
                       }
                       <Link to="/acm">
                           <i className="icon-cup"></i>
                       </Link>
                       {
                           user?.username &&
                               <Link to="/about">
                                   <img className="my-user-icon" src={user?.image} alt="user-avartar" />
                               </Link>
                           ||
                               <Link to="/login">
                                   <span className="active">Login</span>
                               </Link>
                       }
                   </div>
               </div>
           </header>
   )
}



export default Navigation;