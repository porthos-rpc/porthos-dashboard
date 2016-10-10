import React from 'react';
import './styles.css'

class Menu extends React.Component {

    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.state = {since: 30}
        this.items = [
            {value: 30, label: '30 Minutes'},
            {value: 60, label: '1 Hora'},
            {value: 180, label: '3 Horas'},
            {value: 720, label: '12 Horas'},
            {value: 1440, label: '1 Dia'}
        ]
    }

    handleChange (e) {
        this.setState({since: e.target.value});
        this.props.onChangeSince(e)
    }

    render() {
        return (
            <div className="since-box">
                <span className="label">Since:</span>
                {this.items.map(item => (
                    <label key={item.value}>
                        <input key={item.value} name="since" type="radio" id={item.value} onChange={this.handleChange} value={item.value} checked={this.state.since == item.value} />
                        {item.label}
                    </label>
                ))}
            </div>
        );
    }

}

export default Menu;

