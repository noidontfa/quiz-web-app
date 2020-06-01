import React, {useEffect} from "react";
import ChoiceItem from "../choice-item/ChoiceItem";
import {ButtonBaseActions} from "@material-ui/core";
import {animated, useTransition} from "react-spring";
import Navigation from "../../navbar/Navigation";
import Breadcrumbs from "../../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../../breadcrumb/LinkBreadcrumbs";
import QuizInfo from "../../quiz-detail/quiz-info/QuizInfo";
import HistoryTable from "../../quiz-detail/history-table/HistoryTable";

interface P {
    question : QuestionInterface;
    callBackFunction: () => void;
    showChoices : boolean;
}

const QuestionItem : React.FC<P> = ({question,callBackFunction,showChoices}) => {

    const onChooseChoice = (id : number) => {
        document.querySelector(`div[data-id='${id}']`)!.classList.toggle("active");
        if(document.getElementsByClassName("choice active").length) {
            document.getElementById("btn-submit")!.style.pointerEvents = "all"
        } else {
            document.getElementById("btn-submit")!.style.pointerEvents = "none"
        }
    }

    const onSubmit = () => {
        callBackFunction();
    }
    const onKey = (ev : KeyboardEvent) => {
        if(ev.keyCode === 13) {
            const ob = document.getElementById("btn-submit")!.style.pointerEvents;
            if(ob === "all") {
                callBackFunction();
            }
        }
    };

    const transitions = useTransition(showChoices, null ,{
        from: { transform: 'translate3d(0,-40px,0)', opacity: 0},
        enter: { transform: 'translate3d(0,0px,0)', opacity: 1},
        leave: { display: 'none', opacity: 0 },
    });


    useEffect(() => {
        document.addEventListener("keyup", onKey )
        return () => {
            document.removeEventListener("keyup",onKey);
        }
    },[])

    return (
        <div className="quiz-content">
            <div className="row">
                <div className="content">
                    <div className="col-xl-12">
                        <div className="row d-flex justify-content-center">
                            <div className="col-xl-10">
                                <div className="question">

                                    {
                                        transitions.map(({item,key,props}) =>
                                            <animated.div key={key} style={props} >
                                                {question.name}
                                            </animated.div>
                                        )
                                    }
                                </div>
                            </div>
                        </div>
                    </div>
                    {/*// <!-- <div class="col-xl-12">*/}
                    {/*//         <div class="row d-flex justify-content-center">*/}
                    {/*//             <div class="col-xl-10" style="margin-top: 50px;">*/}
                    {/*//                 <div class="question-img">*/}
                    {/*//*/}
                    {/*//                 </div>*/}
                    {/*//             </div>*/}
                    {/*//         </div>*/}
                    {/*//     </div> -->*/}
                    <div className="col-xl-12" style={{flexGrow : 1}}>
                        <div className="row d-flex justify-content-center">
                            {question.choices?.map((choice,index) =>
                                <ChoiceItem showChoices={showChoices} index={index} choice={choice} callBackFunction={onChooseChoice}/>
                            )}
                        </div>
                    </div>
                    <div
                        className="col-xl-12 d-flex justify-content-end align-items-center"
                        style={{marginTop: '50px'}}
                    >
                        <button className="btn-submit" id="btn-submit" onClick={onSubmit} style={{pointerEvents: showChoices ?  "none" : "all"}}>Submit</button>
                    </div>
                </div>
            </div>
        </div>
    )

}

export default QuestionItem;