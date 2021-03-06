import React from "react";
import { Link } from "react-router-dom";

import "./Header.scss";
import "./SearchBar.jsx";
import { FaBeer, FaUser } from "react-icons/fa";
import SearchBar from "./SearchBar.jsx";
import useAuth from "../../hooks/useAuth";
import CheckAuth from "../../api/CheckAuth";
import getCurrUid from "../../api/getCurrUid";
import { getUserData } from "../../api/getUserData";
import { useState, useEffect } from "react";

class Header extends React.Component {
    constructor(props) {
        super(props);
        this.uid = getCurrUid();
    }

    state = {
        userData: null,
    };

    async componentDidMount(){
        if(this.uid>0){
            const response= await fetch(`http://0.0.0.0:8080/user/get?uid=${this.uid}`)
            const result= await response.json()
            this.setState({userData: result.data})
        }
    }

    render() {
        const uid = this.uid;
        const userData = this.state.userData
        console.log(userData)
        return (
            <div className="header">
                <Link
                    to="/"
                    className="title"
                    style={{ textDecoration: "none" }}
                >
                    <h2>MyStackOverflow</h2>
                    <FaBeer className="icon" />
                </Link>
                <div className="search-bar">
                    {this.props.search && (
                        <SearchBar placeholder="Enter Your Keyword here..." />
                    )}
                </div>
                {userData ? (
                    <div className="username">
                        <FaUser className={userData.Status} />
                        <Link
                            to={`/profile/${uid}`}
                            style={{ textDecoration: "none" }}
                        >
                            <h2 className={userData.Status}>
                                {userData.Username}
                            </h2>
                        </Link>
                    </div>
                ) : (
                    <div className="login-btn">
                        <button style={{ textDecoration: "none" }}>
                            <Link
                                to="/login"
                                style={{ textDecoration: "none" }}
                            >
                                Login
                            </Link>
                        </button>
                        <button>
                            <Link
                                to="/signup"
                                style={{ textDecoration: "none" }}
                            >
                                Sign Up
                            </Link>
                        </button>
                    </div>
                )}
            </div>
        );
    }
}

export default Header;
