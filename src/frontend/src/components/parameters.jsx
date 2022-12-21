import { getSchemas, getTasks } from "../api/api"
import { createEffect, createSignal, onMount } from "solid-js";
import { Show } from "solid-js";

export default function Params(props) {
    const [schemas, setSchemas] = createSignal([]);
    const [tasks, setTasks] = createSignal([]);
    const [schema, setSchema] = createSignal(1);

    createEffect(() => {
        const s = schemas().find(el => el.id == schema())
        if (s && !(s.tau == 0 && s.theta == 0)) {
            props.setter(cur => {
                cur.theta = s.theta;
                cur.tau = s.tau;
                return cur;
            })
        }
    })

    onMount(async () => {
        const tasksData = await getTasks();
        const schemaData = await getSchemas();
        setTasks(tasksData.data)
        setSchemas(schemaData.data)
    })

    return <>
        <h4 className="mt-4 text-lg font-bold">Parameters</h4>
        <div className="flex flex-col">
            <div className="flex flex-col my-2">
                <label className="text-gray-800-text-small mb-2">Grid size</label>
                <input
                    className="outline-none ring hover:shadow-xl mr-auto ring-green-400 rounded-lg px-4"
                    value={props.value().n}
                    onInput={(e) => props.setter(cur => {
                        cur.n = e.target.value;
                        return cur;
                    })}
                />
            </div>
            <div className="flex flex-col my-4">
                <label className="text-gray-800-text-small mb-2">Schema</label>
                <select className="outline-none ring ring-green-400 mr-auto pr-10 rounded-lg hover:shadow-xl bg-white py-1 pl-2" value={schema()}
                    onChange={(e) => { setSchema(e.target.value) }}>
                    {schemas() ? schemas().map(s => {
                        return (
                            <option key={s.id} value={s.id}>{s.name}</option>
                        )
                    }) : ""}
                </select>
                <Show when={schema() == 3}>
                    <div className="mt-4 flex flex-col space-y-2">
                        <label className="text-gray-800-text-small mb-2">Theta</label>
                        <input
                            className="outline-none ring hover:shadow-xl mr-auto ring-green-400 rounded-lg px-4"
                            value={props.value().theta}
                            onInput={(e) => props.setter(cur => {
                                cur.theta = e.target.value;
                                return cur;
                            })}
                        />
                    </div>
                    <div className="mt-4 flex flex-col space-y-2">
                        <label className="text-gray-800-text-small mb-2">Tau</label>
                        <input
                            className="outline-none ring hover:shadow-xl mr-auto ring-green-400 rounded-lg px-4"
                            value={props.value().tau}
                            onInput={(e) => props.setter(cur => {
                                cur.tau = e.target.value;
                                return cur;
                            })}
                        />
                    </div>
                </Show>
            </div>
            <div className="flex flex-col my-4">
                <label className="text-gray-800-text-small mb-2">Task</label>
                <select className="outline-none ring ring-green-400 mr-auto pr-10 rounded-lg hover:shadow-xl bg-white py-1 pl-2" value={props.value().task}
                    onChange={(e) => {
                        props.setter((cur) => {
                            cur.task = parseInt(e.target.value);
                            return cur;
                        })
                    }}>
                    {tasks() ? tasks().map(s => {
                        return (
                            <option key={s.id} value={s.id}>{s.name}</option>
                        )
                    }) : ""}
                </select>
            </div>
        </div>
    </>
}