import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Profile.scss";
import { Link } from "react-router-dom";

function Edit() {
  const params = useParams();
  const uid = useState(params.uid);
  return (
      <div>
        <ProfileEditHelper uid={uid} />
      </div>
  );
}


class ProfileEditHelper extends React.Component {
  constructor(props) {
    super(props);
  }

  state = {
    question: null,
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
                <form>
                    <div className="body">
                        <label>
                            username:
                            <input type="text" name="username" defaultValue={udata["Username"]} />
                        </label>
                    </div>
                    <br />
                    <div className="body">
                        <label>
                            city:
                            <input type="text" name="city" defaultValue={udata["City"]} />
                        </label>
                    </div>
                    <br />
                    <div className="body">
                        <label>
                            state:
                            <input type="text" name="state" defaultValue={udata["State"]} />
                        </label>
                    </div>
                    <br />
                    <div className="body">
                        <label>
                            country:
                            <input type="text" name="country" defaultValue={udata["Country"]} />
                        </label>
                    </div>
                    <br />
                    <div className="body">
                        <label>
                            profile:
                            <input size="50" type="text" name="profile" defaultValue={udata["Profile"]} />
                        </label>
                    </div>
                    <br />
                    <input type="submit" value="Submit" />
                </form>
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
export default Edit;
