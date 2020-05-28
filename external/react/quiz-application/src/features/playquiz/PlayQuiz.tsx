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
        const selection = document.getElementsByClassName("choice active");
        let choices : Array<Number> = [];
        for(let i = 0; i < selection.length; i++) {
           choices.push(Number(selection.item(i)!.getAttribute("data-id")));
        }
        let currentScore = score;
        if(choices.length) {
            const question = quiz.questionRefer?.filter((value,index) => {
                return index === questionIndex;
            })

            question?.forEach(q => {
                const result = q.choices?.filter((e) => {
                    return choices.includes(e.id!) && e.isRight
                })
                if(result!.length === choices.length) {
                    currentScore += 100;
                }
            })
        }


        setShowChoices(showChoices => !showChoices);
        setTimeout(() => {
            setShowChoices(showChoices => !showChoices);
            setQuestionIndex(questionIndex => questionIndex + 1);
            setScore(() => currentScore);
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