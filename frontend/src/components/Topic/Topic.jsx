import Select from "react-select";
import React, { useState } from "react";

// const options = [
//     { value: "1", label: "1" },
//     { value: "strawberry", label: "Strawberry" },
//     { value: "vanilla", label: "Vanilla" },
// ];

class Topic extends React.Component {
    state = {
        selectedOption: null,
        options: null
    };
    handleChange = (selectedOption) => {
        this.setState({ selectedOption }, () =>
            console.log(`Option selected:`, this.state.selectedOption)
        );
    };

    async componentDidMount() {
        const response = await fetch(`http://0.0.0.0:8080/topic/list`);
        const results = await response.json();
        this.setState({ options: results.data })
    }

    render() {
        const { options, selectedOption } = this.state;

        return (
            <Select
                value={selectedOption}
                onChange={this.handleChange}
                options={options}
            />
        );
    }
}

export default Topic;
