import React from "react";
import Header from "../../components/Header/Header";
import Lister from "../../components/Lister/Lister";
import AddQuestion from "../../components/Question/AddQuestion";

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
                <br />
                <AddQuestion />
                <div className="listQuestions">
                    {results && <Lister totalData={results["data"]} question={true} currUid={-1}/>}
                </div>
            </div>
        );
    }
}

export default Home;
