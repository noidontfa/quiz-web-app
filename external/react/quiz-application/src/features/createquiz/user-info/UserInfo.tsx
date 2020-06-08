import React from "react";
import userIcon from "../../../assets/cat.jpg";
const UserInfo = () => {

    return <>
        <div className="col-xl-2">
            <img
                className="user-interface-avatar"
                src={userIcon}
                alt="user-avartar"
            />
        </div>
        <div className="col-xl-10">
            <div className="row">
                <div className="col-xl-12">
                    <div className="user-interface-header">
									<span className="user-header">
										User Interface
									</span>
                        <i className="icon-settings user-setting"></i>
                    </div>
                </div>
            </div>
            <div className="row d-flex align-items-center">
                <div className="col-xl-2">
                    <span className="label">Firstname:</span>
                </div>
                <div className="col-xl-2">
                    <span className="field">Gia Thinh</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Lastname:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">Ngo Tran</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Fullname:</span>
                </div>
                <div className="col-xl-2">
                    <span className="field">Ngo Tran Gia Thinh</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Birthday:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">19/11/1998</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Username:</span>
                </div>
                <div className="col-xl-2">
                    <span className="field">thinhntg</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Join Date:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">19/11/1998</span>
                </div>
            </div>
        </div>
    </>
}

export default UserInfo;