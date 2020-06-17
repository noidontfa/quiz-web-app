import React, {useEffect, useState} from "react";
import Navigation from "../navbar/Navigation";
import UserInfo from "../createquiz/user-info/UserInfo";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import axios from "axios";
import logo from "../../assets/Group 622.png";
import {Link} from "react-router-dom";
import {useHistory} from "react-router-dom";
import Cookies  from 'universal-cookie';

const MyQuiz = () => {
    const history = useHistory();
    const cookies = new Cookies();
    const [quizzes,setQuizzes] = useState<Array<QuizInterface>>([])
    const [user,setUser] = useState<UserInterface>({});
    useEffect(() => {
        axios.get(`http:/api/quizzes/${2}/my`)
            .then(res => {
                const data = res.data;
                console.log(data);
                setQuizzes(data);
            })
        const token = cookies.get('token');
        if (token != undefined) {
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


    return <>
        <Navigation user={user}/>

        <main className="my-main-content">
            <div className="my-container">
                <div className="row">
                    <UserInfo user={user}/>
                    <div className="col-xl-12" style={{margin: "40px 0 20px 0"}}>
                        <div className="seprator"></div>
                    </div>

                    <div className="col-xl-12" style={{marginBottom: "20px"}}>
                        <Breadcrumbs>
                            <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                            <LinkBreadcrumbs  active={false} name={'My quizzes'}/>
                        </Breadcrumbs>
                    </div>

                    <div className="col-xl-12">
                        <div className="d-flex justify-content-between">
                            <button className="btn-create" onClick={() => {
                                const path = '/quiz/create';
                                history.push(path);
                            }}>
                                <i className="icon-rocket"></i>
                                <span>Create quiz</span>
                            </button>
                            <div className="my-search-form">
                                <input
                                    className="search-input"
                                    type="text"
                                    placeholder="Search by name"
                                    aria-label="Search"
                                />
                                <i className="icon-magnifier search-icon"></i>
                            </div>
                        </div>
                    </div>

                    <div className="col-xl-12">
                        <table className="table my-table">
                            <thead>
                            <tr>
                                <th scope="col" row-data="index">#</th>
                                <th scope="col" row-data="name">Name</th>
                                <th scope="col" row-data="description">Description</th>
                                <th scope="col" row-data="createDate">Create Date</th>
                                <th scope="col" row-data="actions">Actions</th>
                            </tr>
                            </thead>
                            <tbody>
                            {
                                quizzes.map((quiz,i) =>
                                    <tr>
                                        <th scope="row">{i + 1}</th>
                                        <td>
                                            {quiz.name}
                                        </td>
                                        <td>
                                            {quiz.description}
                                        </td>
                                        <td>{quiz.createdAt}</td>
                                        <td>
                                            <div className="my-table-icon d-flex align-items-center">
                                                <Link to={`/quiz/${quiz.id}`} className="test">
                                                    <i className="icon-eye"></i>
                                                </Link>
                                                <Link to={`/quiz/edit/${quiz.id}`} className="test">
                                                    <i className="icon-note"></i>
                                                </Link>
                                            </div>
                                        </td>
                                    </tr>
                                )
                            }
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </main>
    </>
}

export default MyQuiz;