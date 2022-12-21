import { newPlot, react } from "plotly.js-dist"
import { createEffect } from "solid-js"

export default function Chart3D(props) {
    let d = document.createElement("div")
    d.classList = ["mx-auto"]
    let layout = {
        title: props.title,
        height: 1000,
        width: 1000
    }
    newPlot(d, props.data, layout)
    createEffect(() => {
        react(d, props.data, layout)
    })
    return <>
        {d}
    </>
}