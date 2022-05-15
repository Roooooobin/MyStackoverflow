import { useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import { useState, useEffect } from "react";

const CheckAuth = () => {
    // const localUid = localStorage.getItem("uid");
    // const [userData, setUserData] = useState({});

    // useEffect(() => {
    //     // GET request using fetch inside useEffect React hook
    //     if (localUid > 0) {
    //         fetch(`http://0.0.0.0:8080/user/get?uid=${20}`)
    //             .then((response) => response.json())
    //             .then((data) => setUserData(data.data));
    //     }
    // }, []);
    // console.log(userData);

    // // const { auth } = useAuth();

    // // if(auth?.uid){
    // //     console.log("has data")
    // // }



    // return userData;

    const { auth } = useAuth();
    const location = useLocation();

    // console.log(auth);

    return (
        auth
    )
};

export default CheckAuth;
