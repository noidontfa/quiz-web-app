import React, {useLayoutEffect, useState} from "react";
import userIcon from "../../../assets/cat.jpg";



const SideBar = () => {
    const [leftSidebar,setLeftSidebar] = useState('0');

    useLayoutEffect(() => {
        const mySection = document.getElementById('my-section');
        setLeftSidebar((mySection!.getBoundingClientRect().right + 83 ) + 'px');
        const onResize = () => {
            setLeftSidebar((mySection!.getBoundingClientRect().right + 83 ) + 'px');
            console.log(window.innerWidth);
        }
        window.addEventListener('resize',onResize)
        return () => {
            window.removeEventListener('resize',onResize);
        }
    },[])

    return (
        <div className="my-sidebar" id="my-sidebar" style={{left: leftSidebar}}>
            <div className="user-info">
                <div className="my-user-icon" style={{height: '32px', width: '32px'}}>
                    <img
                        src={userIcon}
                        className="my-user-icon"
                        alt="author"
                        style={{marginRight: '10px'}}
                    />
                </div>
                <div className="my-user-name">
                    <div className="username">thinhntg</div>
                    <div className="fullname">Ngo Tran Gia Thinh</div>
                </div>
            </div>

            <button className="btn-create">
                <i className="icon-rocket"></i>
                <span>Create quiz</span>
            </button>

            <a className="side-link active" href="#">
                <i className="icon-home"></i>
                <span>
                                Home
                            </span>
            </a>

            <div className="seprator"></div>

            <div className="content-header">User interface</div>

            <a className="side-link" href="#">
                <i className="icon-puzzle"></i>
                <span>
                                My quiz
                            </span>
            </a>
            <a className="side-link" href="#">
                <i className="icon-cup"></i>
                <span>
                                Reports
                            </span>
            </a>
            <a className="side-link" href="#">
                <i className="icon-settings"></i>
                <span>
                                Settings
                            </span>
            </a>
            <a className="side-link" href="#">
                <i className="icon-share"></i>
                <span>
                                Log out
                            </span>
            </a>

            <div className="seprator"></div>

            <div className="content-header">Quizzes</div>

            <a className="side-link" href="#">
                <i className="icon-star"></i>
                <span>
                                Ratings
                            </span>
            </a>

            <a className="side-link" href="#">
                <i className="icon-layers"></i>
                <span>
                                Categories
                            </span>
            </a>
            <a className="side-link" href="#">
                <i className="icon-phone"></i>
                <span>
                                Languages
                            </span>
            </a>

            <footer className="my-footer">
                Ngo Tran Gia Thinh © 2020 Copyright
            </footer>
        </div>

    )
}

export default SideBar;

