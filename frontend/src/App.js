import * as React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Loigin/Login";
import Question from "./pages/Question";
import SignUp from "./pages/SignUp";
import SearchResult from "./pages/SearchResult";

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" exact element={<Home />} />
        <Route path="/login" exact element={<Login />} />
        <Route path="/signup" exact element={<SignUp />} />
        <Route path="/result" exact element={<SearchResult />} />
        <Route path="/result/:q" exact element={<SearchResult />} />
        <Route path="/question/:qid" exact element={<Question />} />
      </Routes>
    </div>
  );
}

export default App;
