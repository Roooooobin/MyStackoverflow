import { Link } from "react-router-dom";
import "./QuestionCard.scss";
import date from "date-and-time";
import {FcOk} from "react-icons/fc"

function QuestionCard({ data }) {
    // const datetime = date.parse(data['Time'], 'YYYY-MM-DD hh:mm:ssZZ')

    var toTime = function (dateStr) {
        var date = new Date(dateStr).toJSON();
        return new Date(+new Date(date) + 8 * 3600 * 1000)
            .toISOString()
            .replace(/T/g, " ")
            .replace(/\.[\d]{3}Z/, "");
    };

    const datetime = toTime(data["Time"]);
    const isResolved = data["IsResolved"] === 1;

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
                <p className="time">{datetime}</p>
                <p className="like">Likes: {data["Likes"]}</p>
            </div>
        </div>
    );
}

export default QuestionCard;
