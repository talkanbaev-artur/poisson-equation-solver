import { createEffect, createMemo, createSignal, onMount } from "solid-js"
import { cleanParams, solve } from "./api/api";
import Params from "./components/parameters";
import Chart2D from "./components/2dchart";
import { createAxis, getHalfYValues, getHalfXValues } from "./api/charts";
import Chart3D from "./components/3dchart";

export default function App() {


    const [params, setParams] = createSignal({ n: 10, tau: 1, theta: 0, task: 2 }, { equals: false })
    const [error, serError] = createSignal(0);
    const [solution, setSolution] = createSignal(null);
    const f = async () => {
        let data = await solve(cleanParams(params()));
        setSolution(data.data);
        setK(0);
        let l = setInterval(() => {
            (k() + 1 < data.data.k - 1) ? setK(k() + 1) : null
            if (k() + 1 >= data.data.k - 1) {
                clearInterval(l);
                console.log("stop");
            }
        }, 100);
    }
    onMount(f)
    createEffect(f)

    const [k, setK] = createSignal(1);

    const axis = createMemo(() => createAxis(parseInt(params().n)))

    const a = createMemo(() => solution() ? [getHalfYValues(solution().data[k()], axis(), "num"), getHalfYValues(solution().original, axis(), "org")] : [])
    const b = createMemo(() => solution() ? [getHalfXValues(solution().data[k()], axis(), "num"), getHalfXValues(solution().original, axis(), "org")] : [])

    return <div className="flex">
        <div className="flex mx-8 my-6">
            <div className="flex mr-auto flex-col p-4">
                <h2 className="text-3xl font-bold mb-2">Poisson equation solver</h2>
                <p>Made by Talkanbaev Artur</p>
                <div className="mt-8">
                    <h3 className="font-bold text-lg mb-2">Manual</h3>
                    <p>1. Edit parameters - use the arrows or write the number down</p>
                    <p>2. Use mouse and wheel to navigate the graph</p>
                    <p>3. Hower over graphs to see values</p>
                </div>
                <Params value={params} setter={setParams} />
                <div className="flex flex-col my-6 space-y-4">
                    <h4 className="font-bold text-lg">Error values</h4>
                    <p>Error: {solution() ? solution().err : 0}</p>
                    <p>Sigma: {solution() ? solution().sigma : 0}</p>
                    <p>Iterations: {solution() ? solution().k : 0}</p>
                </div>
            </div>
        </div>
        <div className="flex flex-col flex-1 p-10 border">
            <p>Showing iteration: {k() + 2}</p>
            <div className="flex">
                <div>
                    <Chart2D title={"Y=0.5 chart"} data={a} />
                    {JSON.stringify(a)}
                </div>
                <div>
                    <Chart2D title={"X=0.5 chart"} data={b} />
                </div>
            </div>
            <Chart3D data={solution() ? [{ z: solution().data[solution().k - 1], x: axis(), y: axis(), type: "surface" }] : []} axis={axis()} />
        </div>
    </div>
}