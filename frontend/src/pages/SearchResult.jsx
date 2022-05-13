import React from "react";
import { useParams } from "react-router";
import QuestionCard from "../components/Question/QuestionCard";
import Header from "../components/Header/Header";

function Result(){
    const params = useParams()

    return <SearchResult q={params.q} />
}


class SearchResult extends React.Component {
    constructor(props){
      super(props);
    }

    state = {
        questions:null,
        q:null
    }

    async componentDidMount() {
        const q = this.props.q
        this.setState({q})
        // GET request using fetch with async/await
        const response = await fetch(`http://0.0.0.0:8080/keyword_search/list?keyword=${q}`);
        const questions = await response.json();
        this.setState({questions})
    }

    render() {
        // let q = 1234
        let fakeDatas = [1, 2, 3];
        const listQid = fakeDatas.map((data) => <QuestionCard qid={data} />);
        return (
            <div>
                <Header></Header>
                <h2>This is SearchResult Page!!!</h2>
                <p>The answer for "{this.state.q}" is:</p>
                <ul>{listQid}</ul>

                {/* <button onClick={submitHandler}>click me</button> */}
            </div>
        );
    }
}

export default Result;
