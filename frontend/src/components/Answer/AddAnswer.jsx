// Import React dependencies.
import React, { useState } from "react";
import { Link } from "react-router-dom";
import CheckAuth from "../CheckAuth";
import "./AddAnswer.scss";

const ADDANS_URL = "http://0.0.0.0:8080/answer/add";

const AddAnswer = ({ qid }) => {
    const { userData } = CheckAuth();

    const [ansBody, setAnsBody] = useState("Input your answer here...");
    const [success, setSuccess] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();

        const uid = userData.Uid;
        const tqid = qid["0"];

        let formData = new FormData();
        formData.append("uid", uid);
        formData.append("qid", tqid);
        formData.append("body", ansBody);

        const addAns = {
            method: "POST",
            // headers: { 'content-type': 'multipart/form-data' },
            withCredentials: true,
            body: formData,
        };

        await fetch(ADDANS_URL, addAns);

        setAnsBody("");
        setSuccess(true);
        console.log(success);
    };

    return (
        <>
            {success ? (
                <section>
                    <h1>you have successfully answered question</h1>
                    <br />
                    <p>
                        <Link to={"/"}>Go to home page</Link>
                    </p>
                </section>
            ) : (
                <section>
                    {userData ? (
                        <section className="AddAnswer">
                            <form onSubmit={handleSubmit}>
                                <textarea
                                    className="body"
                                    type="text"
                                    id="ansBody"
                                    value={ansBody}
                                    onChange={(e) => {
                                        setAnsBody(e.target.value);
                                    }}
                                    required
                                />
                                <label htmlFor="topic">Topics: </label>

                                <br />
                                <button>Submit Answer</button>
                            </form>
                        </section>
                    ) : (
                        <section className="AddAnswer">
                            <h3>
                                Please <Link to={"/login"}>Log in</Link> to add
                                answer
                            </h3>
                        </section>
                    )}
                </section>
            )}
        </>
    );
};

export default AddAnswer;
