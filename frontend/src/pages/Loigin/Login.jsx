import React from "react";
import { Link } from "react-router-dom";
import "./Login.scss";
import Header from "../../components/Header/Header";

// function Login() {
//   const loginState = {
//     username: "",
//     password: "",
//   };

//   function loginHandleChange(name, val){
//     this.loginState[name] = val.target.value;
//   };

//   function handleLogin(){
//     console.log(this.loginState.username);
//   }

//   return (
//     <div className="container">
//       <div className="loginTitle">
//         <span>MyStackOverflow</span>
//       </div>
//       <div className="loginContent">
//         <form>
//           <div className="loginInput">
//             <input type="text" placeholder="Username" onChange={val => {val.persist(); loginHandleChange('username',val)}} />
//             <input type="password" placeholder="Password" onChange={val => {loginHandleChange('password',val)}} />
//           </div>
//           <div className="btns">
//             <input type="submit" value="Login" onClick={handleLogin}/>
//             <Link to={"/signup"} style={{ textDecoration: "none" }}><button>Sign Up</button></Link>
//           </div>
//         </form>
//       </div>
//     </div>
//   );
// }

class Login extends React.Component {
    state = {
        username: "",
        pwd: "",
    };

    handleChange = (e) => {
        const { name, value } = e.target;
        this.setState({ [name]: value });
    };

    handleSubmit = (e) => {
        e.preventDefault();
        this.props.isLogin(true);
    };


    render() {

        return (
            <div>
                <div>
                    <Header isLogin={false} />
                </div>
                <div className="container">
                    <div className="loginTitle">
                        <span>MyStackOverflow</span>
                    </div>
                    <div className="loginContent">
                        <form>
                            <div className="loginInput">
                                <input
                                    type="text"
                                    name="username"
                                    placeholder="Username..."
                                    required
                                    onChange={this.handleChange}
                                />
                                <input
                                    type="password"
                                    name="pwd"
                                    placeholder="Password..."
                                    required
                                    onChange={this.handleChange}
                                />
                            </div>
                            <div className="btns">
                                <button
                                    className="loginbtn"
                                    onSubmit={this.handleSubmit}
                                    style={{ textDecoration: "none" }}
                                >
                                    Log In
                                </button>
                                <Link
                                    to={"/signup"}
                                    style={{ textDecoration: "none" }}
                                >
                                    <button className="signupbtn">
                                        Sign Up
                                    </button>
                                </Link>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;
