import React, {useEffect, useState} from "react";

import {useParams} from 'react-router-dom'
import Navigation from "../navbar/Navigation";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import axios from "axios";
import QuizInfo from "./quiz-info/QuizInfo";
import HistoryTable from "./history-table/HistoryTable";
import PlayQuiz from "../playquiz/PlayQuiz";
import "../../style.css";
import {useTransition, animated} from "react-spring";

export function QuizDetail(){
    const { quizId } = useParams();
    const [quiz, setQuiz] = useState<QuizInterface>({});
    const [isPlay,setIsPlay] = useState(false);
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

    const transitions = useTransition(isPlay,null, {
        from: { opacity: 0 },
        enter: { opacity: 1},
        leave: { opacity: 0 , display:'none'},
    });

    return  <>{
        transitions.map(({item,key,props}) =>
            item
                ? <animated.div style={props} key={key}>
                    <PlayQuiz quiz={quiz}/>
                </animated.div>
                :
                <animated.div key={key} style={props} >
                    <main className="my-main-content">
                        <div className="my-container">
                            <div className="row">
                                <Navigation/>
                                <div className="col-xl-12">
                                    <Breadcrumbs>
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
                                            <button className="btn-play" style={{marginBottom: '184px'}} onClick={() => {
                                                setIsPlay(true);
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
                </animated.div>
        )
    }</>

}
export default QuizDetail;
