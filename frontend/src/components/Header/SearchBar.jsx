import { useState } from "react";
import { FaSearch } from "react-icons/fa";
import { Link } from "react-router-dom";

function SearchBar({ placeholder }) {
  var query = "";
  const [searchQuery, setSearchQuery] = useState(query || "");

  return (
    <div className="search">
      <div className="searchInputs">
        <input
          type="text"
          placeholder={placeholder}
          onInput={(e) => setSearchQuery(e.target.value)}
          value={searchQuery}
        />
        <div className="searchIcon">
          <Link to={`/result/${searchQuery}`} onClick={this.forceUpdate()}>
            <FaSearch className="icon" />
          </Link>
          
        </div>
      </div>
    </div>
  );
}

export default SearchBar;
