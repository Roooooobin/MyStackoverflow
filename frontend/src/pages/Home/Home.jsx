import React from "react";
import Header from "../../components/Header/Header";
import Lister from "../../components/Lister/Lister";
import AddQuestion from "../../components/Question/AddQuestion";
import QuestionCard from "../../components/Question/QuestionCard";

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
                    {results &&  <ul className="index">
                            {results.data.map(function (data) {
                                return <QuestionCard data={data} />;
                            })}
                        </ul>}
                </div>
            </div>
        );
    }
}

export default Home;
