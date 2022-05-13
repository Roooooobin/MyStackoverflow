import React from "react";
import { useState } from "react";
import { useParams } from "react-router";
import QuestionCard from "../components/Question/QuestionCard";
import Header from "../components/Header/Header";

function Result() {
    const params = useParams();
    const q = useState(params.q);
    console.log("1",q);

    return (
        <div>
            <Header />
            <SearchResult q={q} />
        </div>
    );
}

class SearchResult extends React.Component {
    constructor(props) {
        super(props);
        
    }

    state = {
        results: null,
        q: null,
    };


    async componentDidMount() {
        const q = this.props.q;
        this.setState({ q });
        // GET request using fetch with async/await
        const response = await fetch(
            `http://0.0.0.0:8080/keyword_search/list?keyword=${q}`
        );
        const results = await response.json();
        this.setState({ results });
    }

    render() {
        const { results } = this.state;
        return (
            <div>
                <p>The answer for "{this.state.q}" is:</p>
                <div>
                    {results &&
                        results["data"]["questions"].map((data) => (
                            <QuestionCard qid={data["Qid"]} />
                        ))}
                </div>

                {/* <button onClick={submitHandler}>click me</button> */}
            </div>
        );
    }
}

export default Result;
