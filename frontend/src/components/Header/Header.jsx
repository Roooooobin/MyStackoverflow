import React from "react";
import { Link } from "react-router-dom";

import "./Header.scss";
import "./SearchBar.jsx";
import { FaBeer } from "react-icons/fa";
import SearchBar from "./SearchBar.jsx";

class Header extends React.Component {

    render() {
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
                    <SearchBar placeholder="Enter Your Question..." />
                </div>
                {!this.props.isLogin && (
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
                {this.props.isLogin && 
                  <div className="username">
                    <h3>username</h3>
                  </div>
                }
            </div>
        );
    }
}

export default Header;
