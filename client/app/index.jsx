import React from 'react';
import {render} from 'react-dom';
import Menu from './components/menu'
import DashboardArea from './components/dashboardArea'

class App extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.state = { cards: [], since: '-30m' }
    }

    render () {
        return (
            <div>
                <Menu onChangeSince={this.handleChange} />
                <DashboardArea cards={this.state.cards} />
            </div>
        );
    }

    componentDidMount() {
        this.newFetchCycle(this.state.since)
    }

    handleChange(e) {
        this.setState({
            since: e.target.value
        }) 

        this.newFetchCycle(e.target.value)
    }

    newFetchCycle(since) {
        this.fetchData(since)
        
        if (this.interval) {
            clearInterval(this.interval)
        }

        this.interval = setInterval(() => this.fetchData(this.state.since), 60 * 1000)
    }

    fetchData(since) {
        let request = new XMLHttpRequest()

        request.open('GET', '/api/methods?since=' + since, true)

        request.onload = function() {
            this.setState({
                cards: JSON.parse(request.response)
            })
        }.bind(this); 

        request.send()
    }
}

render(<App/>, document.getElementById('app'));
