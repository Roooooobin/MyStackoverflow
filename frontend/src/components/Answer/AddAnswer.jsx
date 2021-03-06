// Import React dependencies.
import React, { useState } from "react";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";
import "./AddAnswer.scss";
import Select from "react-select";
import { useEffect } from "react";
import { getTopic } from "../../api/getTopic";
import getCurrUid from "../../api/getCurrUid";

const ADDANS_URL = "http://0.0.0.0:8080/answer/add";
const TOPIC_DEFAULT = "Select Topics";
const BODY_DEFAULT = "Input your answer here...";

const AddAnswer = ({ qid }) => {
    const uid = getCurrUid()


    const [ansBody, setAnsBody] = useState("");
    const [success, setSuccess] = useState(false);
    const [selectTopic, setAnsTopic] = useState(TOPIC_DEFAULT);
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

        const tqid = qid["0"];

        let formData = new FormData();
        formData.append("uid", uid);
        formData.append("qid", tqid);
        formData.append("body", ansBody);
        if(selectTopic.value>0){
            formData.append("tid", selectTopic.value)
        }

        const addAns = {
            method: "POST",
            // headers: { 'content-type': 'multipart/form-data' },
            withCredentials: true,
            body: formData,
        };

        await fetch(ADDANS_URL, addAns);

        setAnsBody("");
        setAnsTopic("");
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
                    {uid > 0 ? (
                        <section className="AddAnswer">
                            <form onSubmit={handleSubmit}>
                                <textarea
                                    className="body"
                                    type="text"
                                    id="ansBody"
                                    value={ansBody}
                                    placeholder={BODY_DEFAULT}
                                    onChange={(e) => {
                                        setAnsBody(e.target.value);
                                    }}
                                    required
                                ></textarea>
                                <Select
                                    value={selectTopic}
                                    onChange={setAnsTopic}
                                    options={list}
                                />
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
