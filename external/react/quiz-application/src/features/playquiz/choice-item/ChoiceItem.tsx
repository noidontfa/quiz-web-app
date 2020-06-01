import React from "react";
import {useTransition,animated} from "react-spring";

interface P {
    choice : ChoiceInterface;
    index : number;
    callBackFunction? : (id : number) => void;
    showChoices : boolean;
}

const ChoiceItem : React.FC<P> = ({choice,index,callBackFunction,showChoices}) => {

    const onClickFunction = () => {
        if (callBackFunction) {
            callBackFunction(choice.id!);
        }
    }
    const transitions = useTransition(showChoices, null ,{
        from: { transform: 'translate3d(0,-40px,0)', opacity: 0},
        enter: { transform: 'translate3d(0,0px,0)', opacity: 1},
        leave: {display: 'none' },
    });

    return   <>
        {
            transitions.map(({ item, key, props}) =>
                <div className="col-xl-3" style={{
                    marginTop: '30px',
                    opacity: item ? (choice.isRight ? "1" : "0") : "1"
                }} onClick={onClickFunction}>
                    <animated.div key={key} style={props} >
                            <div className="choice" data-id={choice.id}>
                                <div className="choice-title">
                                    {index}
                                </div>
                                <div className="choice-content">
                                    {choice.name}
                                </div>
                            </div>
                    </animated.div>
                </div>
            )
        }
    </>
}

export default ChoiceItem;