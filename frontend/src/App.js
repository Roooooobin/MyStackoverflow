import * as React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home/Home";
import Login from "./pages/Loigin/Login";
import Question from "./pages/Question/Question";
import SignUp from "./pages/SignUp";
import Result from "./pages/SearchResult";
import Profile from "./pages/Profile";

class App extends React.Component {
    state = {
        isLogin: false,
    };

    handleLogin = (isLogin) => this.setState({isLogin})

    render() {
        return (
            <div>
                <Routes>
                    <Route path="/" exact element={<Home />} />
                    <Route
                        path="/login"
                        exact
                        element={<Login isLogin={this.handleLogin} />}
                    />
                    <Route path="/signup" exact element={<SignUp />} />
                    <Route path="/result" exact element={<Result />} />
                    <Route path="/result/:q" exact element={<Result />} />
                    <Route path="/question/:qid" exact element={<Question />} />
                    <Route path="/profile/:uid" exact element={<Profile />} />
                </Routes>
            </div>
        );
    }
}

export default App;
