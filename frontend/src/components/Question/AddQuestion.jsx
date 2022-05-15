// Import React dependencies.
import React, { useState } from "react";
import { Link } from "react-router-dom";
import CheckAuth from "../CheckAuth";
import "./AddQuestion.scss";

const ADDQUES_URL = "http://0.0.0.0:8080/question/add";
const BODY_DEFAULT = "Input your details here...";
const TITLE_DEFAULT = "Input your Question Title here...";

const AddQuestion = () => {
    const { userData } = CheckAuth();

    const [quesBody, setQuesBody] = useState(BODY_DEFAULT);
    const [quesTitle, setQuesTitle] = useState(TITLE_DEFAULT);
    const [success, setSuccess] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if(quesBody !== BODY_DEFAULT && quesTitle !== TITLE_DEFAULT){
            const uid = userData.Uid;

            let formData = new FormData();
            formData.append("uid", uid);
            formData.append("body", quesBody);
            formData.append("title", quesTitle);
    
            const addAns = {
                method: "POST",
                // headers: { 'content-type': 'multipart/form-data' },
                withCredentials: true,
                body: formData,
            };
    
            await fetch(ADDQUES_URL, addAns);
    
            setQuesTitle("");
            setQuesBody("");
            setSuccess(true);
            console.log(success);
        }
        else{
            alert("can not post default value");
        }
    };

    return (
        <>
            {success ? (
                <section>
                    <h1>you have successfully posted question</h1>
                    <br />
                    <p>
                        <Link to={"/"}>Go to home page</Link>
                    </p>
                </section>
            ) : (
                <section>
                    {userData ? (
                        <section className="AddQuestion">
                            <form onSubmit={handleSubmit}>
                                <textarea
                                    className="title"
                                    type="text"
                                    id="quesTitle"
                                    value={quesTitle}
                                    defaultValue={TITLE_DEFAULT}
                                    onChange={(e) => {
                                        setQuesTitle(e.target.value);
                                    }}
                                    required
                                />
                                <textarea
                                    className="body"
                                    type="text"
                                    id="quesBody"
                                    value={quesBody}
                                    onChange={(e) => {
                                        setQuesBody(e.target.value);
                                    }}
                                    defaultValue={BODY_DEFAULT}
                                    required
                                />
                                <label htmlFor="topic">Topics: </label>

                                <br />
                                <button>Post Question</button>
                            </form>
                        </section>
                    ) : (
                        <section className="AddQuestion">
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

export default AddQuestion;
