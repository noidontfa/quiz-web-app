import React, {useEffect, useState} from "react";
import Navigation from "../navbar/Navigation";
import UserInfo from "./user-info/UserInfo";
import Breadcrumbs from "../breadcrumb/Breadcrumbs";
import LinkBreadcrumbs from "../breadcrumb/LinkBreadcrumbs";
import axios from "axios";
import {useHistory} from "react-router-dom";


const CreateQuiz = () => {
    const history = useHistory();
    const [quizName,setQuizName] = useState('');
    const [description,setDescription] = useState('');
    const [categoryId,setCategoryId] = useState(0);
    const [timingId,setTimingId] = useState(0);
    const [languageId,setLanguageId] = useState(0);

    const [dataCategories,setDataCategories] = useState<Array<CategoryInterface>>([]);
    const [dataTimings,setDataTimings] = useState<Array<TimingInterface>>([]);
    const [dataLanguages,setDataLanguages] = useState<Array<LanguageInterface>>([]);

    useEffect(() => {
        axios.get("http:/api/categories").then(res => {
            setDataCategories(res.data);
            setCategoryId(res.data[0].id)
        })
            .catch(err => console.log(err));
        axios.get("http:/api/timings").then(res => {
            setDataTimings(res.data);
            setTimingId(res.data[0].id);
        })
            .catch(err => console.log(err));
        axios.get("http:/api/languages").then(res => {
            setDataLanguages(res.data);
            setLanguageId(res.data[0].id);
        })
            .catch(err => console.log(err));


    },[]);

    const onCreateQuiz = (e : React.SyntheticEvent) => {
        e.preventDefault();
        if(quizName === '') {
            alert("quizName is empty");
        } else {
            axios.post("http:/api/quizzes", {
                name: quizName,
                description,
                categoryId,
                timingId,
                languageId,
                createdBy: 2,
            }).then(res => {
                console.log(res.data);
                const path = `edit/${res.data.id}`
                history.push(path);
            }).catch(err => console.log(err));
        }
    }

    return (
        <>
            <Navigation/>
            <main className="my-main-content">
                <div className="my-container">
                    <div className="row">
                        <UserInfo/>

                        <div className="col-xl-12" style={{margin: "40px 0 20px 0"}}>
                            <div className="seprator"></div>
                        </div>

                        <div className="col-xl-12">
                            <Breadcrumbs>
                                <LinkBreadcrumbs to={'/'} active={true} name={'Home'}/>
                                <LinkBreadcrumbs to={'/quiz/my'} active={true} name={'My quizzes'}/>
                                <LinkBreadcrumbs active={false} name={'Create quiz'}/>
                            </Breadcrumbs>
                        </div>

                        <div className="col-xl-12" style={{marginTop: "40px"}}>
                            <form>
                                <div className="row justify-content-center">
                                    <div className="col-xl-8">
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
                                    <div className="col-xl-8">
                                        <div className="my-form-group">
                                            <label htmlFor="name">Description</label>
                                            <textarea
                                                className="my-form-control"
                                                rows={4}
                                                name="description"
                                                placeholder="Description"
                                                value={description}
                                                onChange={e => setDescription(e.target.value)}
                                            ></textarea>
                                        </div>
                                    </div>

                                    <div className="col-xl-8">
                                        <div className="row">
                                            <div className="col-xl-4">
                                                <div className="my-form-group">
                                                    <label htmlFor="category">Category</label>
                                                    <select
                                                        className="my-form-control"
                                                        id="category"
                                                        name="category"
                                                        onChange={e => setCategoryId(Number(e.target.value))}
                                                    >
                                                        {dataCategories.map((e,i) =>
                                                            <option key={e.id} value={e.id}>{e.name}</option>
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
                                                    >
                                                        {dataLanguages.map((e,i) =>
                                                            <option key={e.id} value={e.id}>{e.name}</option>
                                                        )}
                                                    </select>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="col-xl-8">
                                        <div className="d-flex justify-content-end">
                                            <button
                                                type="submit"
                                                className="btn-cancel d-flex justify-content-center"
                                                style={{marginRight: "30px"}}
                                            >
                                                <i className="icon-close"></i>
                                                <span>Cancel</span>
                                            </button>
                                            <button type="submit" className="btn-create" onClick={onCreateQuiz}>
                                                <i className="icon-rocket"></i>
                                                <span>Create quiz</span>
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </main>
        </>
    )
}

export default CreateQuiz;