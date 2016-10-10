import React from 'react';
import {render} from 'react-dom';
import Menu from './components/menu'
import DashboardArea from './components/dashboardArea'

class App extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.state = {cards: [
            {serviceName: "UserContent", methodName: "getUserContents", avgResponseTime: 120, avgThroughput: 5533},
            {serviceName: "UserContent", methodName: "getContentById", avgResponseTime: 20, avgThroughput: 1402},
            {serviceName: "UserContent", methodName: "createContent2", avgResponseTime: 55, avgThroughput: 145},
            {serviceName: "UserContent", methodName: "createContent3", avgResponseTime: 55, avgThroughput: 145},
            {serviceName: "UserContent", methodName: "createContent4", avgResponseTime: 55, avgThroughput: 145},
            {serviceName: "UserContent", methodName: "createContent5", avgResponseTime: 55, avgThroughput: 145},
        ], range: 30};
    }

    render () {
        return (
            <div>
                <Menu onChangeSince={this.handleChange} />
                <DashboardArea cards={this.state.cards} />
            </div>
        );
    }

    handleChange(e) {
        this.setState({
            range: e.target.value
        }) 
    }
}

render(<App/>, document.getElementById('app'));
