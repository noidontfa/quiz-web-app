import React, {useEffect} from "react";
import {animated} from "react-spring";
import ChoiceItem from "./choice-item/ChoiceItem";
import Ratings from "../rating/Ratings";
import axios from "axios";
import {useHistory} from "react-router-dom";
import createHistory from 'history/createBrowserHistory';
interface P {
    quiz : QuizInterface;
    rightChoices: number;
    score : number;
}

const FinishPage  : React.FC<P> = ({quiz,rightChoices,score}) => {
    const history = useHistory();
    const onRatingQuiz = (value : number) => {
        axios.post(`http:/api/ratings/`, {
            userId: 2,
            quizId: quiz.id,
            star: value
        }).then(function (res) {
            alert(res.data);
        }).catch(function (err) {
            console.log(err);
        })
    }

    useEffect(() => {
        axios.post(`http:/api/histories/`, {
            userId: 2,
            quizId: quiz.id,
            score: score,
            numberRightAnswers: rightChoices
        }).then(function (res) {
            alert("Save cored");
        }).catch(function (err) {
            console.log(err);
        });
    })


    return ( <>
        <header className="play-quiz-navbar ">
            <div className="wrapper">
                <div className="wrap-item">
                    <div className="action-button">
						<span>
                            {rightChoices + ' / ' + quiz.totalQuestions }
						</span>
                    </div>
                </div>


                <div className="wrap-item">
                            <div className="action-button">
						                    <span>
                                                    {score}
                                            </span>
                            </div>
                </div>
            </div>
        </header>
        <div className="quiz-content">
            <div className="row">
                <div className="content">
                    <div className="col-xl-12">
                        <div className="row d-flex justify-content-center">
                            <div className="col-xl-10">
                                <div className="finish-content">
                                    <div className="title">
                                        <span>
                                          FINISHED
                                        </span>
                                    </div>
                                    <div className="f-content">
                                        <div className="block">
                                              <span>
                                                Right choices
                                              </span>
                                               <span>
                                                   {rightChoices + ' / ' + quiz.totalQuestions }
                                              </span>
                                        </div>
                                        <div className="block">
                                          <span>
                                            Scores
                                          </span>
                                            <span>
                                                {score}
                                          </span>
                                        </div>
                                    </div>
                                    <div className="f-ratings">
                                        <div className="r">
                                            <Ratings defaultValue={0} readonly={false} callbackFunction={onRatingQuiz}/>
                                        </div>
                                        <div className="b">
                                            <button className="btn-exit" onClick={() => history.push("/")}>Exit</button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </>)

}


export default FinishPage;