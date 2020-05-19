import React from "react";
import {BrowserRouter, NavLink, Route, Switch} from "react-router-dom";
import logo from "../../assets/Group 622.png";
import userIcon from "../../assets/cat.jpg";
import HomePage from "../homepage/HomePage";

interface P {
    to? : string;
    name: string;
    active?: boolean;
}

const LinkBreadcrumbs : React.FC<P> = ({to, active,name}) => {
    return (
        <li className={active ? 'breadcrumb-item' : 'breadcrumb-item active'}>
            {
                active && <NavLink to={to!}>
                            {name}
                            </NavLink>
            }
            {
                active || <span>
                    {name}
                </span>
            }
        </li>
    )
}

export default LinkBreadcrumbs;