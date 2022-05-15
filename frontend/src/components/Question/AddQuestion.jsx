// Import React dependencies.
import React, { useState } from "react";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";
import Topic from "../Topic/Topic";
import { Dropdown } from "rsuite";
import Select from "react-select";
// import 'rsuite/dist/rsuite.min.css';
import { useEffect } from "react";
import { getTopic } from "../../api/getTopic";
import "./AddQuestion.scss";
import getCurrUid from "../../api/getCurrUid";

const ADDQUES_URL = "http://0.0.0.0:8080/question/add";
const BODY_DEFAULT = "Input your details here...";
const TITLE_DEFAULT = "Input your Question Title here...";
const TOPIC_DEFAULT = "Select Topics";

const AddQuestion = () => {
    // const { userData } = CheckAuth();
    const uid = getCurrUid()

    const [quesBody, setQuesBody] = useState("");
    const [quesTitle, setQuesTitle] = useState("");
    const [success, setSuccess] = useState(false);
    const [selectTopic, setQuesTopic] = useState(TOPIC_DEFAULT);
    const [list, setList] = useState([]);

    useEffect(() => {
        let mounted = true;
        getTopic().then((items) => {
            if (mounted) {
                setList(items);
            }
        });
        return () => (mounted = false);
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (quesBody !== BODY_DEFAULT && quesTitle !== TITLE_DEFAULT) {
            const uid = uid;

            let formData = new FormData();
            formData.append("uid", uid);
            formData.append("body", quesBody);
            formData.append("title", quesTitle);
            if(selectTopic.value>0){
                formData.append("tid", selectTopic.value)
            }

            

            const addAns = {
                method: "POST",
                // headers: { 'content-type': 'multipart/form-data' },
                withCredentials: true,
                body: formData,
            };

            await fetch(ADDQUES_URL, addAns);

            setQuesTitle("");
            setQuesBody("");
            setQuesTopic("");
            setSuccess(true);
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
                    {uid>0 ? (
                        <section className="AddQuestion">
                            <form onSubmit={handleSubmit}>
                                <textarea
                                    className="title"
                                    id="quesTitle"
                                    placeholder={TITLE_DEFAULT}
                                    value={quesTitle}
                                    onChange={(e) => {
                                        setQuesTitle(e.target.value);
                                    }}
                                    required
                                ></textarea>
                                <textarea
                                    className="body"
                                    type="text"
                                    id="quesBody"
                                    placeholder={BODY_DEFAULT}
                                    value={quesBody}
                                    onChange={(e) => {
                                        setQuesBody(e.target.value);
                                    }}
                                    required
                                ></textarea>
                                <Select
                                    value={selectTopic}
                                    onChange={setQuesTopic}
                                    options={list}
                                />
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
