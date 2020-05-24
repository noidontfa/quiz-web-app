import React from "react";

interface P {
    choice : ChoiceInterface;
    index : number;
    callBackFunction? : (id : number) => void;
}

const ChoiceItem : React.FC<P> = ({choice,index,callBackFunction}) => {

    const onClickFunction = () => {
        if (callBackFunction) {
            callBackFunction(choice.id!);
        }
    }

    return   <div className="col-xl-3" style={{marginTop: '30px'}} onClick={onClickFunction}>
                <div className="choice">
                    <div className="choice-title">
                        {index}
                    </div>
                    <div className="choice-content">
                        {choice.name}
                    </div>
                </div>
            </div>
}

export default ChoiceItem;