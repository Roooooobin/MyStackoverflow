import React from "react";
import { useState } from "react";
import { useParams } from "react-router";
import Header from "../components/Header/Header";
import Lister from "../components/Lister/Lister";
import AddQuestion from "../components/Question/AddQuestion";

function Result() {
    const params = useParams();
    const q = useState(params.q);
    console.log("1", q);

    return (
        <div>
            <SearchResult q={q} />
        </div>
    );
}

class SearchResult extends React.Component {
    constructor(props) {
        super(props);
    }

    state = {
        topics: null,
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

        // const rspTopic = await fetch(`http://0.0.0.0:8080/topic/list`);
        // const resTopic = await rspTopic.json();
        // const topics = resTopic.data;
        // console.log("search reasult topic ", topics)
        this.setState({ results });
    }

    render() {
        const { results } = this.state;
        let lister;
        if (results) {
            if (Object.keys(results.data.questions).length !== 0) {
                lister = (
                    <div>
                        <h2>The Question for "{this.state.q}" is:</h2>{" "}
                        <Lister
                            totalData={results.data.questions}
                            question={true}
                        />
                    </div>
                );
            } else {
                lister = (
                    <>
                        <h2>
                            Sorry, We don't find any question about "
                            {this.state.q}
                            "!
                        </h2>
                        <AddQuestion />
                    </>
                );
            }
        } else {
            lister = (
                <>
                    <h2>
                        Sorry, We don't find any question about "{this.state.q}
                        "!
                    </h2>
                    <AddQuestion/>
                </>
            );
        }
        return (
            <div>
                <Header search={false} />

                <div>{lister}</div>
            </div>
        );
    }
}

export default Result;
