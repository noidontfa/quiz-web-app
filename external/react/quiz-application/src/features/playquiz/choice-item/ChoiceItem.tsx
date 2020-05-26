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
        from: { transform: 'translate3d(0,-40px,0)' },
        enter: { transform: 'translate3d(0,0px,0)' },
        leave: { transform: 'translate3d(0,-40px,0)' , display: 'none'},
    });

    return   <>
        {
            transitions.map(({ item, key, props}) =>
                <div className="col-xl-3" style={{
                    marginTop: '30px',
                    display: item ? (choice.isRight ? "block" : "none") : "block"
                }} onClick={onClickFunction}>
                    <animated.div key={key} style={props} >
                            <div className="choice">
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