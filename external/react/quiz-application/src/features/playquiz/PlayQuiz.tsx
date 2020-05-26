import React, {useEffect, useState} from "react";
import {CSSTransition} from "react-transition-group";
import NavigationPlay from "./navigation/NavigationPlay";
import ChoiceItem from "./choice-item/ChoiceItem";
import QuestionItem from "./question-item/QuestionItem";
interface P {
    quiz : QuizInterface
}





const PlayQuiz : React.FC<P> = ({quiz}) => {


    const [questionIndex,setQuestionIndex] = useState(0);
    const [score, setScore] = useState(100);
    const [showChoices, setShowChoices] = useState(false);
    const onNextQuestion = () => {
        //call back show right question => setitme setstate;
        setShowChoices(showChoices => !showChoices);
        setTimeout(() => {
            setShowChoices(showChoices => !showChoices);
            setQuestionIndex(questionIndex => questionIndex + 1);
            setScore(score => score + 100)
        },2000);
    };

    function OnRender() : JSX.Element {
        if(questionIndex !== (quiz.totalQuestion! + 1)) {
            return (
                <NavigationPlay  showChoices={showChoices} question={questionIndex} score={score} sec={5} totalQuestion={quiz.totalQuestion!} callbackFunction={onNextQuestion}/>
            )
        }
        return <h1>Finished</h1>;
    }

    return (
                <>
                    <OnRender/>
                    { quiz.questionRefer?.map((q,index) => {
                            if(index === questionIndex) {
                                return (<QuestionItem showChoices={showChoices} question={q} callBackFunction={onNextQuestion}/>);
                            }
                        })
                    }

                </>
    )
}

export default PlayQuiz;