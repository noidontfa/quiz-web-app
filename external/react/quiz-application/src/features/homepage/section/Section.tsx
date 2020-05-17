import React, {EffectCallback, useEffect, useState} from "react";
import userIcon from "../../../assets/cat.jpg";

import axios from 'axios';
import BigCard from "./card/BigCard";
import MediumCard from "./card/MediumCard";


const Section = () => {
    const [quizData,setQuizData] = useState<Array<QuizInterface>>([]);

    const doStuff = async () => {
        try {
            const response = await axios.get("http:/api/quizzes");
            console.log(response.data);
            setQuizData(response.data);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(  () => { doStuff()},[])


    return (

        <div className="my-section" id="my-section">
            {
                quizData
                    .filter((e,i) => i === 0)
                    .map(e => (<BigCard quiz={e}/>))
            }
            {

                quizData.map((e,i) => {
                    if(i !== 0) {
                        return  (<MediumCard key= {e.id} quiz={e}/>)
                    }
                })
            }
            {/*<BigCard/>*/}
            {/*<MediumCard/>*/}
            {/*<MediumCard/>*/}
            {/*<MediumCard/>*/}
        </div>
    )
}

export default Section;