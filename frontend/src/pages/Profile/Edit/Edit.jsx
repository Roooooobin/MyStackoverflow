import React, { useState } from "react";
import { useParams } from "react-router";
import Header from "../../../components/Header/Header";
import { FcOk } from "react-icons/fc";
import "./Profile.scss";
import { Link } from "react-router-dom";
import { getUserData } from "../../../api/getUserData";
import getCurrUid from "../../../api/getCurrUid";
import md5 from "js-md5";

const EDITUSER_URL = `http://0.0.0.0:8080/user/edit`;

class Edit extends React.Component {
    constructor(props) {
        super(props);
    }

    state = {
        eUsername: this.props.udata.Username,
        ePassword: this.props.udata.Password,
        eCity: this.props.udata.City,
        eEmail: this.props.udata.Email,
        eCountry: this.props.udata.Country,
        eState: this.props.udata.State,
        eProfile: this.props.udata.Profile,
    };

    submitHandler = async (e) => {
        e.preventDefault();
        const uid = getCurrUid();
        let formData = new FormData();
        let changed = false;
        if (this.state.eUsername !== this.props.udata.Username) {
            changed = true;
            this.props.udata.Username = this.state.eUsername;
            formData.append("username", this.state.eUsername);
        }
        if (this.state.ePassword !== this.props.udata.Password) {
            changed = true;
            this.props.udata.Password = this.state.ePassword;
            formData.append("password", this.state.ePassword);
        }
        if (this.state.eEmail !== this.props.udata.Email) {
            changed = true;
            this.props.udata.Email = this.state.eEmail;
            formData.append("email", this.state.eEmail);
        }
        if (this.state.eCity !== this.props.udata.City) {
            changed = true;
            this.props.udata.City = this.state.eCity;
            formData.append("city", this.state.eCity);
        }
        if (this.state.eState !== this.props.udata.State) {
            changed = true;
            this.props.udata.State = this.state.eState;
            formData.append("state", this.state.eState);
        }
        if (this.state.eCountry !== this.props.udata.Country) {
            changed = true;
            this.props.udata.Country = this.state.eCountry;
            formData.append("country", this.state.eCountry);
        }
        if (this.state.eProfile !== this.props.udata.Profile) {
            changed = true;
            this.props.udata.Profile = this.state.eProfile;
            formData.append("profile", this.state.eProfile);
        }

        if (changed) {
            formData.append("uid", uid);
            const editUser = {
                method: "POST",
                withCredentials: true,
                body: formData,
            };

            const response = await fetch(EDITUSER_URL, editUser);
            alert("finish update!!!")
        } else {
            alert("no change has made!!!");
        }
    };

    handleBackClick = (e) => {
        this.props.handleEditclick(e.target.value);
    };

    handleChange = (e) => {
        let id = e.target.id;
        this.setState({ [id]: e.target.value });
    };

    render() {
        const udata = this.props.udata;
        return (
            <div>
                {udata && (
                    <form onSubmit={this.submitHandler}>
                        <label htmlFor="eUsername">Username</label>
                        <input
                            type="text"
                            id="eUsername"
                            onChange={this.handleChange}
                            value={this.state.eUsername}
                            required
                        />
                        <br />
                        <label htmlFor="ePassword">Password</label>
                        <input
                            type="password"
                            id="ePassword"
                            onChange={this.handleChange}
                            value={this.state.ePassword}
                            required
                        />
                        <br />
                        <label htmlFor="eEmail">Email</label>
                        <input
                            type="text"
                            id="eEmail"
                            onChange={this.handleChange}
                            value={this.state.eEmail}
                            required
                        />
                        <br />
                        <label htmlFor="eCity">City</label>
                        <input
                            type="text"
                            id="eCity"
                            onChange={this.handleChange}
                            value={this.state.eCity}
                            required
                        />
                        <br />
                        <label htmlFor="eState">State</label>
                        <input
                            type="text"
                            id="eState"
                            onChange={this.handleChange}
                            value={this.state.eState}
                            required
                        />
                        <br />
                        <label htmlFor="eCountry">Country</label>
                        <input
                            type="text"
                            id="eCountry"
                            onChange={this.handleChange}
                            value={this.state.eCountry}
                            required
                        />
                        <br />
                        <label htmlFor="eProfile">Profile</label>
                        <textarea
                            className="eProfile"
                            id="eProfile"
                            onChange={this.handleChange}
                            value={this.state.eProfile}
                            required
                        />
                        <br />
                        <button>Finish Edit</button>
                    </form>
                )}
                <button onClick={this.handleBackClick}>Back</button>
            </div>
        );
    }
}

export default Edit;
