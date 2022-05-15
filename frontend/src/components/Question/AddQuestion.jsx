// Import React dependencies.
import React, { useState } from "react";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";
import Topic from "../Topic/Topic";
import { Dropdown } from "rsuite";
// import 'rsuite/dist/rsuite.min.css'; 
import "./AddQuestion.scss";


const ADDQUES_URL = "http://0.0.0.0:8080/question/add";
const BODY_DEFAULT = "Input your details here...";
const TITLE_DEFAULT = "Input your Question Title here...";
const TOPIC_DEFAULT = "Select Topics";

const AddQuestion = () => {
    const { userData } = CheckAuth();

    const [quesBody, setQuesBody] = useState(BODY_DEFAULT);
    const [quesTitle, setQuesTitle] = useState(TITLE_DEFAULT);
    const [success, setSuccess] = useState(false);
    const [quesTopic, setQuesTopic] = useState(TOPIC_DEFAULT)

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (quesBody !== BODY_DEFAULT && quesTitle !== TITLE_DEFAULT) {
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
        } else {
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
                                <Topic />
                                {/* <Dropdown title={quesTopic} onChange={(e) => {
                                        setQuesTopic(e.target.value);
                                    }}>
                                    <Dropdown.Item value="new file" onClick={(e)=>{
                                        setQuesTopic(e.target.value)
                                    }}>
                                        New File
                                    </Dropdown.Item>
                                    <Dropdown.Item>
                                        New File with Current Profile
                                    </Dropdown.Item>
                                    <Dropdown.Item>
                                        Download As...
                                    </Dropdown.Item>
                                    <Dropdown.Item>Export PDF</Dropdown.Item>
                                    <Dropdown.Item>Export HTML</Dropdown.Item>
                                    <Dropdown.Item>Settings</Dropdown.Item>
                                    <Dropdown.Item>About</Dropdown.Item>
                                </Dropdown> */}
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
                            <Topic />
                        </section>
                    )}
                </section>
            )}
        </>
    );
};

export default AddQuestion;
