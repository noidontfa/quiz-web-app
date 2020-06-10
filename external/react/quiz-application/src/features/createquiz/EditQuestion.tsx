import React, {useEffect, useRef, useState} from "react";
import Navigation from "../navbar/Navigation";
import UserInfo from "./user-info/UserInfo";
import axios from "axios";
import {useParams} from "react-router-dom";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import Editable from "../editable-lable/Editable.";
const EditQuiz = () => {
    const inputRef = useRef() as React.MutableRefObject<HTMLTextAreaElement>;
    const { quizId } = useParams();
    const [quizName,setQuizName] = useState('');
    const [questions,setQuestions] = useState<Array<QuestionInterface>>([]);
    const doStuff =  async () => {
        try {
            const response = await axios.get(`http:/api/quizzes/${quizId}`);
            // console.log(response.data);
            const quiz : QuizInterface = response.data;
            // setQuiz(() => quizTest);
            // console.log(response.data);
            setQuizName(quiz.name!);
            setQuestions(quiz.questionRefer!);

        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        doStuff();
    },[])

    const onSaveQuestion = () => {
        axios.post(`http:/api/questions/${quizId}`,
            questions
        )
            .catch(err => alert(err));
    }



    return <>
        <Navigation/>

        <main className="my-main-content">
            <div className="my-container">
                <div className="row">
                    <UserInfo/>

                    <div className="col-xl-12" style={{margin: "40px 0 20px 0"}}>
                        <div className="seprator"></div>
                    </div>

                    <div className="col-xl-12">
                        <div className="row">
                            <div className="col-xl-6">
                                <Breadcrumbs>
                                    <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                                    <LinkBreadcrumbs to={'/'} active={true} name={'My quizzes'}/>
                                    <LinkBreadcrumbs to={`/quiz/edit/${quizId}`} active={true} name={quizName}/>
                                    <LinkBreadcrumbs active={false} name={'Question list'}/>
                                </Breadcrumbs>
                            </div>

                            <div className="col-xl-6 nav-edit-quiz">
                                <button className="quiz-detail">
                                    Quiz detail
                                </button>

                                <button className="questions active">
                                    Questions
                                </button>
                            </div>
                        </div>
                    </div>

                    <div className="col-xl-12" style={{marginTop: "15px"}}>
                        <div className="d-flex justify-content-between">
                            <button className="btn-create">
                                <i className="icon-folder"></i>
                                <span>Import csv</span>
                            </button>
                        </div>
                    </div>

                    <div className="col-xl-12">
                        <table className="table my-table">
                            <thead>
                            <tr>
                                <th scope="col" row-data="index">#</th>
                                <th scope="col" row-data="questions_name">Question name</th>
                                <th scope="col" row-data="choices">Choices</th>
                                <th scope="col" row-data="is_right">Is right</th>
                            </tr>
                            </thead>
                            <tbody>
                            {questions.map((q,i) => (
                                <tr>
                                    <th scope="row">{i + 1}</th>
                                    <td>
                                        <div className='table-content'>
                                            {/*<label className="text-table-content">{q.name}</label>*/}
                                            <Editable
                                                text={q.name}
                                                placeholder="Write a task name"
                                                type="textarea"
                                                childRef={inputRef}
                                                callBackFunction={onSaveQuestion}
                                            >
                                                <textarea
                                                    placeholder="Write a task name"
                                                    value={q.name}
                                                    data-id={q.id}
                                                    style={{width: "100%"}}
                                                    onChange={e => {
                                                       const questionId = Number(e.target.getAttribute("data-id"));
                                                       setQuestions( [...questions].map(object => {
                                                           if(object.id === questionId)
                                                               return {
                                                                   ...object,
                                                                   name: e.target.value
                                                               }
                                                           else return object;
                                                       }))
                                                    }}
                                                    rows={4}
                                                    ref={inputRef}
                                                />
                                            </Editable>
                                            {/*<div className="table-action">*/}
                                                {/*    <i className="icon-note"></i>*/}
                                                {/*    <i className="icon-trash"></i>*/}
                                                {/*</div>*/}
                                            {/*</input>*/}
                                        </div>
                                    </td>
                                    <td colSpan={2}>
                                        {q.choices?.map(e => (
                                            <div className="table-content choice-content">
                                                    {<Editable
                                                        text={e.name}
                                                        placeholder="Write a task name"
                                                        type="textarea"
                                                        childRef={inputRef}
                                                        callBackFunction={onSaveQuestion}
                                                    >
                                                <textarea
                                                    placeholder="Write a task name"
                                                    value={e.name}
                                                    data-id={e.id}
                                                    style={{width: "100%"}}
                                                    onChange={c => {
                                                        const choiceId = Number(c.target.getAttribute("data-id"));
                                                        setQuestions([...questions].map(object => {
                                                            return {
                                                                ...object,
                                                                choices: object.choices?.map(ch => {
                                                                    if(ch.id === choiceId) {
                                                                        return {
                                                                            ...ch,
                                                                            name: c.target.value
                                                                        }
                                                                    } else
                                                                    return ch;
                                                                })
                                                            }
                                                        }))
                                                    }}
                                                    rows={4}
                                                    ref={inputRef}
                                                />
                                                    </Editable>}
                                                    {/*<div className="table-action">*/}
                                                    {/*    <i className="icon-note"></i>*/}
                                                    {/*    <i className="icon-trash"></i>*/}
                                                    {/*</div>*/}

                                                <input
                                                    type="checkbox"
                                                    defaultChecked={e.isRight}
                                                    className="table-checkbox"
                                                    data-id={e.id}
                                                    onChange={c => {
                                                        const choiceId = Number(c.target.getAttribute("data-id"));
                                                        const questionD = [...questions].map(object => {
                                                            return {
                                                                ...object,
                                                                choices: object.choices?.map(ch => {
                                                                    if(ch.id === choiceId) {
                                                                        return {
                                                                            ...ch,
                                                                            isRight: !ch.isRight
                                                                        }
                                                                    } else
                                                                        return ch;
                                                                })
                                                            }
                                                        });
                                                        axios.post(`http:/api/questions/${quizId}`,
                                                            questionD
                                                        )
                                                            .catch(err => alert(err));
                                                    }}
                                                />
                                            </div>
                                        ))}
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </main>
    </>
}

export default EditQuiz;