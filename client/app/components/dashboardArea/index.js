import React from 'react';
import './styles.css'
import Card from '../card'

class DashboardArea extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className="dashboardArea">
                {this.props.cards.map(card => (
                    <Card key={card.serviceName + card.methodName} data={card} />
                ))}
            </div>
        );
    }

}

export default DashboardArea;

