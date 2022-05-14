import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProvider } from "./context/authProvidor";

import "./index.css";
import App from "./App";

ReactDOM.render(
    <BrowserRouter>
        <AuthProvider>
          <Routes>
            <Route path="/*" element={<App />} />
          </Routes>
        </AuthProvider>
    </BrowserRouter>,
    document.getElementById("root")
);
