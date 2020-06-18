import React, {useEffect, useState} from "react";
import userIcon from "../../../assets/cat.jpg";
import axios from "axios";
interface P {
    quiz : QuizInterface;

}

const HistoryTable : React.FC<P> = ({quiz}) => {

    const [histories,setHistories] = useState<Array<HistoryInterface>>([])

    useEffect(() => {
           const datenow = new Date();
           const formatedDate = `${datenow.getFullYear()}-${(datenow.getMonth() + 1).toString().padStart(2,'0')}-${datenow.getDate().toString().padStart(2,'0')}`;
           axios.get(`http:/api/histories/d?date=${formatedDate}&quizid=${quiz.id}`)
               .then(function (res) {
                   if(res.data)
                       setHistories(res.data);
               })
    },[quiz])

    return (
        <div className="col-xl-12" style={{marginTop: '75px'}}>
            <table className="table my-table">
                <thead>
                <tr>
                    <th scope="col" row-data="index">#</th>
                    <th scope="col" row-data="history-user-name">Name</th>
                    <th scope="col" row-data="history-score">Scores</th>
                </tr>
                </thead>
                <tbody>
                {histories.map((h,i) =>
                    <tr>
                        <th scope="row">{i + 1}</th>
                        <td>
                            <div className="d-flex align-item-end">
                                <img
                                    className="quiz-detail-user-icon"
                                    src={userIcon}
                                    alt="user-avartar"
                                />
                                <span className="quiz-detail-user-name"
                                >{h.userRefer?.firstName + " " + h.userRefer?.lastName}</span
                                >
                            </div>
                        </td>
                        <td>
                            <span className="quiz-detail-score">{h.score}</span>
                        </td>
                    </tr>
                )}
                {/*<tr>*/}
                {/*    <th scope="row">1</th>*/}
                {/*    <td>*/}
                {/*        <div className="d-flex align-item-end">*/}
                {/*            <img*/}
                {/*                className="quiz-detail-user-icon"*/}
                {/*                src={userIcon}*/}
                {/*                alt="user-avartar"*/}
                {/*            />*/}
                {/*            <span className="quiz-detail-user-name"*/}
                {/*            >Ngo Tran Gia Thinh</span*/}
                {/*            >*/}
                {/*        </div>*/}
                {/*    </td>*/}
                {/*    <td>*/}
                {/*        <span className="quiz-detail-score">500,123</span>*/}
                {/*    </td>*/}
                {/*</tr>*/}
                {/*<tr>*/}
                {/*    <th scope="row">2</th>*/}
                {/*    <td>*/}
                {/*        <div className="d-flex align-item-end">*/}
                {/*            <img*/}
                {/*                className="quiz-detail-user-icon"*/}
                {/*                src={userIcon}*/}
                {/*                alt="user-avartar"*/}
                {/*            />*/}
                {/*            <span className="quiz-detail-user-name"*/}
                {/*            >Ngo Tran Gia Thinh</span*/}
                {/*            >*/}
                {/*        </div>*/}
                {/*    </td>*/}
                {/*    <td>*/}
                {/*        <span className="quiz-detail-score">500,123</span>*/}
                {/*    </td>*/}
                {/*</tr>*/}
                {/*<tr>*/}
                {/*    <th scope="row">3</th>*/}
                {/*    <td>*/}
                {/*        <div className="d-flex align-item-end">*/}
                {/*            <img*/}
                {/*                className="quiz-detail-user-icon"*/}
                {/*                src={userIcon}*/}
                {/*                alt="user-avartar"*/}
                {/*            />*/}
                {/*            <span className="quiz-detail-user-name"*/}
                {/*            >Ngo Tran Gia Thinh</span*/}
                {/*            >*/}
                {/*        </div>*/}
                {/*    </td>*/}
                {/*    <td>*/}
                {/*        <span className="quiz-detail-score">500,123</span>*/}
                {/*    </td>*/}
                {/*</tr>*/}
                </tbody>
            </table>
        </div>
    )
}

export default HistoryTable;