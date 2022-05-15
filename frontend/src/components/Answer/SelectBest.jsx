import React from "react";
import { FcOk } from "react-icons/fc";
import "./AnswerLike.scss";
import getCurrUid from "../../api/getCurrUid"

const ADDQUESLIKE_URL = "http://0.0.0.0:8080/answer/select";

class SelectBest extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clicked: false,
        };
    }

    handleClick = async (e) => {
        const uid = getCurrUid()
        if (uid<0) {
            alert("please login first!!");
        } else {
            if (!this.state.clicked) {
                const aid = this.props.aid;

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
                    alert("Can not select a best answer if you are not the user who post the question!");
                } else {
                    this.setState({
                        clicked: true,
                    });
                }
            }else{
                alert("You have already select this answer as best!!!");
            }
        }
    };
    render() {
        return (
            <div className="like">
                <button value={'select'} onClick={this.handleClick}>
                    <FcOk />
                </button>
            </div>
        );
    }
}

export default SelectBest;
