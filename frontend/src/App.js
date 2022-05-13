import * as React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Loigin/Login";
import Question from "./pages/Question";
import SignUp from "./pages/SignUp";
import Result from "./pages/SearchResult";

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
                    <Route path="/result" element={<Result />} />
                    <Route path="/result/:q" exact element={<Result />} />
                    <Route path="/question/:qid" exact element={<Question />} />
                </Routes>
            </div>
        );
    }
}

export default App;
