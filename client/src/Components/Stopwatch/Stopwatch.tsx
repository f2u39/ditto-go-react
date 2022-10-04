import * as React from "react";
import { useEffect, useState } from "react";

export default function Stopwatch() {
    const [time, setTime] = useState(0)

    useEffect(() => {
        const interval = setInterval(() => {
            setTime(prevTime => prevTime + 10)
        }, 10)
        return () => clearInterval(interval)
    }, [])

    return (
        <>
            {("0" + Math.floor((time/3600000)%60)).slice(-2)}h {' '}
            {("0" + Math.floor((time/60000)%60)).slice(-2)}m {' '}
            {("0" + Math.floor((time/1000)%60)).slice(-2)}s {' '}
            {("0" + (time/10)%1000).slice(-2)}ms
        </>
    )
}
