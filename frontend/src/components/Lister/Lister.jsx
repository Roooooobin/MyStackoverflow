import React from "react";
import QuestionCard from "../Question/QuestionCard";
import "./Lister.scss";

class Lister extends React.Component{
    constructor(props){
        super(props);
        this.pageNext=this.pageNext.bind(this);
        this.setPage=this.setPage.bind(this);
        this.state = {
            indexList:[],
            totalData:this.props.totalData,
            current: 1, 
            pageSize:10, 
            goValue:0,  
            totalPage:0,
        };

    }

    componentWillMount(){
        this.setState({
            totalPage:Math.ceil( this.state.totalData.length/this.state.pageSize),
        })
        this.pageNext(this.state.goValue)

    }

    setPage(num){
        this.setState({
            indexList:this.state.totalData.slice(num,num+this.state.pageSize)
        })
    }


    pageNext (num) {
        this.setPage(num)
    }

    render() {

        return (
            <div className="main">
                <div className="top_bar">
                </div>
                <div className="lists">
                    <ul className="index">
                        {this.state.indexList.map(function (data) {
                            return <QuestionCard data={data} />
                        })}
                    </ul>

                    <PageButton { ...this.state } pageNext={this.pageNext} />

                </div>
            </div>
        );
    }
}

class PageButton extends React.Component {

    constructor(props) {
        super(props);
        this.setNext=this.setNext.bind(this);
        this.setUp=this.setUp.bind(this);
        this.state={
            num: 0,
            pagenum:this.props.current
        }
    }

    //下一页
    setNext(){
        if(this.state.pagenum < this.props.totalPage){
            this.setState({
                num:this.state.num + this.props.pageSize,
                pagenum:this.state.pagenum + 1
            },function () {
                console.log(this.state)
                this.props.pageNext(this.state.num)
            })
        }
    }

    //上一页
    setUp(){
        if(this.state.pagenum > 1){
            this.setState({
                num:this.state.num - this.props.pageSize,
                pagenum:this.state.pagenum - 1
            },function () {
                console.log(this.state)
                this.props.pageNext(this.state.num)
            })
        }
    }

    render() {
        return (
            <div className="change_page">
                <button onClick={ this.setUp } >Previous</button>
                <span>{ this.state.pagenum }Page/{ this.props.totalPage } Pages </span>
                <button onClick={ this.setNext }>Next</button>

            </div>
        );
    }
}

export default Lister;