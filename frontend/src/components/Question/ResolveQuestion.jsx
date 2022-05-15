import React from "react";
import {FcOk} from "react-icons/fc";
import "./QuestionLike.scss";
import getCurrUid from "../../api/getCurrUid"

const ADDQUESLIKE_URL = "http://0.0.0.0:8080/question/resolve";

class ResolveQuestion extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clicked: false,
        };
    }

    handleClick = async (e) => {
        const uid = getCurrUid()
        if (uid < 0) {
            alert("please login first!!");
        } else {

            const qid = this.props.qid[0];

            let formData = new FormData();
            formData.append("uid", uid);
            formData.append("qid", qid);

            const addQuesLike = {
                method: "POST",
                // headers: { 'content-type': 'multipart/form-data' },
                withCredentials: true,
                body: formData,
            };

            const response = await fetch(ADDQUESLIKE_URL, addQuesLike);
            if (response.headers.get("Content-Length") > 0) {
                alert("You have already mark this question as resolved!!!");
            } else {
                this.setState({
                    clicked: true,
                });
            }

        }
    };

    render() {
        return (
            <div className="like">
                <button value={"resolve"} onClick={this.handleClick}>
                    <FcOk/>
                </button>
            </div>
        );
    }
}

export default ResolveQuestion;
