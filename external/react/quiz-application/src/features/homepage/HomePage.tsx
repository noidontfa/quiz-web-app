import React,{useEffect, useState,useLayoutEffect} from "react";

import Sidebar from "./sidebar/Sidebar";
import Section from "./section/Section";
import Navigation from "../navbar/Navigation";
export function HomePage(){
    return (
        <main className="my-main-content">
            <div className="my-container">
                <Navigation/>
                <Section/>
                <Sidebar/>
            </div>
        </main>
    )
}
export default HomePage;
