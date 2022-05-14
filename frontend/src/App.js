import * as React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home/Home";
import Login from "./pages/Loigin/Login";
import Question from "./pages/Question/Question";
import SignUp from "./pages/SignUp/SignUp";
import Result from "./pages/SearchResult";
import Profile from "./pages/Profile/Profile";
import ProfileEdit from "./pages/Profile/Edit/Edit";

class App extends React.Component {

    render() {
        return (
            <div>
                <Routes>
                    <Route path="/" exact element={<Home />} />
                    <Route
                        path="/login"
                        exact
                        element={<Login/>}
                    />
                    <Route path="/signup" exact element={<SignUp />} />
                    <Route path="/result" exact element={<Result />} />
                    <Route path="/result/:q" exact element={<Result />} />
                    <Route path="/question/:qid" exact element={<Question />} />
                    <Route path="/profile/:uid" exact element={<Profile />} />
                    <Route path="/profile/edit/:uid" exact element={<ProfileEdit />} />
                </Routes>
            </div>
        );
    }
}

export default App;
