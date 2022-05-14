import { useContext } from "react";
import AuthContext from "../context/authProvidor";

const useAuth = ()=>{
    return useContext(AuthContext);
};

export default useAuth;