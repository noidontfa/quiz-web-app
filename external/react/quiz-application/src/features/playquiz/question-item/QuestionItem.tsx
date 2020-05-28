import React from "react";
import ChoiceItem from "../choice-item/ChoiceItem";

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

    return (
        <div className="quiz-content">
            <div className="row">
                <div className="content">
                    <div className="col-xl-12">
                        <div className="row d-flex justify-content-center">
                            <div className="col-xl-10">
                                <div className="question">
                                    {question.name}
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
                        <button className="btn-submit" id="btn-submit" onClick={onSubmit} style={{pointerEvents: "none"}}>Submit</button>
                    </div>
                </div>
            </div>
        </div>
    )

}

export default QuestionItem;