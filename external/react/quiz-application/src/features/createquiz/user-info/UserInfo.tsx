import React from "react";
import userIcon from "../../../assets/cat.jpg";

interface P {
    user? : UserInterface;
}

const UserInfo : React.FC<P> = ({user}) => {

    return <>
        <div className="col-xl-2">
            <img
                className="user-interface-avatar"
                src={user?.image}
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
                    <span className="field">{user?.firstName}</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Lastname:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">{user?.lastName}</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Fullname:</span>
                </div>
                <div className="col-xl-2">
                    <span className="field">{user?.lastName + " " + user?.lastName}</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Birthday:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">{user?.dayOfBirth}</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Username:</span>
                </div>
                <div className="col-xl-2">
                    <span className="field">{user?.username}</span>
                </div>
                <div className="col-xl-2">
                    <span className="label">Join Date:</span>
                </div>
                <div className="col-xl-6">
                    <span className="field">{user?.dayOfBirth}</span>
                </div>
            </div>
        </div>
    </>
}

export default UserInfo;