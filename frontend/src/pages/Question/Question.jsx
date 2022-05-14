import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Question.scss";
import { Link } from "react-router-dom";
import Lister from "../../components/Lister/Lister";

function Question() {
    const params = useParams();
    const qid = useState(params.qid);

    return (
        <div>
            <QuestionHelper qid={qid} />
        </div>
    );
}

var toTime = function (dateStr) {
    var date = new Date(dateStr).toJSON();
    return new Date(+new Date(date) + 8 * 3600 * 1000)
        .toISOString()
        .replace(/T/g, " ")
        .replace(/\.[\d]{3}Z/, "");
};

class QuestionHelper extends React.Component {
    constructor(props) {
        super(props);
    }

    state = {
        question: null,
        answers: null,
        quser: null,
    };

    async componentDidMount() {
        const qid = this.props.qid;
        // GET request using fetch with async/await
        const respQuestion = await fetch(
            `http://0.0.0.0:8080/question/get?qid=${qid}`
        );
        const question = await respQuestion.json();

        const respAnswer = await fetch(
            `http://0.0.0.0:8080/answer/list?qid=${qid}`
        );
        const answers = await respAnswer.json();

        const respUser = await fetch(
            `http://0.0.0.0:8080/user/get?uid=${question["data"]["Uid"]}`
        );
        const quser = await respUser.json();

        this.setState({ question, answers, quser });
    }

    render() {
        const { question, answers, quser } = this.state;
        let qpart, apart;
        if (question) {
            const data = question.data;
            const isResolved = data["IsResolved"] === 1;
            const datetime = toTime(data["Time"]);
            qpart = (
                <div className="questionPart">
                    <div className="top-line">
                        <label>{data["Title"]}</label>
                        {isResolved && <FcOk className="isResolved" />}
                    </div>

                    <div className="body">
                        <p>{data["Body"]}</p>
                    </div>

                    <div className="btm-line">
                        <div className="name-time">
                            {quser && (
                                <div>
                                    Posted by:{" "}
                                    <Link to={`/profile/${quser.data.Uid}`}>{quser.data.Username}</Link>
                                </div>
                            )}
                            <p className="time">{datetime}</p>
                        </div>
                        <p className="like">Likes: {data["Likes"]}</p>
                    </div>
                </div>
            );
        }
        if (answers) {
            const data = answers.data;
            if (Object.keys(data).length !== 0) {
                apart = (
                    <div className="answersPart">
                        <Lister totalData={answers.data} answer={true} />
                    </div>
                );
            } else {
                apart = <div>No answer</div>;
            }
        }

        return (
            <div className="main">
                <Header search={true} />
                <div>
                    <h2>Question: </h2>
                    {qpart}
                </div>
                <div>
                    <h2>Answer: </h2>
                    {apart}
                </div>
            </div>
        );
    }
}

export default Question;
