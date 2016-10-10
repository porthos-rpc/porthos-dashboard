import React from 'react';
import { label } from './styles.css'

class Menu extends React.Component {

  constructor(props) {
    super(props);
    this.state = {since: 30};
    this.onChangeSince = this.onChangeSince.bind(this);
  }

  onChangeSince () {
    this.setState({since: 60});
  }

  render() {
    return (
      <div className="since-box">
        <span className="label">Since:</span>
        <input name="since" type="radio" onChange={this.onChangeSince} value="30" checked="checked"/> 30 Minutes
        <input name="since" type="radio" onChange={this.onChangeSince} value="60"/> 60 Minutes
        <input name="since" type="radio" onChange={this.onChangeSince} value="180"/> 3 Hour
        <input name="since" type="radio" onChange={this.onChangeSince} value="720"/> 6 Hours
        <input name="since" type="radio" onChange={this.onChangeSince} value="1440"/> 1 Day
      </div>
    );
  }

}

export default Menu;

