import React, {useLayoutEffect, useState} from "react";
import userIcon from "../../../assets/cat.jpg";
import {useHistory} from "react-router-dom";
import {Link} from "react-router-dom";

interface P {
    user? : UserInterface;
}

const SideBar : React.FC<P> = ({user}) => {
    const [leftSidebar,setLeftSidebar] = useState('0');
    const history = useHistory();

    useLayoutEffect(() => {
        const mySection = document.getElementById('my-section');
        setLeftSidebar((mySection!.getBoundingClientRect().right + 82 ) + 'px');
        const onResize = () => {
            setLeftSidebar( (mySection!.getBoundingClientRect().right + 92 ) + 'px');
            console.log(window.innerWidth);
        }
        window.addEventListener('resize',onResize)
        return () => {
            window.removeEventListener('resize',onResize);
        }
    },[])

    return (
        <div className="my-sidebar" id="my-sidebar" style={{left: leftSidebar}}>
            {user &&
                <div className="user-info">
                    <div className="my-user-icon" style={{height: '32px', width: '32px'}}>
                        <img
                            src={user.image}
                            className="my-user-icon"
                            alt="author"
                            style={{marginRight: '10px'}}
                        />
                    </div>
                    <div className="my-user-name">
                        <div className="username">{user.username}</div>
                        <div className="fullname">{user.firstName + " " + user.lastName}</div>
                    </div>
                </div>
            }

            <button className="btn-create" onClick={() => {history.push("/quiz/create")}}>
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

            <Link className="side-link" to="/quiz/my">
                <i className="icon-puzzle"></i>
                <span>
                                My quiz
                            </span>
            </Link>
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

