import { useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";

const CheckAuth = () => {
    const { auth } = useAuth();
    const location = useLocation();

    console.log(auth);

    return (
        auth
    )
};

export default CheckAuth;