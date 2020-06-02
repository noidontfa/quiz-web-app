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



const quizTest : QuizInterface = {
    id: 42,
    name: "Quiz 1",
    totalQuestions: 3,
    ratings: 3.5,
    timingRefer: {
        id: 2,
        name:"",
        sec: 5
    },
    questionRefer: [
        {
            id: 45,
            name: "Question 1",
            choices: [
                {
                    id: 92,
                    name: "Choice 1",
                    isRight: true,
                    image: ""
                },
                {
                    id: 93,
                    name: "Choice 2",
                    isRight: true,
                    image: ""
                }
            ],
            image: ""
        },
        {
            id: 46,
            name: "Question 2",
            choices: [
                {
                    id: 94,
                    name: "Choice 1",
                    isRight: true,
                    image: ""
                },
                {
                    id: 95,
                    name: "Choice 2",
                    isRight: false,
                    image: ""
                }
            ],
            image: ""
        },
        {
            id: 46,
            name: "Question 3",
            choices: [
                {
                    id: 96,
                    name: "Choice 1",
                    isRight: true,
                    image: ""
                },
                {
                    id: 97,
                    name: "Choice 2",
                    isRight: false,
                    image: ""
                },
                {
                    id: 98,
                    name: "Choice 3",
                    isRight: true,
                    image: ""
                },
                {
                    id: 99,
                    name: "Choice 4",
                    isRight: false,
                    image: ""
                }
            ],
            image: ""
        },

    ],
}


export function QuizDetail(){
    const { quizId } = useParams();
    const [quiz, setQuiz] = useState<QuizInterface>({});
    const [isPlay,setIsPlay] = useState(false);
    const doStuff =  async () => {
        console.log(quiz);
        try {
            // const response = await axios.get(`http:/api/quizzes/${quizId}`);
            // console.log(response.data);
            // setQuiz(response.data);
            setQuiz(() => quizTest);
            // console.log(response.data);
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
                                        <HistoryTable quiz={quiz}/>
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
