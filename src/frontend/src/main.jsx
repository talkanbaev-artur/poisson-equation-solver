import { render } from "solid-js/web";
import { Router, Routes, Route } from "@solidjs/router";
import "./index.css"
import App from "./App";

render(
    () => (
        <Router>
            <Routes>
                <Route path="/" component={App} />
            </Routes>
        </Router>
    ),
    document.getElementById("root")
);