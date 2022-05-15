import { Link } from "react-router-dom";
import "./QuestionCard.scss";
import { FcOk } from "react-icons/fc";
import React from "react";

class QuestionCard extends React.Component {
    // const datetime = date.parse(data['Time'], 'YYYY-MM-DD hh:mm:ssZZ')
    constructor(props) {
        super(props);
    }

    state = {
        quser: null,
    };

    toTime(dateStr) {
        var date = new Date(dateStr).toJSON();
        return new Date(+new Date(date) + 8 * 3600 * 1000)
            .toISOString()
            .replace(/T/g, " ")
            .replace(/\.[\d]{3}Z/, "");
    }

    async componentDidMount() {
        const uid = this.props.data["Uid"];
        const respUser = await fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`);
        const quser = await respUser.json();

        this.setState({ quser });
    }

    render() {
        const data = this.props.data;
        const datetime = this.toTime(data["Time"]);
        const isResolved = data["IsResolved"] === 1;
        const { quser } = this.state;

        return (
            <div className="questionCard">
                <div className="top-line">
                    <Link to={`/question/${data["Qid"]}`}>
                        <label>{data["Title"]}</label>
                    </Link>
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
                    </div>
                    {data.Topics && <p>Topics: {data["Topics"]}</p>}
                    <p>Number of Ansers: {data["NumOfAnswer"]}</p>
                    <p className="like">Likes: {data["Likes"]}</p>
                </div>
            </div>
        );
    }
}

export default QuestionCard;
