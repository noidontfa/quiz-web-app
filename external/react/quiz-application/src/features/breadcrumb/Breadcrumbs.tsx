import React, {ReactNode} from "react";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";
import logo from "../../assets/Group 622.png";
import userIcon from "../../assets/cat.jpg";
import HomePage from "../homepage/HomePage";

interface P {
    children : ReactNode
}

const Breadcrumbs : React.FC<P> = ({children}) => {
    return (
            <nav aria-label="breadcrumb">
                <ol className="breadcrumb my-breadcrumb">
                    {children}
                </ol>
            </nav>
    )
}
export default Breadcrumbs;