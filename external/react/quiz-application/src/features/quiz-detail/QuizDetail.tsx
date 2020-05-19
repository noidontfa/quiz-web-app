import React, {useEffect, useState} from "react";

import {useParams, useHistory} from 'react-router-dom'
import Navigation from "../navbar/Navigation";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import userIcon from "../../assets/cat.jpg";
import axios from "axios";
import Ratings from "../rating/Ratings";
import QuizInfo from "./quiz-info/QuizInfo";
import HistoryTable from "./history-table/HistoryTable";

export function QuizDetail(){
    const { quizId } = useParams();
    const [quiz, setQuiz] = useState<QuizInterface>({});
    const history = useHistory()

    const doStuff =  async () => {
        try {
            const response = await axios.get(`http:/api/quizzes/${quizId}`);
            console.log(response.data);
            setQuiz(response.data);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        doStuff();
    },[])
    return(

        <main className="my-main-content">
            <div className="my-container">
                <div className="row">
                    <Navigation/>
                    <div className="col-xl-12">
                        <Breadcrumbs >
                            <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                            <LinkBreadcrumbs active={false} name={quiz.name!}/>
                        </Breadcrumbs>
                    </div>



                    <div className="col-xl-9" style={{marginTop: '32px'}}>
                        <div className="row">
                            <QuizInfo quiz={quiz}/>
                            <HistoryTable/>
                        </div>
                    </div>
                    <div className="col-xl-3" style={{marginTop: '32px'}}>
                        <div className="row">
                            <div className="col-xl-12 d-flex justify-content-center">
                                <button className="btn-play" style={{marginBottom: '184px'}} onClick={ () => {
                                    const path = `/quiz/play/${quiz.id}`
                                    history.push(path);
                                }}>
                                    <i className="icon-puzzle"></i>
                                    <span>Play quiz</span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    )
}
export default QuizDetail;
