import React from "react";
import userIcon from "../../../../assets/cat.jpg";
import {useHistory} from "react-router-dom";

interface P {
    quiz : QuizInterface
}

const MediumCard : React.FC<P> = ({quiz}) => {
    const history = useHistory();
    const quizDetail = () => {
        const path = `/quiz/${quiz.id}`
        history.push(path);
    }

    return (
        <div className="my-card" onClick={quizDetail}>
            <div className="my-medium-card-img">
                <img
                    src="https://picsum.photos/253/180"
                    alt="img big card"
                    style={{height: '100%'}}
                />
            </div>
            <div className="my-medium-card-content">
                <div className="row">
                    <div className="col-sm-12 my-medium-card-content-header">
                        {quiz.name}
                    </div>
                    <div className="col-sm-12 my-medium-card-content-description">
                        {quiz.description}
                    </div>
                    <div className="col-sm-12 my-medium-card-content-author">
                        <div className="d-flex align-items-center">
                            <img
                                src={userIcon}
                                className="my-user-icon"
                                alt="author"
                                style={{marginRight: '10px'}}
                            />
                            {quiz.userRefer?.firstName + ' ' + quiz.userRefer?.lastName}
                        </div>
                        <div className="d-flex align-items-center">
										<span style={{color: '#348009'}}>
											15
										</span>
                            <i className="icon-book-open" style={{color: '#348009'}}></i>
                            <span style={{color: '#f7d66d'}}>
											{quiz.ratings}
										</span>
                            <i className="icon-star" style={{color: '#f7d66d'}}></i>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default MediumCard;