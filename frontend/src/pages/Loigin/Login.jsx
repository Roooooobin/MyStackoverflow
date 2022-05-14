import React from "react";
import { Link } from "react-router-dom";
import "./Login.scss";
import Header from "../../components/Header/Header";
import { useRef, useState, useEffect, useContext } from "react";
import AuthContext from "../../context/authProvidor";
import md5 from "js-md5";

const LOGIN_URL = "http://0.0.0.0:8080/user/authorize";

const Login = () => {
    const { setAuth } = useContext(AuthContext);
    const userRef = useRef();
    const errRef = useRef();

    const [username, setUsername] = useState("");
    const [pwd, setPwd] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const [success, setSuccess] = useState(false);

    useEffect(() => {
        userRef.current.focus();
    }, []);

    useEffect(() => {
        setErrMsg("");
    }, [username, pwd]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const authUser = {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                withCredentials: true,
                body: JSON.stringify({
                    username: username,
                    password: md5(pwd),
                }),
            };

            const response = await fetch(LOGIN_URL, authUser);

            console.log(JSON.stringify(response?.data));

            const uid = response.data.uid;
            let login;
            if (uid > 0) {
                login = true;
            } else {
                login = false;
            }

            setAuth({ uid, username, login });

            setUsername("");
            setPwd("");
            setSuccess(true);
        } catch (err) {
            if (!err?.response) {
                setErrMsg("No Server Response");
            } else if (err.response?.status === 400) {
                setErrMsg("Missing username or password");
            } else if (err.response?.status === 401) {
                setErrMsg("Unauthorized");
            } else if (err.response?.status === 404) {
                setErrMsg("No Auth API");
            } else {
                setErrMsg("Login failed");
            }
            errRef.current.focus();
        }
    };

    return (
        <>
            {success ? (
                <section>
                    <h1>you have logged in</h1>
                    <br />
                    <p>
                        <Link to={"/"}>Go to home page</Link>
                    </p>
                </section>
            ) : (
                <section>
                    <p
                        ref={errRef}
                        className={errMsg ? "errmsg" : "offscreen"}
                        aria-live="assertive"
                    >
                        {errMsg}
                    </p>
                    <h1>Sign In</h1>
                    <form onSubmit={handleSubmit}>
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            ref={userRef}
                            autoComplete="off"
                            onChange={(e) => setUsername(e.target.value)}
                            value={username}
                            required
                        />
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            onChange={(e) => setPwd(e.target.value)}
                            value={pwd}
                            required
                        />
                        <button>Sign In</button>
                    </form>
                    <p>
                        Need an Account?
                        <br />
                        <span className="line">
                            <Link to={"/signup"}>Sign Up</Link>
                        </span>
                    </p>
                </section>
            )}
        </>
    );
};

// class Login extends React.Component {
//     state = {
//         username: "",
//         pwd: "",
//     };

//     handleChange = (e) => {
//         const { name, value } = e.target;
//         this.setState({ [name]: value });
//     };

//     handleSubmit = (e) => {
//         e.preventDefault();
//         this.props.isLogin(true);
//     };

//     render() {
//         return (
//             <div>
//                 <div>
//                     <Header isLogin={false} search={false} />
//                 </div>
//                 <div className="container">
//                     <div className="loginTitle">
//                         <span>MyStackOverflow</span>
//                     </div>
//                     <div className="loginContent">
//                         <form>
//                             <div className="loginInput">
//                                 <input
//                                     type="text"
//                                     name="username"
//                                     placeholder="Username..."
//                                     required
//                                     onChange={this.handleChange}
//                                 />
//                                 <input
//                                     type="password"
//                                     name="pwd"
//                                     placeholder="Password..."
//                                     required
//                                     onChange={this.handleChange}
//                                 />
//                             </div>
//                             <div className="btns">
//                                 <button
//                                     className="loginbtn"
//                                     onSubmit={this.handleSubmit}
//                                     style={{ textDecoration: "none" }}
//                                 >
//                                     Log In
//                                 </button>
//                                 <Link
//                                     to={"/signup"}
//                                     style={{ textDecoration: "none" }}
//                                 >
//                                     <button className="signupbtn">
//                                         Sign Up
//                                     </button>
//                                 </Link>
//                             </div>
//                         </form>
//                     </div>
//                 </div>
//             </div>
//         );
//     }
// }

export default Login;
