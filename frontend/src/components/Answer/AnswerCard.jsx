import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./AnswerCard.scss";
import { FcCheckmark } from "react-icons/fc";
import AnswerLike from "./AnswerLike";
import getCurrUid from "../../api/getCurrUid";
import SelectBest from "./SelectBest";

class AnswerCard extends React.Component {
    constructor(props) {
        super(props);
    }

    state = {
        auser: null,
        question: null,
    };

    toTime(dateStr) {
        var date = new Date(dateStr).toJSON();
        return new Date(+new Date(date) - 4 * 3600 * 1000)
            .toISOString()
            .replace(/T/g, " ")
            .replace(/\.[\d]{3}Z/, "");
    }

    async componentDidMount() {
        const uid = this.props.data["Uid"];
        const respUser = await fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`);
        const auser = await respUser.json();

        const qid = this.props.data["Qid"];
        const respQuestion = await fetch(`http://0.0.0.0:8080/question/get?qid=${qid}`)
        const question = await respQuestion.json();

        this.setState({ auser, question });
    }

    render() {
        const curId = getCurrUid();
        const data = this.props.data;
        const { auser, question } = this.state;
        const isBest = data.IsBest === 1;
        return (
            <div className="AnswerCard">
                <div className="top-line">
                    <div className="name-time">
                        {auser && (
                            <div>
                                Posted by:{" "}
                                <Link to={`/profile/${auser.data.Uid}`}>
                                    {auser.data.Username}
                                </Link>
                            </div>
                        )}
                        <p className="time">{this.toTime(data.Time)}</p>
                    </div>

                    {isBest && <FcCheckmark className="FcCheckmark" />}
                </div>
                <div className="body">
                    <p>{data.Body}</p>
                </div>

                <div className="btm-line">
                    <div className="btm-left">
                        <p>Likes: {data.Likes}</p>
                        <p>Rating: {data.Rating}</p>
                        <p>Topics: {data.Topics}</p>
                    </div>
                    {curId > 0 && <AnswerLike
                        likes={data.Likes}
                        aid={data.Aid}
                        uid={this.props.currUid}
                    />}
                    {auser && (
                        <div>
                            {curId === question.data.Uid.toString() ?
                                <SelectBest
                                    likes={data.Likes}
                                    aid={data.Aid}
                                    uid={this.props.currUid}
                                />: ''}
                        </div>
                    )}
                </div>
            </div>
        );
    }
}

export default AnswerCard;
