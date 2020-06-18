import React, {useEffect, useState} from "react";
import {CSSTransition} from "react-transition-group";
import NavigationPlay from "./navigation/NavigationPlay";
import ChoiceItem from "./choice-item/ChoiceItem";
import QuestionItem from "./question-item/QuestionItem";
import FinishPage from "./FinishPage";
import Cookies  from 'universal-cookie';
import axios from "axios";

interface P {
    quiz : QuizInterface
}





const PlayQuiz : React.FC<P> = ({quiz}) => {

    const [questionIndex,setQuestionIndex] = useState(0);
    const [score, setScore] = useState(100);
    const [rightChoices,setRightChoices] = useState(0);
    const [showChoices, setShowChoices] = useState(false);
    const [user,setUser] = useState<UserInterface>({});
    const cookies = new Cookies();
    useEffect(() => {
        const token = cookies.get("token");
        if ( token != undefined ){
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

    const onNextQuestion = () => {
        //call back show right question => setitme setstate;
        const selection = document.getElementsByClassName("choice active");
        let choices : Array<number> = [];
        for(let i = 0; i < selection.length; i++) {
           choices.push(Number(selection.item(i)!.getAttribute("data-id")));
        }
        console.log(choices);
        let currentScore = score;
        if(choices.length) {
            const question = quiz.questionRefer?.filter((value,index) => {
                return index === questionIndex;
            })

            question?.forEach(q => {
                const result = q.choices?.filter((e) => e.isRight).map(e => e.id);
                console.log(result);
                console.log(choices);
                const isRight = choices.some(e => result?.includes(e))
                if(isRight && result?.length === choices.length) {
                    currentScore += 100;
                    setRightChoices(rightChoices => rightChoices + 1);
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
        if(questionIndex !== (quiz.totalQuestions!)) {
            return (
                <NavigationPlay key={quiz.id} showChoices={showChoices} question={questionIndex + 1} score={score} sec={quiz.timingRefer?.sec!} totalQuestion={quiz.totalQuestions!} callbackFunction={onNextQuestion}/>
            )
        }
        return <FinishPage key={quiz.id} quiz={quiz} rightChoices={rightChoices} score={score} user={user}/>
    }

    return (
                <>
                    <OnRender/>
                    { quiz.questionRefer?.map((q,index) => {
                            if(index === questionIndex) {
                                return (<QuestionItem key={q.id} showChoices={showChoices} question={q} callBackFunction={onNextQuestion}/>);
                            }
                        })
                    }

                </>
    )
}

export default PlayQuiz;