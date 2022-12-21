import { newPlot, react } from "plotly.js-dist"
import { createEffect } from "solid-js"

export default function Chart2D(props) {
    let d = document.createElement("div")
    let layout = {
        title: props.title,
    }
    newPlot(d, props.data(), layout)
    createEffect(() => {
        react(d, props.data(), layout)
    })
    return <>
        {d}
    </>
}