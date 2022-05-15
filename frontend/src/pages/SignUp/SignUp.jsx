import { useRef, useState, useEffect, useContext, React } from "react";
import AuthContext from "../../context/authProvidor";
import md5 from "js-md5";
import { Link } from "react-router-dom";

const SIGNUP_URL = "http://0.0.0.0:8080/user/add";

const SignUp = () => {
    const { setAuth } = useContext(AuthContext);
    const userRef = useRef();
    const errRef = useRef();

    const [username, setUsername] = useState("");
    const [pwd, setPwd] = useState("");
    const [email, setEmail] = useState("");
    const [city, setCity] = useState("");
    const [state, setState] = useState("");
    const [country, setCountry] = useState("");
    const [profile, setProfile] = useState("");
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

        let formData = new FormData();
        formData.append('username', username);
        formData.append('password', md5(pwd));
        formData.append('email', email)
        formData.append('city', city)
        formData.append('state', state)
        formData.append('country', country)
        formData.append('profile', profile)

        const registUser = {
            method: "POST",
            // headers: { 'content-type': 'multipart/form-data' },
            withCredentials: true,
            body: formData,
        };


        const response = await fetch(SIGNUP_URL, registUser);
        const data = await response.json();
        const uid = data.data;

        if (uid > 0) {
            const respUser = await fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`)
            const userResult = await respUser.json();
            const userData = userResult.data;
            
            setAuth({ userData });
            localStorage.setItem("uid", uid)

            setUsername("");
            setPwd("");
            setEmail("");
            setCity("");
            setState("");
            setCountry("");
            setProfile("");
            setSuccess(true);
        } else {
            setErrMsg("Err on creating user, please try another user name");
        }

        // setUsername("");
        // setPwd("");
        // setEmail("");
        // setCity("");
        // setState("");
        // setCountry("");
        // setProfile("");
        // setSuccess(true);
    };

    return (
        <>
            {success ? (
                <section>
                    <h1>you have successfully signed up</h1>
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
                    <h1>Sign Up</h1>
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
                        <br />
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            onChange={(e) => setPwd(e.target.value)}
                            value={pwd}
                            required
                        />
                        <br />
                        <label htmlFor="email">Email:</label>
                        <input
                            type="email"
                            id="email"
                            onChange={(e) => setEmail(e.target.value)}
                            value={email}
                            required
                        />
                        <br />
                        <label htmlFor="city">City:</label>
                        <input
                            type="text"
                            id="city"
                            onChange={(e) => setCity(e.target.value)}
                            value={city}
                            required
                        />
                        <br />
                        <label htmlFor="state">State:</label>
                        <input
                            type="text"
                            id="state"
                            onChange={(e) => setState(e.target.value)}
                            value={state}
                            required
                        />
                        <br />
                        <label htmlFor="country">Country:</label>
                        <input
                            type="text"
                            id="country"
                            onChange={(e) => setCountry(e.target.value)}
                            value={country}
                            required
                        />
                        <br />
                        <label htmlFor="profile">Profile:</label>
                        <input
                            type="text"
                            id="profile"
                            onChange={(e) => setProfile(e.target.value)}
                            value={profile}
                            required
                            size="50"
                        />
                        <button>Sign Up</button>
                    </form>
                    <p>
                        Already Have an Account?
                        <br />
                        <span className="line">
                            <Link to={"/login"}>Sign In</Link>
                        </span>
                    </p>
                </section>
            )}
        </>
    );
};

export default SignUp;
