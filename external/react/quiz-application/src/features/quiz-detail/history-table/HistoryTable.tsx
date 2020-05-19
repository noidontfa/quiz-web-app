import React from "react";
import userIcon from "../../../assets/cat.jpg";

const HistoryTable = () => {
    return (
        <div className="col-xl-12" style={{marginTop: '75px'}}>
            <table className="table my-table">
                <thead>
                <tr>
                    <th scope="col" row-data="index">#</th>
                    <th scope="col" row-data="history-user-name">Name</th>
                    <th scope="col" row-data="history-score">Actions</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <th scope="row">1</th>
                    <td>
                        <div className="d-flex align-item-end">
                            <img
                                className="quiz-detail-user-icon"
                                src={userIcon}
                                alt="user-avartar"
                            />
                            <span className="quiz-detail-user-name"
                            >Ngo Tran Gia Thinh</span
                            >
                        </div>
                    </td>
                    <td>
                        <span className="quiz-detail-score">500,123</span>
                    </td>
                </tr>
                <tr>
                    <th scope="row">2</th>
                    <td>
                        <div className="d-flex align-item-end">
                            <img
                                className="quiz-detail-user-icon"
                                src={userIcon}
                                alt="user-avartar"
                            />
                            <span className="quiz-detail-user-name"
                            >Ngo Tran Gia Thinh</span
                            >
                        </div>
                    </td>
                    <td>
                        <span className="quiz-detail-score">500,123</span>
                    </td>
                </tr>
                <tr>
                    <th scope="row">3</th>
                    <td>
                        <div className="d-flex align-item-end">
                            <img
                                className="quiz-detail-user-icon"
                                src={userIcon}
                                alt="user-avartar"
                            />
                            <span className="quiz-detail-user-name"
                            >Ngo Tran Gia Thinh</span
                            >
                        </div>
                    </td>
                    <td>
                        <span className="quiz-detail-score">500,123</span>
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
    )
}

export default HistoryTable;