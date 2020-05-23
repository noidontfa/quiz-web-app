import React from "react";
import {CSSTransition} from "react-transition-group";
import NavigationPlay from "./navigation/NavigationPlay";
interface P {
    quiz : QuizInterface
}

const quizTest = {
    id: 42,
    name: "Quiz 1",
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
                    isRight: false,
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
            name: "Question 2",
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
                    name: "Choice 2",
                    isRight: false,
                    image: ""
                },
                {
                    id: 99,
                    name: "Choice 2",
                    isRight: false,
                    image: ""
                }
            ],
            totalQuestion: 2,
            image: ""
        },

    ],
}



const PlayQuiz : React.FC<P> = ({quiz}) => {




    return (
                <>
                    <NavigationPlay question={1} score={500} sec={15} totalQuestion={2}/>
                    <div className="quiz-content">
                        <div className="row">
                            <div className="content">
                                <div className="col-xl-12">
                                    <div className="row d-flex justify-content-center">
                                        <div className="col-xl-10">
                                            <div className="question">
                                                Let's Find Out With BTS Member You're Most Like Based On Your
                                                Favorite TV Shows? Let's Find Out With BTS Member You're Most
                                                Like Based On Your Favorite TV Shows? Let's Find Out With BTS
                                                Member You're Most Like Based On Your Favorite TV Shows?
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
                                        <div className="col-xl-3" style={{marginTop: '30px'}}>
                                            <div className="choice">
                                                <div className="choice-title">
                                                    1
                                                </div>
                                                <div className="choice-content">
                                                    Day la cau tra loi a
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-xl-3" style={{marginTop: '30px'}}>
                                            <div className="choice">
                                                <div className="choice-title">
                                                    1
                                                </div>
                                                <div className="choice-content">
                                                    Day la cau tra loi a
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-xl-3" style={{marginTop: '30px'}}>
                                            <div className="choice">
                                                <div className="choice-title">
                                                    1
                                                </div>
                                                <div className="choice-content">
                                                    Day la cau tra loi a
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-xl-3" style={{marginTop: '30px'}}>
                                            <div className="choice">
                                                <div className="choice-title">
                                                    1
                                                </div>
                                                <div className="choice-content">
                                                    Day la cau tra loi a
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div
                                    className="col-xl-12 d-flex justify-content-end align-items-center"
                                    style={{marginTop: '50px'}}
                                >
                                    <button className="btn-submit">Submit</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </>
    )
}

export default PlayQuiz;