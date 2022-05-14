import React from "react";
import Header from "../../components/Header/Header";
import Lister from "../../components/Lister/Lister";

class Home extends React.Component {
    state = {
        results: null,
    };

    async componentDidMount() {
        const response = await fetch("http://0.0.0.0:8080/question/list");
        const results = await response.json();
        this.setState({ results });
    }

    render() {
        const { results } = this.state;

        return (
            <div>
                <Header search={true} />
                {/* <label>This is Home!!!</label> */}
                <div className="listQuestions">
                    {results && <Lister totalData={results["data"]} question={true}/>}
                </div>
            </div>
        );
    }
}

export default Home;
