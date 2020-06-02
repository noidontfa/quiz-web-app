import React from "react";
import Ratings from "../../rating/Ratings";

interface P {
    quiz : QuizInterface
}

const QuizInfo : React.FC<P> = ({quiz}) => {
    return (
        <>
        <div className="col-xl-12">
            <img
                className="edit-quiz-image"
                src={quiz.image}
                alt="user-avartar"
            />
        </div>
        <div className="col-xl-12">
            <div className="quiz-detail-name" style={{marginBottom: '18px'}}>
                {quiz.name}
            </div>
        </div>
        <div className="col-xl-12">
            <div className="row">
                <div className="col-xl-6">
                    <div className="quiz-detail-description">
                        {quiz.description}
                    </div>
                </div>
                <div className="col-xl-6">
                    <div className="d-flex justify-content-end align-items-end" style={{height: '100%'}}>

                        <Ratings key={quiz.ratings} defaultValue={quiz.ratings!} readonly={true}/>
                    </div>
                </div>
            </div>
        </div>
        </>
    )
}

export default QuizInfo;