import React from "react";
import { Link } from "react-router-dom";

import "./Header.scss";
import "./SearchBar.jsx";
import { FaBeer, FaUser } from "react-icons/fa";
import SearchBar from "./SearchBar.jsx";
import useAuth from "../../hooks/useAuth";
import CheckAuth from "../../api/CheckAuth";

const Header = ({ search }) => {
    const { userData } = CheckAuth();

    return (
        <div className="header">
            <Link to="/" className="title" style={{ textDecoration: "none" }}>
                <h2>MyStackOverflow</h2>
                <FaBeer className="icon" />
            </Link>
            <div className="search-bar">
                {search && <SearchBar placeholder="Enter Your Question..." />}
            </div>
            {!userData && (
                <div className="login-btn">
                    <button style={{ textDecoration: "none" }}>
                        <Link to="/login" style={{ textDecoration: "none" }}>
                            Login
                        </Link>
                    </button>
                    <button>
                        <Link to="/signup" style={{ textDecoration: "none" }}>
                            Sign Up
                        </Link>
                    </button>
                </div>
            )}
            {userData && (
                <div className = "username">
                    <FaUser className={userData.Status}/>
                    <Link
                        to={`/profile/${userData.Uid}`}
                        style={{ textDecoration: "none" }}
                        
                    >
                        
                        <h2 className={userData.Status}>{userData.Username}</h2>
                    </Link>
                </div>
            )}
        </div>
    );
};

export default Header;
