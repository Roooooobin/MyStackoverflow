import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Profile.scss";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";
import useAuth from "../../hooks/useAuth";
import getCurrUid from "../../api/getCurrUid";
import Edit from "./Edit/Edit";

function Profile() {
    const params = useParams();
    const uid = useState(params.uid);

    return (
        <div>
            <ProfileHelper uid={uid} />
        </div>
    );
}

class ProfileHelper extends React.Component {
    constructor(props) {
        super(props);
        this.currUid = getCurrUid();
        this.handleEditChange = this.handleEditChange.bind(this)
    }

    state = {
        isEdit:false,
        quser: null,
        logout: false,
        eUsername: "",
        ePassword: "",
        eEmail: "",
        eCity: "",
        eState: "",
        eCountry: "",
        eProfile: "",
    };

    async componentDidMount() {
        const uid = this.props.uid[0];
        const respUser = await fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`);
        const quser = await respUser.json();

        this.setState({
            quser: quser,
            eUsername: quser.Username,
            ePassword: quser.Password,
            eEmail: quser.eEmail,
            eCity: quser.eCity,
            eState: quser.eState,
            eCountry: quser.eCountry,
            eProfile: quser.eProfile,
        });
    }

    handleClick = async (e) => {
        console.log("logout");
        localStorage.setItem("uid", -1);
        this.setState({ logout: true });
    };

    handleEditclick = async (e) => {
        this.setState({ isEdit: true });
    };

    handleFinishEdit = async (e) => {
        this.setState({ isEdit: false });
    };

    handleEditChange(e) {
        const id = e.target.name
        this.setState({ [id]: e.target.value });
        console.log(this.state.eUsername)
    }

    checkUser = (uid, currUid) => {
        return uid === currUid;
    };

    render() {
        const { quser } = this.state;
        let upart;
        let thisUser;
        if (quser) {
            const udata = quser.data;
            console.log("profile", udata)
            thisUser = this.checkUser(this.props.uid[0], this.currUid);
            upart = (
                <div>
                    {this.state.isEdit ? (
                        <div className="editPart">
                            <Edit isEdit={this.props.isEdit} udata={this.state.quser.data} handleEditclick={this.handleFinishEdit}/>
                        </div>
                    ) : (
                        <div className="profilePart">
                            <div className="body">
                                <h2>{udata["Username"]}</h2>
                            </div>
                            <div className="body">
                                <p>email: {udata["Email"]}</p>
                                <p>city: {udata["City"]}</p>
                                <p>state: {udata["State"]}</p>
                                <p>country: {udata["Country"]}</p>
                            </div>
                            <div className="profile">
                                <p>profile: {udata["Profile"]}</p>
                            </div>
                            <div className="edit">
                                {thisUser && (
                                    <button onClick={this.handleEditclick}>
                                        Edit
                                    </button>
                                )}
                            </div>
                        </div>
                    )}
                </div>
            );
        }
        return (
            <>
                {this.state.logout ? (
                    <section>
                        <h2>you have logged out</h2>
                        <Link to={"/"}>Home</Link>
                    </section>
                ) : (
                    <div className="main">
                        <Header search={true} />
                        <div>{upart}</div>
                        {thisUser && !this.state.isEdit && (
                            <button
                                className="logout"
                                onClick={this.handleClick}
                            >
                                Log out
                            </button>
                        )}
                    </div>
                )}
            </>
        );
    }
}
export default Profile;
