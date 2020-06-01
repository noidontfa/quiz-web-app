import React, {useEffect, useState} from "react";
import {Rating} from '@material-ui/lab'

interface P {
    defaultValue : number;
    readonly? : boolean;
    callbackFunction? : (value : number) => void;
}

const Ratings : React.FC<P> = ({defaultValue,readonly,callbackFunction}) => {
    const [value,setValue] = useState(defaultValue || 0)

    return (
        <Rating
            precision={0.5}
            readOnly={readonly ? true : false}
            onChange={(event, newValue) => {
                setValue(newValue!);
                callbackFunction!(newValue!);
            }}
            defaultValue={value}
            // onChangeActive={(event, newHover) => {
            //     setHover(newHover);
            // }}
        />
    )
}

export default Ratings;