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

    const onNextQuestion = () => {
        if (questionIndex + 1 > quiz.totalQuestion!) {
            alert("Finished");
        } else {
            setQuestionIndex(questionIndex + 1);
        }
    };


    return (
                <>
                    <NavigationPlay question={questionIndex} score={500} sec={15} totalQuestion={quiz.totalQuestion!}/>
                    {
                        quiz.questionRefer?.map((q,index) => {
                            if(index === questionIndex) {
                                return (<QuestionItem question={q} callBackFunction={onNextQuestion}/>);
                            }
                        })
                    }
                </>
    )
}

export default PlayQuiz;