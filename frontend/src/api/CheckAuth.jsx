import { useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";
import { useState, useEffect } from "react";
import getCurrUid from "./getCurrUid";
import { getUserData } from "./getUserData";

const CheckAuth = () => {
    const localUid = getCurrUid();
    const [userData, setUserData] = useState({});
    const { auth } = useAuth();

    useEffect(() => {
        // GET request using fetch inside useEffect React hook
        let mounted = true;
        if(localUid > 0 && !auth?.uid){
            getUserData(localUid).then((data)=>{
                if(mounted){
                    setUserData(data)
                    auth = userData
                }
            });
        }
        return () => (mounted = false)
    }, []);

    console.log(userData);

    // const { auth } = useAuth();

    // if(auth?.uid){
    //     console.log("has data")
    // }



    

    // const { auth } = useAuth();
    // const location = useLocation();

    return (
        auth
    )
};

export default CheckAuth;
