import { Link } from "react-router-dom";
import "./Login.scss";

function Login() {
  const loginState = {
    username: "",
    password: "",
  };

  function loginHandleChange(name, val){
    this.loginState[name] = val.target.value;
  };

  return (
    <div className="container">
      <div className="loginTitle">
        <span>MyStackOverflow</span>
      </div>
      <div className="loginInput">
        <form>
          <input type="text" placeholder="Username" onChange={val => {val.persist(); loginHandleChange('username',val)}} />
          <input type="text" placeholder="Password" onChange={val => {loginHandleChange('password',val)}} />
        </form>
      </div>
      <div className="btns">
        <span>This is login btns part</span>
      </div>
    </div>
  );
}

export default Login;
