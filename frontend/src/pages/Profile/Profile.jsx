import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Profile.scss";
import { Link } from "react-router-dom";
import CheckAuth from "../../api/CheckAuth";


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
  }

  state = {
    quser: null,
  };

  async componentDidMount() {
    const uid = this.props.uid;
    const respUser = await fetch(
        `http://0.0.0.0:8080/user/get?uid=${uid}`
    );
    const quser = await respUser.json();

    this.setState({ quser });
  }

  render() {
    const { quser } = this.state;
    let upart;
    if (quser) {
      const udata = quser.data;
      upart = (
          <div className="profilePart">
            <div className="body">
              <h2>{udata["Username"]}</h2>
            </div>
            <div className="body">
              <p>email: {udata["Email"]}</p>
            </div>
            <div className="body">
              <p>city: {udata["City"]}</p>
            </div>
            <div className="body">
              <p>state: {udata["State"]}</p>
            </div>
            <div className="body">
              <p>country: {udata["Country"]}</p>
            </div>
            <div className="body">
              <p>profile: {udata["Profile"]}</p>
                <div>
                    {udata["Uid"] === 4 ? (<Link to={`/profile/edit/${udata["Uid"]}`}>Edit</Link>) : ''}
                </div>
            </div>

          </div>
      );
    }
    return (
        <div className="main">
          <Header search={true} />
          <div>
            {upart}
          </div>
        </div>
    );
  }
}
export default Profile;
