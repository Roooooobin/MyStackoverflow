import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Profile.scss";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";
import useAuth from "../../hooks/useAuth";
import getCurrUid from "../../api/getCurrUid";

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
    }

    state = {
        quser: null,
        logout: false,
    };

    async componentDidMount() {
        const uid = this.props.uid[0];
        const respUser = await fetch(`http://0.0.0.0:8080/user/get?uid=${uid}`);
        const quser = await respUser.json();

        this.setState({ quser });
    }

    handleClick = async (e) => {
        console.log("logout");
        localStorage.setItem("uid", -1);
        this.setState({ logout: true });
    };

    checkUser = (uid, currUid) => {
      return uid === currUid
    }

    render() {
        const { quser } = this.state;
        let upart;
        let thisUser
        if (quser) {
            const udata = quser.data;
            thisUser = this.checkUser(this.props.uid[0], this.currUid)
            upart = (
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
                        {thisUser && <Link to={`/profile/edit/${udata["Uid"]}`}>Edit</Link>}
                    </div>
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
                        {thisUser && <button className="logout" onClick={this.handleClick}>
                            Log out
                        </button>}
                    </div>
                )}
            </>
        );
    }
}
export default Profile;
