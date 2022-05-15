import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../components/Header/Header";
import { FcOk, FcLike } from "react-icons/fc";
import "./Question.scss";
import { Link } from "react-router-dom";
import Lister from "../../components/Lister/Lister";
import AddAnswer from "../../components/Answer/AddAnswer";
import QuestionLike from "../../components/Question/QuestionLike";
import CheckAuth from "../../api/CheckAuth";

function Question() {
    const {curUser} = CheckAuth()
    const params = useParams();
    const qid = useState(params.qid);

    return (
        <div>
            <QuestionHelper qid={qid} currUid={curUser?.Uid}/>
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
        this.state = {
            question: null,
            answers: null,
            quser: null,
            Likes: null,
        };
    }

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

        this.setState({
            question,
            answers,
            quser,
        });
    }

    render() {
        const { question, answers, quser } = this.state;
        let qpart, apart;
        if (question) {
            const qid = this.props.qid
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
                                    <Link to={`/profile/${quser.data.Uid}`}>
                                        {quser.data.Username}
                                    </Link>
                                </div>
                            )}
                            <p className="time">{datetime}</p>
                            {data.Topics && <p>Topics: {data["Topics"]}</p>}
                            {data["NumOfAnswer"] && (
                                <p>Number of Ansers: {data["NumOfAnswer"]}</p>
                            )}
                        </div>
                        <div className="like">
                            <QuestionLike likes={data.Likes} qid={qid} uid={this.props.currUid}/>
                        </div>
                    </div>
                </div>
            );
        }
        if (answers) {
            const data = answers.data;
            if (Object.keys(data).length !== 0) {
                apart = (
                    <div className="answersPart">
                        <div className="addAnswer">
                            <AddAnswer qid={this.props.qid} />
                        </div>
                        <Lister totalData={answers.data} answer={true} currUid={this.props.currUid}/>
                    </div>
                );
            } else {
                apart = (
                    <div className="addAnswer">
                        <AddAnswer qid={this.props.qid} />
                    </div>
                );
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
