import React,{useEffect, useState,useLayoutEffect} from "react";

import Sidebar from "./sidebar/Sidebar";
import Section from "./section/Section";
import Navigation from "../navbar/Navigation";
import Cookies  from 'universal-cookie';
import axios from "axios";
export function HomePage(){
    const cookies = new Cookies();
    const [user,setUser] = useState<UserInterface>({});
    useEffect(() => {
        const token = cookies.get("token");
        if ( token != undefined ){
            axios.get("http:/api/user/info", {
                headers: {
                    Authorization: 'Bearer ' + token
                }
            }).then(res => {
                console.log(res.data);
                setUser(res.data);
            })
        }
    },[])
    return (
        <main className="my-main-content">
            <div className="my-container">
                <Navigation user={user}/>
                <Section/>
                <Sidebar user={user}/>
            </div>
        </main>
    )
}
export default HomePage;
