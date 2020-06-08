import React, {useEffect, useState} from "react";
import axios from "axios";
import {useParams} from "react-router-dom";
import Navigation from "../navbar/Navigation";
import UserInfo from "./user-info/UserInfo";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";


const EditQuiz = () => {


    const { quizId } = useParams();
    const [quizName,setQuizName] = useState('');
    const [description,setDescription] = useState('');
    const [categoryId,setCategoryId] = useState(0);
    const [timingId,setTimingId] = useState(0);
    const [languageId,setLanguageId] = useState(0);
    const [stateId,setStateId] = useState(0);
    const [image,setImage] = useState("");
    const [filename,setFilename] = useState("");


    const [dataCategories,setDataCategories] = useState<Array<CategoryInterface>>([]);
    const [dataTimings,setDataTimings] = useState<Array<TimingInterface>>([]);
    const [dataLanguages,setDataLanguages] = useState<Array<LanguageInterface>>([]);
    const [dataStates,setDataStates] = useState<Array<StateInterface>>([]);



    const doStuff =  async () => {
        try {
            const response = await axios.get(`http:/api/quizzes/${quizId}`);
            // console.log(response.data);
            const quiz : QuizInterface = response.data;
            setQuizName(quiz.name!);
            setDescription(quiz.description!);
            setCategoryId(quiz.categoryRefer?.id!);
            setTimingId(quiz.timingRefer?.id!);
            setLanguageId(quiz.languageRefer?.id!);
            setImage(quiz.image!);
            setStateId(quiz.stateRefer?.id!);
            // setQuiz(() => quizTest);
            // console.log(response.data);


        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        doStuff();
        axios.get("http:/api/categories").then(res => {
            setDataCategories(res.data);
        })
            .catch(err => console.log(err));
        axios.get("http:/api/timings").then(res => {
            setDataTimings(res.data);
        })
            .catch(err => console.log(err));
        axios.get("http:/api/languages").then(res => {
            setDataLanguages(res.data);
        })
            .catch(err => console.log(err));

        axios.get("http:/api/states").then(res => {
            setDataStates(res.data);
        })
            .catch(err => console.log(err));

    },[])

    const toBase64 = (file : any) => new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result);
        reader.onerror = error => reject(error);
    });

    const onSelectImage = async (e : React.SyntheticEvent) => {
        e.preventDefault();
        const imageDom = document.getElementById("image") as HTMLInputElement;
        const baseFile = imageDom.files![0];
        const fileBase64 =  await toBase64(baseFile);
        setImage(fileBase64 as string);
        setFilename(baseFile.name)
    }

    const onEditQuiz = (e : React.SyntheticEvent) => {
        e.preventDefault();
        if(quizName === '') {
            alert("quizName is empty");
        } else {
            axios.put(`http:/api/quizzes/${quizId}`, {
                name: quizName,
                description,
                categoryId,
                timingId,
                languageId,
                stateId,
                filename,
                image: image.split(',')[1],
                createdBy: 2
            }).then(res => {
                console.log(res.data);
                alert("Me log")
                window.location.reload();
            }).catch(err => console.log(err));
            console.log(quizName);
        }
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

                    </div>
                    <div className="col-xl-12">
                        <div className="row">
                            <div className="col-xl-6 d-flex align-items-center">
                                <Breadcrumbs>
                                    <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                                    <LinkBreadcrumbs to={'/'} active={true} name={'My quizzes'}/>
                                    <LinkBreadcrumbs active={false} name={'Quiz Application 1'}/>
                                </Breadcrumbs>
                            </div>
                            <div className="col-xl-6 nav-edit-quiz">
                                <div className="my-form-group">
                                    <label htmlFor="state">State</label>
                                    <select
                                        className="my-form-control"
                                        id="state"
                                        name="state"
                                        value={stateId}
                                        onChange={e => setStateId(Number(e.target.value))}
                                    >
                                        {dataStates.map((e,i) =>
                                            <option key={e.id} value={e.id}>{e.name}</option>
                                        )}
                                    </select>
                                </div>
                                <button className="quiz-detail active">
                                    Quiz detail
                                </button>

                                <button className="questions">
                                    Questions
                                </button>
                            </div>
                        </div>
                    </div>

                    <div className="col-xl-12" style={{marginTop: "40px"}}>
                        <div className="row">
                            <div className="col-xl-8">
                                <div className="row">
                                    <div className="col-xl-12">
                                        <img
                                            className="edit-quiz-image"
                                            src={image}
                                            alt="user-avartar"
                                        />
                                    </div>
                                    <div className="col-xl-12">
                                        <div className="my-form-group">
                                            <label className="label-requried" htmlFor="name">Quiz name</label>
                                            <input
                                                type="text"
                                                className="my-form-control"
                                                id="name"
                                                name="name"
                                                placeholder="Quiz name"
                                                value={quizName}
                                                onChange={e => setQuizName(e.target.value)}
                                            />
                                        </div>
                                    </div>

                                    <div className="col-xl-12">
                                        <div className="my-form-group">
                                            <label htmlFor="name">Description</label>
                                            <textarea
                                                className="my-form-control"
                                                rows={4}
                                                name="description"
                                                placeholder="Description"
                                                value={description}
                                                onChange={e => setDescription(e.target.value)}
                                            />

                                        </div>
                                    </div>

                                    <div className="col-xl-12">
                                        <div className="row">
                                            <div className="col-xl-4">
                                                <div className="my-form-group">
                                                    <label htmlFor="category">Category</label>
                                                    <select
                                                        className="my-form-control"
                                                        id="category"
                                                        name="category"
                                                        onChange={e => setCategoryId(Number(e.target.value))}
                                                        value={categoryId}
                                                    >
                                                        {dataCategories.map((e,i) =>
                                                            <option key={e.id} value={e.id} >{e.name}</option>
                                                        )}
                                                    </select>
                                                </div>
                                            </div>
                                            <div className="col-xl-4">
                                                <div className="my-form-group">
                                                    <label htmlFor="timing">Timing</label>
                                                    <select
                                                        className="my-form-control"
                                                        id="timing"
                                                        name="timing"
                                                        onChange={e => setTimingId(Number(e.target.value))}
                                                        value={timingId}
                                                    >
                                                        {dataTimings.map((e,i) =>
                                                            <option key={e.id} value={e.id}>{e.name}</option>
                                                        )}
                                                    </select>
                                                </div>
                                            </div>
                                            <div className="col-xl-4">
                                                <div className="my-form-group">
                                                    <label htmlFor="language">Language</label>
                                                    <select
                                                        className="my-form-control"
                                                        id="language"
                                                        name="language"
                                                        onChange={e => setLanguageId(Number(e.target.value))}
                                                        value={languageId}
                                                    >
                                                        {dataLanguages.map((e,i) =>
                                                            <option key={e.id} value={e.id}>{e.name}</option>
                                                        )}
                                                    </select>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="col-xl-12">
                                        <div className="d-flex justify-content-end">
                                            <button type="submit" className="btn-save-changes" onClick={onEditQuiz}>
                                                <i className="icon-folder"></i>
                                                <span>Save quiz</span>
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="col-xl-4">
                                <div className="row">
                                    <div className="col-xl-12 d-flex justify-content-center">
                                        <label className="btn-upload" style={{marginBottom: "184px"}}>
                                            <input type="file" id="image" style={{display: "none"}}  onChange={onSelectImage}/>
                                            <i className="icon-cloud-download"></i>
                                            <span>Upload image</span>
                                        </label>
                                    </div>
                                    <div className="col-xl-12 d-flex justify-content-center">
                                        <button type="submit" className="btn-delete">
                                            <i className="icon-trash"></i>
                                            <span>Delete quiz</span>
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </>
}

export default EditQuiz;