import React, { useState, useEffect } from "react";
import "./Editable.css";

const Editable = ({
                      text,
                      type,
                      placeholder,
                      children,
                      childRef,
                      callBackFunction,
                      callBackDeleteFunction,
                      ...props
                  }) => {

    const [isEditing, setEditing] = useState(false);

    useEffect(() => {
        if (childRef && childRef.current && isEditing === true) {
            childRef.current.focus();
        }
    }, [isEditing, childRef]);

    const handleKeyDown = (event, type) => {
        const { key } = event;
        const keys = ["Escape", "Tab"];
        const allKeys = [...keys];
        if (
            (type === "textarea" && keys.indexOf(key) > -1) ||
            (type !== "textarea" && allKeys.indexOf(key) > -1) ||
            (type === "textarea" && event.ctrlKey && key == "Enter")
        ) {
            callBackFunction();
            setEditing(false);
        }
    };

    return (
        <span style={{width: "100%",cursor: "pointer"}}>
            {isEditing ? (
                <div
                    onBlur={() =>{ callBackFunction(); setEditing(false)}}
                    onKeyDown={e => handleKeyDown(e, type)}
                    style={{flexGrow: "1", width: "100%"}}
                >
                    {children}

                </div>
            ) : (
                <div
                    className={`text-table-content `}

                >
          <span className={`${text ? "text-black" : "text-gray-500"}`}   onClick={() => setEditing(true)}>
            {text || placeholder || "Editable content"}
          </span>
                    <div className="table-action" onClick={() => {callBackDeleteFunction(); setEditing(false)}}>
                        <i className="icon-trash"></i>
                    </div>
                </div>
            )}
        </span>
    );
};

export default Editable;