import React, {useEffect, useRef, useState} from "react";
import Navigation from "../navbar/Navigation";
import UserInfo from "./user-info/UserInfo";
import axios from "axios";
import {useParams} from "react-router-dom";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import Editable from "../editable-lable/Editable.";
import {useHistory} from "react-router-dom";
import Cookies from 'universal-cookie';

const EditQuiz = () => {
    const history = useHistory();
    const cookies = new Cookies();
    const inputRef = useRef() as React.MutableRefObject<HTMLTextAreaElement>;
    const { quizId } = useParams();
    const [quizName,setQuizName] = useState('');
    const [questions,setQuestions] = useState<Array<QuestionInterface>>([]);
    const [user,setUser] = useState<UserInterface>({});

    const [questionName,setQuestionName] = useState('');
    const [choiceName,setChoiceName] = useState('');
    const [choiceIsRight,setChoiceIsRight] = useState(false);
    const [csvValue,setCSVValue] = useState('');

    const [questionId,setQuestionId] = useState(0);
    const doStuff =  async () => {
        try {
            const response = await axios.get(`http:/api/quizzes/${quizId}`);
            // console.log(response.data);
            const quiz : QuizInterface = response.data;
            // setQuiz(() => quizTest);
            // console.log(response.data);
            setQuizName(quiz.name!);
            setQuestions(quiz.questionRefer!);
            setQuestionId(quiz.questionRefer![0].id!);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        doStuff();

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

    const onSaveQuestion = () => {
        axios.post(`http:/api/questions/${quizId}`,
            questions
        )
            .catch(err => alert(err));
    }
    const onDeleteQuestion = (questionId : Number) => {
        const ques = questions.filter(e => e.id === questionId)[0];
        const flag = window.confirm("Are you sure to delete this questions: " + ques.name + " ?");
        if(flag) {
            axios.delete(`http:/api/questions/${questionId}`)
                .then(res => {
                    const questionF = [...questions].filter(e => e.id !== questionId);
                    setQuestions(questionF);
                })
                .catch(err => alert(err))
        }

    }
    const onDeleteChoice = (choice : ChoiceInterface) => {
        const flag = window.confirm("Are you sure to delete this choice: " + choice.name + " ?");
        if (flag) {
            axios.delete(`http:/api/choices/${choice.id}`)
                .then(res => {
                    const questionD = [...questions].map(object => {
                       return {
                           ...object,
                           choices: object.choices?.filter(ch => ch.id !== choice.id)
                       }
                    })
                    setQuestions(questionD);
                })
                .catch(err => alert(err))
        }
    }

    const onCreateQuestion = () => {
        if (questionName === '') {
            alert("Question is Null");
            return
        }
        axios.post(`http:/api/questions/${quizId}`, [{
            name: questionName
        }]).then(res => {
                const questionD = [...questions,res.data[0]];
                setQuestions(questionD);
            }
        )

    }

    const onCreateChoice = () => {
        if (questionId === 0 || choiceName === '') {
            alert("Not enough field");
            return;
        }
        const data = [{
            id: questionId,
            choices: [
                {
                    name: choiceName,
                    isRight: choiceIsRight
                }
            ]
        }];
        axios.post(`http:/api/questions/${quizId}`, data).then(res => {
            const ques = res.data[0];
            const questionD = [...questions].map(q => {
                if (q.id === questionId) {
                    return {
                        ...q,
                        choices: ques.choices,
                    }
                } else
                    return q;
            })
            setQuestions(questionD);
        })
    }

    const onCreateMultiData = () => {
        const lines = csvValue.split('\n');
        let questionsQ : Array<any> = [];
        lines.forEach(line => {
            const datas = line.split('\t');
            const question = datas[0];
            let choices = [];
            for(let i = 1 ; i < datas.length - 1; i++) {
                const choice = datas[i];
                const isRight = datas[i+1] === '1' ? true : false;
                i++;
                choices.push({
                    name: choice,
                    isRight
                })
            }
            if (line.length) {
                questionsQ.push({
                    name: question,
                    choices
                })
            }
        })
        if (questionsQ.length) {
            axios.post(`http:/api/questions/${quizId}`, questionsQ).then(res => {
                    const questionD = [...questions,...res.data];
                    setQuestions(questionD);
                }
            )
        }
    }



    return <>
        <Navigation user={user}/>

        <main className="my-main-content">
            <div className="my-container">
                <div className="row">
                    <UserInfo user={user}/>

                    <div className="col-xl-12" style={{margin: "40px 0 20px 0"}}>
                        <div className="seprator"></div>
                    </div>

                    <div className="col-xl-12">
                        <div className="row">
                            <div className="col-xl-6">
                                <Breadcrumbs>
                                    <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                                    <LinkBreadcrumbs to={'/quiz/myn'} active={true} name={'My quizzes'}/>
                                    <LinkBreadcrumbs to={`/quiz/edit/${quizId}`} active={true} name={quizName}/>
                                    <LinkBreadcrumbs active={false} name={'Question list'}/>
                                </Breadcrumbs>
                            </div>

                            <div className="col-xl-6 nav-edit-quiz">
                                <button className="quiz-detail" onClick={() => {
                                    const path = `/quiz/edit/${quizId}`;
                                    history.push(path);
                                }}>
                                    Quiz detail
                                </button>

                                <button className="questions active">
                                    Questions
                                </button>
                            </div>
                        </div>
                    </div>

                    <div className="col-xl-12" style={{marginTop: "15px"}}>
                        <div className="d-flex justify-content-start">
                            <button className="btn-upload"
                                    style={{ marginRight: "5px"}}
                                    data-toggle="modal"
                                    data-target="#questionModel"
                                    onClick={() => setQuestionName("")}
                            >
                                <i className="icon-plus"></i>
                                <span>Question</span>
                            </button>

                            <button className="btn-save-changes"
                                    style={{ marginRight: "5px"}}
                                    data-toggle="modal"
                                    data-target="#choiceModal"
                                    onClick={() => {setChoiceName(""); setChoiceIsRight(false) ; setQuestionId(questions[0]!.id!)}}
                            >
                                <i className="icon-plus"></i>
                                <span>Choice</span>
                            </button>

                            <button className="btn-create"
                                    data-toggle="modal"
                                    data-target="#csvModal"
                            >
                                <i className="icon-folder"></i>
                                <span>Import</span>
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
                                                callBackDeleteFunction={() => onDeleteQuestion(q.id!)}
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
                                                        callBackDeleteFunction={() => onDeleteChoice(e)}
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
                                                        ).then(() => setQuestions(questionD)

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


            <div className="modal fade" id="questionModel" tabIndex={-1} role="dialog"
                 aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div className="modal-dialog modal-dialog-centered modal-lg" role="document">
                    <div className="modal-content">
                        {/*<div className="modal-header">*/}
                        {/*    <h5 className="modal-title" id="exampleModalLabel">Modal title</h5>*/}
                        {/*    <button type="button" className="close" data-dismiss="modal" aria-label="Close">*/}
                        {/*        <span aria-hidden="true">&times;</span>*/}
                        {/*    </button>*/}
                        {/*</div>*/}
                        <div className="modal-body">
                            <div className="my-form-group">
                                <label className="lable" htmlFor="name">Question Name</label>
                                <textarea
                                    className="my-form-control"
                                    id="createQuestion"
                                    name="name"
                                    rows={10}
                                    placeholder="Type..."
                                    value={questionName}
                                    onChange={e => setQuestionName(e.target.value)}
                                />
                            </div>
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn-create" data-dismiss="modal" onClick={onCreateQuestion}>
                                <i className="icon-plus"></i>
                                Create
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <div className="modal fade" id="choiceModal" tabIndex={-1} role="dialog"
                 aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div className="modal-dialog modal-dialog-centered modal-lg" role="document">
                    <div className="modal-content">
                        {/*<div className="modal-header">*/}
                        {/*    <h5 className="modal-title" id="exampleModalLabel">Modal title</h5>*/}
                        {/*    <button type="button" className="close" data-dismiss="modal" aria-label="Close">*/}
                        {/*        <span aria-hidden="true">&times;</span>*/}
                        {/*    </button>*/}
                        {/*</div>*/}
                        <div className="modal-body">
                            <div className="my-form-group">
                                <label className="lable">Question</label>
                                <select
                                    className="my-form-control"
                                    value={questionId}
                                    onChange={e => setQuestionId(Number(e.target.value))}
                                >
                                    {
                                        questions.map((e,i) =>
                                        <option key={e.id} value={e.id}>{i + 1 +" - " +e.name}</option>
                                    )}
                                </select>

                                <label className="lable" htmlFor="#createChoice" style={{marginTop: "5px"}}>Choice Name</label>
                                <textarea
                                    className="my-form-control"
                                    id="createChoice"
                                    rows={10}
                                    placeholder="Type..."
                                    value={choiceName}
                                    onChange={e => setChoiceName(e.target.value)}
                                />
                                <label className="lable" style={{display: "inline"}} >Is right:</label>
                                <input
                                    style={{marginLeft: "10px"}}
                                    type="checkbox"
                                    checked={choiceIsRight}
                                    className="table-checkbox"
                                    onChange={c => {
                                        setChoiceIsRight(!choiceIsRight);
                                    }}
                                />
                            </div>
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn-create" data-dismiss="modal" onClick={onCreateChoice}>
                                <i className="icon-plus"></i>
                                Create
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <div className="modal fade" id="csvModal" tabIndex={-1} role="dialog"
                 aria-labelledby="csvModal" aria-hidden="true">
                <div className="modal-dialog modal-dialog-centered modal-lg" role="document">
                    <div className="modal-content">
                        {/*<div className="modal-header">*/}
                        {/*    <h5 className="modal-title" id="exampleModalLabel">Modal title</h5>*/}
                        {/*    <button type="button" className="close" data-dismiss="modal" aria-label="Close">*/}
                        {/*        <span aria-hidden="true">&times;</span>*/}
                        {/*    </button>*/}
                        {/*</div>*/}
                        <div className="modal-body">
                            <div className="my-form-group">
                                <label className="label">Rule: 1 is true, 0 is false</label>
                                <label className="label">Example: question choice1 1 choice2 0 choice3 0 choice4 0</label>
                                <label className="label" htmlFor="name">Copy csv format here:</label>
                                <textarea
                                    className="my-form-control"
                                    id="createQuestion"
                                    name="name"
                                    rows={10}
                                    placeholder="Type..."
                                    value={csvValue}
                                    onChange={e => {
                                        setCSVValue(e.target.value);
                                    }}
                                />
                            </div>
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn-create" data-dismiss="modal" onClick={onCreateMultiData}>
                                <i className="icon-folder   "></i>
                                Import
                            </button>
                        </div>
                    </div>
                </div>
            </div>

        </main>
    </>
}

export default EditQuiz;