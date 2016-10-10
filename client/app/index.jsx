import React from 'react';
import {render} from 'react-dom';
import Menu from './components/menu'
import DashboardArea from './components/dashboardArea'

class App extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.state = {cards: [
            {serviceName: "UserContent", methodName: "getUserContents", avgResponseTime: 120, avgThroughput: 5533, minResponseTime: 10, maxResponseTime: 120, minThroughput: 20, maxThroughput: 5533,
                throughputHistory: [{value: 4000}, {value: 3000}, {value: 2000}, {value: 2780}, {value: 1890}, {value: 2390}, {value: 1490}]},
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
