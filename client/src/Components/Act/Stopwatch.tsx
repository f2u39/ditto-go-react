import { useState } from "react";

export default function Stopwatch() {
    const [stopwatch, setStopwatch] = useState(defaultStopwatch)

    const defaultStopwatch = {
        is_counting: false,
        start_time: '',
        type: '',
        game_title: '',
    }
}