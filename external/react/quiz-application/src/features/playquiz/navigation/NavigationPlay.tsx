import React, {useEffect, useLayoutEffect, useState} from "react";
import {setInterval} from "timers";
import {animated, useTransition} from "react-spring";
interface P {
    question: number;
    totalQuestion: number;
    score: number;
    sec: number;
    showChoices : boolean;
    callbackFunction?: () => void;
}
const NavigationPlay : React.FC<P>= ({question,totalQuestion,score,sec,callbackFunction,showChoices}) => {

    const [togglePause,setTogglePause] = useState(true);
    const [widthStyle, setWidthStyle] = useState('100%');
    const [timing,setTiming] = useState(() => {
        return sec;
    });
    const TimingTransition = `width ${timing}s ease-out`

    useEffect(()=> {
        const btnPause = document.getElementById("btn_pause");
        const timingBar = document.getElementById("timing-bar");
        btnPause!.addEventListener('click',function (e) {
            if (timingBar!.parentElement!.classList.contains("dec")) {
                const computedStyle = window.getComputedStyle(timingBar!);
                const widthStyle = computedStyle.getPropertyValue("width")
                setWidthStyle(widthStyle);
            }
            setTogglePause(togglePause => !togglePause);
        });

        timingBar!.addEventListener('transitioncancel', function (e) {
            setTiming(timing => timing - e.elapsedTime);
        })

        timingBar!.addEventListener('transitionend', function (e) {
            callbackFunction!();
            timingBar!.style.removeProperty("width");
            setTiming(sec => sec + timing);
            setTogglePause( togglePause => !togglePause);
        })

    },[])
    useEffect(() => {
        setTimeout(() => {
            setTogglePause(togglePause => !togglePause)
        },300);
    },[score])

    const transitions = useTransition(score, null, {
        from: { opacity: showChoices ? 1 : 0},
        enter: { opacity: 1 },
        leave: { display: "none" },
    })


    return <>
        <header className="play-quiz-navbar ">
            <div className="wrapper">
                <div className="wrap-item">
                    <div
                        className={togglePause ? "action-button pause" : "action-button"}
                        id="btn_pause"
                        style={{marginRight: '15px'}}
                    >
                        <i className="icon-control-pause"></i>
                        <i className="icon-control-play"></i>
                    </div>
                    <div className="action-button">
						<span>
                            {question + "/" + totalQuestion}
						</span>
                    </div>
                </div>

                            {
                                transitions.map(({item,key,props}) =>
                                   score &&  <div className="wrap-item">
                                        <div className="action-button">
						                    <span>
                                                <animated.div key={key} style={props} >
                                                    {score}
                                                </animated.div>
                                            </span>
                                        </div>
                                    </div>
                                )
                            }

            </div>
        </header>

        <div  style={{ transition: TimingTransition}} className={togglePause ? "no-transition" :  "dec"} >
            <div  id="timing-bar" className="timing-bar" style={{width: !showChoices ? widthStyle : "0" }}></div>
        </div>
    </>



}

export default NavigationPlay;