import React from "react";
import { Link } from "react-router-dom";

import "./Header.scss";
import "./SearchBar.jsx";
import { FaBeer } from "react-icons/fa";
import SearchBar from "./SearchBar.jsx";

// class Header extends Component{

//     render(){
//         titleHandler = e =>{
//             const history = useNavigate();
//             history.push(`/`)
//         };

//         return (
//             <div className="header">
//               <div className="title" >
//                 <h2>MyStackOverflow</h2>
//                 <FaBeer className="icon" />
//               </div>
//               <div className="search-bar">
//                 <SearchBar placeholder="Enter Your Question..."/>
//               </div>
//               <div className="login-btn">
//                 <button><Link to="/login">Login</Link></button>
//                 <button><Link to="/signup">Sign Up</Link></button>
//               </div>
//             </div>
//           );
//     }
// }

function Header() {
  return (
    <div className="header">
      <Link to="/" className="title" style={{ textDecoration: 'none' }}>
        <h2>MyStackOverflow</h2>
        <FaBeer className="icon" />
      </Link>
      <div className="search-bar">
        <SearchBar placeholder="Enter Your Question..." />
      </div>
      <div className="login-btn">
        <button>
          <Link to="/login">Login</Link>
        </button>
        <button>
          <Link to="/signup">Sign Up</Link>
        </button>
      </div>
    </div>
  );
}

export default Header;
