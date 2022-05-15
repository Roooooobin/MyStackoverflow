import React from "react";
import { FcLike } from "react-icons/fc";
import "./AnswerLike.scss";
import getCurrUid from "../../api/getCurrUid"

const ADDQUESLIKE_URL = "http://0.0.0.0:8080/answer/like";

class AnswerLike extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clicked: false,
            num: this.props.likes ? this.props.likes : 0,
        };
    }

    handleClick = async (e) => {
        const uid = getCurrUid()
        if (uid<0) {
            alert("please login first!!");
        } else {
            if (!this.state.clicked) {
                const aid = this.props.aid;
                const val = this.state.num + 1;

                let formData = new FormData();
                formData.append("uid", uid);
                formData.append("aid", aid);

                const addAnsLike = {
                    method: "POST",
                    // headers: { 'content-type': 'multipart/form-data' },
                    withCredentials: true,
                    body: formData,
                };

                const response = await fetch(ADDQUESLIKE_URL, addAnsLike);
                if (response.headers.get("Content-Length") > 0) {
                    alert("You have already liked this Answer!!!");
                } else {
                    this.setState({
                        num: val,
                        clicked: true,
                    });
                }
            }else{
                alert("You have already liked this Answer!!!");
            }
        }
    };
    render() {
        const val = this.state.num;
        return (
            <div className="like">
                <button value={Number(val)} onClick={this.handleClick}>
                    <FcLike />
                </button>
                <h3>{val}</h3>
            </div>
        );
    }
}

export default AnswerLike;
